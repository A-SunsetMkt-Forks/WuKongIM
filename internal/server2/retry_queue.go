package server

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

// RetryQueue 重试队列
type RetryQueue struct {
	inFlightPQ       inFlightPqueue
	inFlightMessages map[string]*retryMessage
	inFlightMutex    sync.Mutex
	s                *Server
	fakeMessageID    int64
}

// NewRetryQueue NewRetryQueue
func NewRetryQueue(s *Server) *RetryQueue {

	return &RetryQueue{
		inFlightPQ:       newInFlightPqueue(1024),
		inFlightMessages: make(map[string]*retryMessage),
		s:                s,
		fakeMessageID:    10000,
	}
}

func (r *RetryQueue) startInFlightTimeout(msg *retryMessage) {
	now := time.Now()
	msg.pri = now.Add(r.s.opts.MessageRetry.Interval).UnixNano()
	r.pushInFlightMessage(msg)
	r.addToInFlightPQ(msg)

}

func (r *RetryQueue) addToInFlightPQ(msg *retryMessage) {
	r.inFlightMutex.Lock()
	defer r.inFlightMutex.Unlock()
	r.inFlightPQ.Push(msg)

}
func (r *RetryQueue) pushInFlightMessage(msg *retryMessage) {
	r.inFlightMutex.Lock()
	defer r.inFlightMutex.Unlock()
	key := r.getInFlightKey(msg.connId, msg.messageId)
	_, ok := r.inFlightMessages[key]
	if ok {
		r.s.Warn("ID already in flight", zap.Int64("connId", msg.connId), zap.Int64("messageId", msg.messageId))
		return
	}
	r.inFlightMessages[key] = msg

}

func (r *RetryQueue) popInFlightMessage(connId int64, messageId int64) (*retryMessage, error) {
	r.inFlightMutex.Lock()
	defer r.inFlightMutex.Unlock()
	key := r.getInFlightKey(connId, messageId)
	msg, ok := r.inFlightMessages[key]
	if !ok {
		r.s.Warn("ID not in flight", zap.Int64("connId", msg.connId), zap.Int64("messageId", msg.messageId))
		return nil, errors.New("ID not in flight")
	}
	delete(r.inFlightMessages, key)
	return msg, nil
}

func (r *RetryQueue) getInFlightKey(connId int64, messageId int64) string {
	var b strings.Builder
	b.WriteString(strconv.FormatInt(connId, 10))
	b.WriteString(":")
	b.WriteString(strconv.FormatInt(messageId, 10))
	return b.String()
}
func (r *RetryQueue) finishMessage(connId int64, messageId int64) error {
	msg, err := r.popInFlightMessage(connId, messageId)
	if err != nil {
		return err
	}
	r.removeFromInFlightPQ(msg)

	return nil
}
func (r *RetryQueue) removeFromInFlightPQ(msg *retryMessage) {
	r.inFlightMutex.Lock()
	if msg.index == -1 {
		// this item has already been popped off the pqueue
		r.inFlightMutex.Unlock()
		return
	}
	r.inFlightPQ.Remove(msg.index)
	r.inFlightMutex.Unlock()
}

func (r *RetryQueue) processInFlightQueue(t int64) {
	for {
		r.inFlightMutex.Lock()
		msg, _ := r.inFlightPQ.PeekAndShift(t)
		r.inFlightMutex.Unlock()
		if msg == nil {
			break
		}
		err := r.finishMessage(msg.connId, msg.messageId)
		if err != nil {
			r.s.Error("processInFlightQueue-finishMessage失败", zap.Error(err), zap.Int64("connId", msg.connId), zap.Int64("messageId", msg.messageId))
			break
		}
		r.s.retryManager.retry(msg) // 重试
	}
}

// Start 开始运行重试
func (r *RetryQueue) Start() {
	r.s.Schedule(r.s.opts.MessageRetry.ScanInterval, func() {
		now := time.Now().UnixNano()
		r.processInFlightQueue(now)
	})
}

func (r *RetryQueue) Stop() {
}