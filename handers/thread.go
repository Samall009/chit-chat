package handers

import (
	"net/http"
	"test/blog/models"
)

// GET /threads/new
// 创建群组页面
func NewThread(writer http.ResponseWriter, request *http.Request) {
	// 判断用户是否登陆
	_, err := session(writer, request)
	if err != nil {
		// 重定向至登录页
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "new.thread")
	}
}

// POST /thread/create
// 执行群组创建逻辑
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	// 判断用户是否登陆
	sess, err := session(writer, request)
	if err != nil {
		// 重定向至登录页
		http.Redirect(writer, request, "/login", 302)
	} else {
		// 判断表单数据
		err = request.ParseForm()
		if err != nil {
			error_message(writer, request, "Cannot parse form")
		}
		// 获取登陆用户
		user, err := sess.User()
		if err != nil {
			error_message(writer, request, "Cannot get user from session")
		}
		// 获取值
		topic := request.PostFormValue("topic")
		// 创建新的群组
		if _, err := user.CreateThread(topic); err != nil {
			error_message(writer, request, "Cannot create thread")
		}
		// 重定向至首页
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// 通过 ID 渲染指定群组页面
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	// 获取参数列表
	vals := request.URL.Query()
	// 获取ID
	uuid := vals.Get("id")
	// 获取群组信息
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		error_message(writer, request, "Cannot read thread")
	} else {
		// 判断用户登陆
		_, err := session(writer, request)
		// 渲染模板
		if err != nil {
			generateHTML(writer, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
