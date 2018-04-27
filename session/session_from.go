package session

import (
	"sync"
	"time"
	"fmt"
)

type Storage interface {
	//初始化一个session，id根据需要生成后传入
	InitSession(sid string, maxAge int64) (Session, error)
	//根据sid，获得当前session
	SetSession(session Session) error
	//销毁session
	DestroySession(sid string) error
	//回收
	GCSession()
}

//session来自内存
type FromMemory struct {
	//由于session包含所有的请求
	//并行时，保证数据独立、一致、安全
	lock     sync.Mutex //互斥锁
	sessions map[string]Session
}

//实例化一个内存实现
func newFromMemory() *FromMemory {
	return &FromMemory{
		sessions: make(map[string]Session, 0),
	}
}

//初始换会话session，这个结构体操作实现Session接口
func (fm *FromMemory) InitSession(sid string, maxAge int64) (Session, error) {
	fm.lock.Lock()
	defer fm.lock.Unlock()

	newSession := newSessionFromMemory()
	newSession.sid = sid
	if maxAge != 0 {
		newSession.maxAge = maxAge
	}
	newSession.lastAccessedTime = time.Now()

	fm.sessions[sid] = newSession //内存管理map
	return newSession, nil
}

//设置
func (fm *FromMemory) SetSession(session Session) error {
	fm.sessions[session.GetId()] = session
	return nil
}

//销毁session
func (fm *FromMemory) DestroySession(sid string) error {
	if _, ok := fm.sessions[sid]; ok {
		delete(fm.sessions, sid)
		return nil
	}
	return nil
}

//监判超时
func (fm *FromMemory) GCSession() {

	sessions := fm.sessions

	//fmt.Println("gc session")

	if len(sessions) < 1 {
		return
	}

	//fmt.Println("current active sessions ", sessions)

	for k, v := range sessions {
		t := (v.(*SessionFromMemory).lastAccessedTime.Unix()) + (v.(*SessionFromMemory).maxAge)

		if t < time.Now().Unix() { //超时了

			fmt.Println("timeout-------->", v)
			delete(fm.sessions, k)
		}
	}

}