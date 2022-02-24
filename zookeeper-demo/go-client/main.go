package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var (
	hosts       = []string{"127.0.0.1:2181"}
	path        = "/wtzk"
	flags int32 = zk.FlagEphemeral
	data        = []byte("zk data 001")
	acls        = zk.WorldACL(zk.PermAll)
)

func main() {
	// 创建监听的option，用于初始化zk
	eventCallbackOption := zk.WithEventCallback(callback)
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5, eventCallbackOption)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 开始监听path
	_, _, event, err := conn.ExistsW(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	var a zk.Event
	a = <-event
	callback(a)

	// 触发创建数据操作
	// create(conn, path, data)

}

// zk watch 回调函数
func callback(event zk.Event) {
	fmt.Println("path: ", event.Path)
	fmt.Println("type: ", event.Type.String())
	fmt.Println("state: ", event.State.String())
}

// 创建数据
func create(conn *zk.Conn, path string, data []byte) {
	_, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Printf("创建数据失败: %v\n", err)
		return
	}
	fmt.Println("创建数据成功")
}
