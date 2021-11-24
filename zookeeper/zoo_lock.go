package zookeeper

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/go-zookeeper/zk"
)

const (
	LockBasePath = "/zlock"
)

var (
	ErrHasBeenLocked = errors.New("HasBeenLocked by others")
)

type ZLocker struct {
	key  string
	path string
	conn *zk.Conn
}

func NewZLocker(key string, conn *zk.Conn) (*ZLocker, error) {
	return &ZLocker{
		key:  key,
		path: "",
		conn: conn,
	}, nil
}

// lock
func (l *ZLocker) lock() error {

	path, err :=
		l.conn.CreateProtectedEphemeralSequential(LockBasePath+"/"+l.key, []byte(""), zk.WorldACL(zk.PermAll))
	if err != nil {
		return fmt.Errorf("error to lock, %+v", err)
	}

	l.path = path

	children, _, err := l.conn.Children(LockBasePath)
	if err != nil {
		return fmt.Errorf("error to lock, %+v", err)
	}
	if len(children) == 0 {
		return fmt.Errorf("error to lock, failed to get lock from zk")
	}
	sortChildren := make([]string, len(children))
	for i, v := range children {
		split := strings.Split(v, "-")
		sortChildren[i] = split[1]
	}

	sort.Strings(sortChildren)
	log.Printf("local path is %s, children is %s", l.path, sortChildren)
	// lock has been got by others
	pathes := strings.Split(l.path, "-")
	if pathes[1] != sortChildren[0] {
		return ErrHasBeenLocked
	}
	// get lock suc!
	return nil
}

// unLock
func (l *ZLocker) unLock() error {
	_, stat, err := l.conn.Get(l.path)
	if err != nil {
		return fmt.Errorf("get lock state error, %+v", err)
	}

	return l.conn.Delete(l.path, stat.Version)
}
