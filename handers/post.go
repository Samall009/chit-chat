package handers

import (
	"fmt"
	"net/http"
	"test/blog/models"
)

// POST /thread/post
// 在指定群组下创建新主题
func PostThread(writer http.ResponseWriter, request *http.Request) {
	// 判断是否登陆
	sess, err := session(writer, request)
	if err != nil {
		// 重定向登录页
		http.Redirect(writer, request, "/login", 302)
	} else {
		// 表单数据
		err = request.ParseForm()
		if err != nil {
			fmt.Println("Cannot parse form")
		}
		// 获取用户
		user, err := sess.User()
		if err != nil {

			fmt.Println("Cannot get user from session")
		}
		// 获取请求数据
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		// 根据uuid获取群组
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			error_message(writer, request, "Cannot read thread")
		}
		// 创建聊天
		if _, err := user.CreatePost(thread, body); err != nil {
			fmt.Println("Cannot create post")
		}
		// 重定向至群组
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
