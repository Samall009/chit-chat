package main

import (
	"log"
	"net/http"
	. "test/blog/routers"
)

func main() {
	startWebServer("2222")
}

// 服务启动函数
func startWebServer(port string) {
	// 创建路由
	r := NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 通过 router.go 中定义的路由器来分发请求

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // 启动协程监听请求

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
