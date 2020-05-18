package handers

import (
	"net/http"
	"test/blog/models"
)

// GET /login
// 登录页面
func Login(Writer http.ResponseWriter, Request *http.Request) {
	// 获取模板渲染器指针
	t := parseTemplateFiles("auth.layout", "navbar", "login")
	// 渲染模板
	t.Execute(Writer, nil)
}

// GET /signup
// 注册页面
func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

// POST /signup
// 注册新用户
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	// 判断数据
	if err != nil {
		danger(err, "Cannot parse form")
	}
	// 实例化结构体
	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot parse form")
	}
	// 重定向至登陆页面
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// 通过邮箱和密码字段对用户进行认证
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	// 判断是否存在数据
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot find user")
	}
	// 获取用户信息
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	// 判断密码是否正确
	if user.Password == models.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		// 添加Cookie
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		// 写入Cookie
		http.SetCookie(writer, &cookie)
		// 重定向至首页
		http.Redirect(writer, request, "/", 302)
	} else {
		// 重定向至错误页面
		error_message(writer, request, "密码错误")
	}
}

// GET /logout
// 用户退出
func Logout(writer http.ResponseWriter, request *http.Request) {
	// 获取cookie信息
	cookie, err := request.Cookie("_cookie")
	// 判断cookie是否正确获取
	if err != http.ErrNoCookie {
		danger(err, "Failed to get cookie")
		// 获取session结构体
		session := models.Session{Uuid: cookie.Value}
		info(session, "session")
		// 删除会话
		err = session.DeleteByUUID()

		http.SetCookie(writer, &http.Cookie{})

	}
	// 重定向至首页
	http.Redirect(writer, request, "/", 302)
}
