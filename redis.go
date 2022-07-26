
package main
 
import (
    "fmt"
    "github.com/garyburd/redigo/redis"
 
)
 
func main() {
    c , err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println(err)
	}
    defer c.Close()
    if _, err := c.Do("AUTH", "123456"); err != nil {
		c.Close()
	}
    _,err = c.Do("Set","key1",998)   //redis写入数据
	if err != nil{
		fmt.Println(err)
	}
    r ,err := redis.Int(c.Do("Get","key1"))  //类型断言接口
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(r)
    //修改为HSet 定义Hash值
	_,err = c.Do("HSet","user01","name","john")
	if err != nil{
		fmt.Println(err)
	}
 
    //定义年龄
	_,err = c.Do("HSet","user01","age",18)
	if err != nil{
		fmt.Println(err)
	}
    _,err=c.Do("HMSet","user02","name","john","age",18)
    if err!=nil{
        fmt.Println(err)
    }
    res1, err := redis.String(c.Do("Get", "AUTH"))
	if err != nil {
        fmt.Println("get string failed,", err)
        return
    }
    fmt.Println(res1)


    

}
