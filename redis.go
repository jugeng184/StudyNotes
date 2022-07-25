
package main
 
import (
    "fmt"
 
)
 
func main() {
    client, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("redis connect failed,", err)
        return
    } 
 
    fmt.Println("redis connect success")
	_, err = client.Do("Set", "abc", 100)
    if err != nil {
        fmt.Println("set string failed", err)
        return
    }
    _, err = client.Do("Set", "36D", "good")
    if err != nil {
        fmt.Println("set string failed", err)
        return
    }
	res, err := redis.Int(client.Do("Get", "abc"))
    if err != nil {
        fmt.Println("get string failed,", err)
        return
    }
    fmt.Println(res)
    // redigo 通过redis.String()函数来获取字符串
    res1, err := redis.String(client.Do("Get", "36D"))
    if err != nil {
        fmt.Println("get string failed,", err)
        return
    }
    fmt.Println(res1)
	_, err = client.Do("lpush", "NBAplayer", "Jordon", "Kobe", "Lebron")
    if err != nil {
        fmt.Println("push element failed")
        return
    }

 
    defer client.Close()
}
