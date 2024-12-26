package eventbus

import (
	wkproto "github.com/WuKongIM/WuKongIMGoProto"
)

var User *userPlus

type IUser interface {
	// AddEvent 添加事件
	AddEvent(uid string, event *Event)
	// Advance 推进事件（让事件池不需要等待直接执行下一轮事件）
	Advance(uid string)
	// 查询连接信息
	ConnsByUid(uid string) []*Conn
	AuthedConnsByUid(uid string) []*Conn
	ConnCountByUid(uid string) int
	ConnsByDeviceFlag(uid string, deviceFlag wkproto.DeviceFlag) []*Conn
	ConnCountByDeviceFlag(uid string, deviceFlag wkproto.DeviceFlag) int
	ConnById(uid string, fromNode uint64, id int64) *Conn
	LocalConnById(uid string, id int64) *Conn
	LocalConnByUid(uid string) []*Conn

	// UpdateConn 更新连接
	UpdateConn(conn *Conn)
}

type userPlus struct {
	user IUser
}

func newUserPlus(user IUser) *userPlus {
	return &userPlus{
		user: user,
	}
}

func (u *userPlus) AddEvent(uid string, event *Event) {
	u.user.AddEvent(uid, event)
}

func (u *userPlus) AddEvents(uid string, events []*Event) {
	for _, event := range events {
		u.user.AddEvent(uid, event)
	}
	u.user.Advance(uid)
}

func (u *userPlus) Advance(uid string) {
	u.user.Advance(uid)
}

// ========================================== conn ==========================================
// Connect 请求连接
func (u *userPlus) Connect(conn *Conn, connectPacket *wkproto.ConnectPacket) {
	u.user.AddEvent(conn.Uid, &Event{
		Type:  EventConnect,
		Frame: connectPacket,
		Conn:  conn,
	})
}

// ConnsByDeviceFlag 根据设备标识获取连接
func (u *userPlus) ConnsByDeviceFlag(uid string, deviceFlag wkproto.DeviceFlag) []*Conn {
	return u.user.ConnsByDeviceFlag(uid, deviceFlag)
}

func (u *userPlus) ConnCountByDeviceFlag(uid string, deviceFlag wkproto.DeviceFlag) int {
	return u.user.ConnCountByDeviceFlag(uid, deviceFlag)
}

// ConnsByUid 根据用户uid获取连接
func (u *userPlus) ConnsByUid(uid string) []*Conn {
	return u.user.ConnsByUid(uid)
}

// AuthedConnsByUid 根据用户uid获取已认证的连接
func (u *userPlus) AuthedConnsByUid(uid string) []*Conn {
	return u.user.AuthedConnsByUid(uid)
}

func (u *userPlus) ConnCountByUid(uid string) int {
	return u.user.ConnCountByUid(uid)
}

// LocalConnById 获取本地连接
func (u *userPlus) LocalConnById(uid string, id int64) *Conn {
	return u.user.LocalConnById(uid, id)
}

// LocalConnByUid 获取本地连接
func (u *userPlus) LocalConnByUid(uid string) []*Conn {
	return u.user.LocalConnByUid(uid)
}

// ConnById 获取连接
func (u *userPlus) ConnById(uid string, fromNode uint64, id int64) *Conn {
	return u.user.ConnById(uid, fromNode, id)
}

// UpdateConn 更新连接
func (u *userPlus) UpdateConn(conn *Conn) {
	u.user.UpdateConn(conn)
}

// ConnWrite 连接写包
func (u *userPlus) ConnWrite(conn *Conn, frame wkproto.Frame) {
	u.user.AddEvent(conn.Uid, &Event{
		Type:  EventConnWriteFrame,
		Conn:  conn,
		Frame: frame,
	})
}

// CloseConn 关闭连接
func (u *userPlus) CloseConn(conn *Conn) {
	u.user.AddEvent(conn.Uid, &Event{
		Type: EventConnClose,
		Conn: conn,
	})
}