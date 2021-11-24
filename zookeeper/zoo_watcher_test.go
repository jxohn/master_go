package zookeeper

import (
	"log"
	"testing"
)

func TestZookeeperChildW(t *testing.T) {
	log.Println("zookeeper_watcher.go test begin...")
	ChildW()
	log.Println("zookeeper_watcher.go test end...")
}
