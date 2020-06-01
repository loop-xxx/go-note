package session

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
)

type MemSession struct{
	core map[string]interface{}
	rwm  sync.RWMutex
	last time.Time
}

func (mp *MemSession)Get(key string) (value interface{}, ok bool){
	mp.rwm.RLock()
	value, ok = mp.core[key]
	mp.rwm.RUnlock()
	return
}

func (mp *MemSession)Set(key string , value interface{}){
	mp.rwm.Lock()
	mp.rwm.RLocker()
	mp.core[key] = value
	mp.rwm.Unlock()
}

func (mp *MemSession)Del(key string){
	delete(mp.core, key)
}

func newMS(cacheSize uint) (mp *MemSession){
	mp = &MemSession{
		core: make(map[string]interface{}, cacheSize),
		rwm:sync.RWMutex{},
	}
	return
}


type MemManager struct{
	core map[string] *MemSession
	rwm sync.RWMutex
	cancel context.CancelFunc
	wg sync.WaitGroup
}

func (mp *MemManager) timeoutHandler(ctx context.Context, timeout time.Duration, interval time.Duration){
	defer mp.wg.Done()
	while:
	for{
		select {
		case now := <- time.Tick(interval):
			mp.rwm.Lock()
			for sessionID, msp := range mp.core{
				if now.Sub(msp.last) > timeout  {
					delete(mp.core, sessionID)
					log.Printf("%s is invailed", sessionID)
				}
			}
			mp.rwm.Unlock()
		case <- ctx.Done():
				break while
		}
	}
}

// BuildMemManager ... 创建Manager对象
func BuildMemManager(cacheSize uint, timeout time.Duration, interval time.Duration)(m Manager){
	mmp := &MemManager{
		core: make(map[string]*MemSession, cacheSize),
		rwm:  sync.RWMutex{},
		wg: sync.WaitGroup{},
	}

	//开启session timeout检测goroutine
	ctx, cancel := context.WithCancel(context.Background())
	mmp.cancel = cancel
	mmp.wg.Add(1)
	go mmp.timeoutHandler(ctx ,timeout, interval)

	//返回接口类型
	m = mmp
	return
}

// Destroy ... 关闭timeout检测goroutine
func (mp *MemManager)Done()(err error){
	mp.cancel()
	mp.wg.Wait()

	return
}

func (mp *MemManager)GetSessionByID(sessionID string)(s Session, exist bool){
	now := time.Now()

	mp.rwm.RLock()
	//每次获取Session时,更新session的最后使用时间
	if msp, ok := mp.core[sessionID]; ok{
		msp.rwm.Lock()
		msp.last = now
		msp.rwm.Unlock()
		s = msp
		exist = ok
	}
	mp.rwm.RUnlock()

	return
}

func (mp *MemManager)NewSession(cacheSize uint) (sessionID string, s Session){
	msp := newMS(cacheSize)
	//初始化session的最后使用时间
	msp.last = time.Now()

	//设置返回值
	sessionID = uuid.NewV4().String()
	s = msp

	//将新创建的Session添加到Manager
	mp.rwm.Lock()
	mp.core[sessionID] = msp
	mp.rwm.Unlock()

	return
}