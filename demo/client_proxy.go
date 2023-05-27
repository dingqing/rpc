package main

import (
	"context"
	"encoding/gob"
	"log"
	"reflect"

	"github.com/dingqing/rpc/consumer"
	"github.com/dingqing/rpc/proxy"
)

func main() {
	nodes := []string{"localhost:8881"}
	conf := &proxy.Config{Nodes: nodes, Env: "dev"}
	discovery := proxy.New(conf)

	gob.Register(User{})
	cli := consumer.NewClientProxy("UserService", consumer.DefaultOption, discovery)

	var GetUserById func(id int) (User, error)
	res, err := cli.Call(context.Background(), "User.GetUserById", &GetUserById, 1)
	if err != nil {
		log.Println("call error:", err)
	} else {
		val := res.([]reflect.Value)
		user := val[0].Interface().(User)
		log.Println("rpc return result:", user)
	}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
