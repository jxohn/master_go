package zookeeper

import (
	"log"
	"time"

	"github.com/go-zookeeper/zk"
)

// ZookeeperChildW
// zookeeper连接, 监听child变更事件
func ChildW() {
	// 1. 建立连接
	conn, events, err := zk.Connect([]string{"localhost:2181"}, time.Second)
	if err != nil {
		log.Panic(err)
	}

	// 连接过程中的event
	select {
	case e := <-events:
		log.Printf("type is %s, state is %s, ", e.Type, e.State)
	default:
		log.Printf("no events")
	}

	for {
		// 监听节点
		children, _, chanEvents, err := conn.ChildrenW("/child")
		if err != nil {
			log.Panic(err)
		}
		log.Println(children)

		select {
		case event := <-chanEvents:
			switch event.Type {
			// 当前监听节点[/child]被删除时
			case zk.EventNodeDeleted:
				log.Printf("[delete] node path : %s", event.Path)
			// 当前监听节点[/child]创建时
			case zk.EventNodeCreated:
				log.Printf("[create] node path : %s", event.Path)
			// 当前监听节点[/child]数据改变时
			case zk.EventNodeDataChanged:
				log.Printf("[DataChange] node path : %s", event.Path)
			// 当前监听节点子节点[/child/sub]变更时
			case zk.EventNodeChildrenChanged:
				log.Printf("[ChildrenChange] node path : %s", event.Path)
			default:
				log.Println(event)
			}
		}
	}
}
