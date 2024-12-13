package reactor

import (
	"fmt"
	"time"

	"github.com/WuKongIM/WuKongIM/internal/reactor"
	"github.com/WuKongIM/WuKongIM/pkg/wklog"
	"github.com/lni/goutils/syncutil"
	"github.com/valyala/fastrand"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

type reactorSub struct {
	users        *list // 用户列表
	stopped      atomic.Bool
	actionQueue  *actionQueue  // action队列
	tickInterval time.Duration // tick间隔时间
	stopper      *syncutil.Stopper
	advanceC     chan struct{} // 推进事件
	index        int
	wklog.Log
	// 连续readEvent次数
	continReadEventCount int

	tmpUsers []*User
	r        *Reactor
}

func newReactorSub(index int, r *Reactor) *reactorSub {

	return &reactorSub{
		index:        index,
		r:            r,
		tickInterval: options.TickInterval,
		stopper:      syncutil.NewStopper(),
		advanceC:     make(chan struct{}, 1),
		users:        newList(),
		actionQueue:  newActionQueue(options.ReceiveQueueLength, false, 0, options.MaxReceiveQueueSize),
		Log:          wklog.NewWKLog(fmt.Sprintf("user.reactorSub[%d]", index)),
	}
}

func (r *reactorSub) start() error {
	r.stopper.RunWorker(r.loop)
	return nil
}

func (r *reactorSub) stop() {
	r.stopped.Store(true)
	r.stopper.Stop()
}

func (r *reactorSub) loop() {

	p := float64(fastrand.Uint32()) / (1 << 32)
	// 以避免系统中因定时器、周期性任务或请求间隔完全一致而导致的同步问题（例如拥堵或资源竞争）。
	jitter := time.Duration(p * float64(r.tickInterval/2))
	tick := time.NewTicker(r.tickInterval + jitter)
	defer tick.Stop()

	for !r.stopped.Load() {

		if r.continReadEventCount < 100 {
			// 读取事件
			r.readEvents()
		} else {
			r.continReadEventCount = 0
			r.Warn("too many consecutive ready", zap.Int("continReadEventCount", r.continReadEventCount))
		}

		select {
		// case <-r.advanceC:
		case <-tick.C:
			r.continReadEventCount = 0
			r.tick()
		case <-r.stopper.ShouldStop():
			return
		}
	}
}

func (r *reactorSub) readEvents() {

	hasEvent := false

	event := r.handleEvents()
	if event {
		hasEvent = true
	}
	event = r.handleReceivedActions()
	if event {
		hasEvent = true
	}

	if hasEvent { // 如果有事件 接着推进
		r.continReadEventCount++
		r.advance()
	}
}

// 处理本地事件
func (r *reactorSub) handleEvents() bool {
	r.users.read(&r.tmpUsers)
	hasEvent := true

	for hasEvent && !r.stopped.Load() {
		hasEvent = false
		for _, user := range r.tmpUsers {
			has := r.handleEvent(user)
			if has {
				hasEvent = true
			}
		}
	}
	r.tmpUsers = r.tmpUsers[:0]
	return hasEvent
}

func (r *reactorSub) handleEvent(u *User) bool {
	if !u.hasReady() {
		return false
	}
	actions := u.ready()
	if len(actions) == 0 {
		return false
	}

	for _, action := range actions {
		switch action.Type {
		case reactor.UserActionUserClose:
			r.users.remove(u.uid)
		}
	}

	r.r.send(actions)

	return true
}

// 处理收到的action
func (r *reactorSub) handleReceivedActions() bool {
	actions := r.actionQueue.get()
	if len(actions) == 0 {
		return false
	}
	for _, a := range actions {
		user := r.users.get(a.Uid)
		if user == nil {
			continue
		}
		if a.No != "" && a.No != user.no {
			continue
		}
		user.step(a)
	}

	return true
}

func (r *reactorSub) advance() {
	select {
	case r.advanceC <- struct{}{}:
	default:
	}
}

func (r *reactorSub) tick() {
	r.users.read(&r.tmpUsers)
	if !r.stopped.Load() {
		for _, user := range r.tmpUsers {
			user.tick()
		}
	}

	r.tmpUsers = r.tmpUsers[:0]
}

func (r *reactorSub) addAction(a reactor.UserAction) bool {
	r.Info("addAction==", zap.String("uid", a.Uid), zap.String("type", a.Type.String()))
	return r.actionQueue.add(a)
}

func (r *reactorSub) mustAddAction(a reactor.UserAction) {
	r.actionQueue.mustAdd(a)
}

func (r *reactorSub) user(uid string) *User {

	return r.users.get(uid)

}

func (r *reactorSub) addUser(u *User) {
	r.users.add(u)
}
