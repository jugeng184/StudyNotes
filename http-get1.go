package main
 
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main(){
    resp, err := http.Get("http://baidu.com/")
	if err != nil {
		// handle error
		}
		defer resp.Body.Close()   //关闭body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {   //处理错误
			// handle error
		}
		fmt.Println(string(body))
	}
