package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func main() {
	input := "test1"

	switch input {
	case "test1":
		zkTest1()
	case "test2":
		zkTest2()
	case "test3":
		zkTest3()
	default:
		fmt.Println("需要加参数【test1 | test2 | test3】")
	}
}

/**
 * 获取一个zk连接
 * @return {[type]}
 */
func zkGetConnect(zkList []string) (conn *zk.Conn) {
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

/**
 * 测试连接
 * @return
 */
func zkTest1() {
	zkList := []string{"localhost:2183"}
	conn := zkGetConnect(zkList)

	defer conn.Close()
	conn.Create("/go_servers", nil, 0, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

/**
 * 测试临时节点
 * @return {[type]}
 */
func zkTest2() {
	zkList := []string{"localhost:2183"}
	conn := zkGetConnect(zkList)

	defer conn.Close()
	conn.Create("/testadaadsasdsaw", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

/**
 * 获取所有节点
 */
func zkTest3() {
	zkList := []string{"localhost:2183"}
	conn := zkGetConnect(zkList)

	defer conn.Close()

	children, _, err := conn.Children("/go_servers")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", children)
}
