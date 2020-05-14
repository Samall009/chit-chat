package handers

import (
	"net/http"
	"test/blog/models"
)

// 首页处理器
func Index(Writer http.ResponseWriter, Request *http.Request) {
	// 数据获取
	threads, err := models.Threads()

	// 获取用户信息
	_, err = session(Writer, Request)

	// 判断是否登陆
	if err == nil {
		// 数据赋值
		generateHTML(Writer, threads, "layout", "auth.navbar", "index")
	} else {
		generateHTML(Writer, threads, "layout", "navbar", "index")
	}
}

// 全局错误处理
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}
