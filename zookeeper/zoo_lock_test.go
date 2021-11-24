package zookeeper

import (
	"sync"
	"testing"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/stretchr/testify/assert"
)

func TestZLocker(t *testing.T) {
	// 无抢占加锁
	conn, _, err := zk.Connect([]string{"localhost:2181"}, time.Second)
	assert.Nil(t, err)
	locker, err := NewZLocker("lock", conn)
	assert.Nil(t, err)
	err = locker.lock()
	assert.Nil(t, err)
	err = locker.unLock()
	assert.Nil(t, err)

	// 抢占式加锁
	errorChan := make(chan error, 10)
	waitGroup := sync.WaitGroup{}

	lockers := make([]*ZLocker, 10)
	for i := 0; i < 10; i++ {
		zLocker, err := NewZLocker("lock", conn)
		assert.Nil(t, err)
		lockers[i] = zLocker
	}

	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go func(x int) {
			defer waitGroup.Done()
			err := lockers[x].lock()
			if err != nil {
				errorChan <- err
			}
		}(i)
	}
	waitGroup.Wait()
	close(errorChan)
	i := 0
	for e := range errorChan {
		i++
		assert.Equal(t, ErrHasBeenLocked, e)
	}
	assert.Equal(t, 9, i)
}
