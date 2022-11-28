package main

import (
	"github.com/gin-gonic/gin"
)

/*
	func sayHello(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadFile("./hello.txt")
		_, _ = fmt.Fprintln(w,string(b))
	}

	func main() {
		http.HandleFunc("/hello",sayHello)
		err := http.ListenAndServe(":9090",nil)
		if err != nil{
			 fmt.Printf("http server failed, err:%v\n",err)
			return
		}
	}
*/
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}

func main() {
	r := gin.Default() //返回默认的路由引擎

	// 指定用户使用GET请求访问/hello时，执行sayHello这个函数
	r.GET("/hello", sayHello)

	//启动服务
	r.Run(":9090")
}
