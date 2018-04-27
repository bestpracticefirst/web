package session

import (
	"sync"
	"time"
)

type Session interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Remove(key interface{}) error
	GetId() string
}

//session实现
type SessionFromMemory struct {
	//唯一标示
	sid              string
	lock             sync.Mutex                  //一把互斥锁
	lastAccessedTime time.Time                   //最后访问时间
	maxAge           int64                       //超时时间
	data             map[interface{}]interface{} //主数据
}

//实例化
func newSessionFromMemory() *SessionFromMemory {
	return &SessionFromMemory{
		data:   make(map[interface{}]interface{}),
		maxAge: 60 * 30, //默认30分钟
	}
}

//同一个会话均可调用，进行设置，改操作必须拥有排斥锁
func (si *SessionFromMemory) Set(key, value interface{}) {
	si.lock.Lock()
	defer si.lock.Unlock()
	si.data[key] = value
}

func (si *SessionFromMemory) Get(key interface{}) interface{} {
	if value := si.data[key]; value != nil {
		return value
	}
	return nil
}
func (si *SessionFromMemory) Remove(key interface{}) error {
	if value := si.data[key]; value != nil {
		delete(si.data, key)
	}
	return nil
}
func (si *SessionFromMemory) GetId() string {
	return si.sid
}