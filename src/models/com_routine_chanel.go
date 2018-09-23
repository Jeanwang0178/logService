package models

import (
	"fmt"
	"github.com/hpcloud/tail"
	"logService/src/common"
	"math/rand"
	"sync"
)

type GoroutineChannel struct {
	gid   uint64
	name  string
	tails tail.Tail
	msg   chan string
}

type GoroutineChannelMap struct {
	mutex      sync.Mutex
	grchannels map[string]*GoroutineChannel
}

func (m *GoroutineChannelMap) unregister(name string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.grchannels[name]; !ok {
		return fmt.Errorf("goroutine channel not find :%q ", name)
	}
	delete(m.grchannels, name)
	return nil
}

func (m *GoroutineChannelMap) register(chanKey string, tails tail.Tail) (chanName string, err error) {
	gchannel := &GoroutineChannel{
		gid:   uint64(rand.Int63()),
		name:  chanKey,
		tails: tails,
	}
	gchannel.msg = make(chan string)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.grchannels == nil {
		m.grchannels = make(map[string]*GoroutineChannel)
	} else if _, ok := m.grchannels[gchannel.name]; ok {
		return "", fmt.Errorf("go routine channel already defined : %q ", gchannel.name)
	}

	m.grchannels[gchannel.name] = gchannel
	common.Logger.Info("register chanName :" + gchannel.name)
	return gchannel.name, nil
}
