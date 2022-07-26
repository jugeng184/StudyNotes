package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当程序启动的时候，就初始化链接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {

	conn := pool.Get() //先从pool 取出一个链接
	defer conn.Close() //关闭这个链接，用完了就不要占着位置

	// //添加一个认证
	// if _, err := conn.Do("AUTH", "123456"); err != nil {
	// 	conn.Close()
	// }

	if _, err := conn.Do("SET", "name", "tom"); err!=nil{
		conn.Close()
	}
	

	//取数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)
}
