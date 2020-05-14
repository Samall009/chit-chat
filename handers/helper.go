package handers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"test/blog/models"
)

var logger *log.Logger

// 初始化定义
// 日志记录
func init() {
	file, err := os.OpenFile("logs/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// 为什么不命名为 error？避免和 error 类型重名
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// 异常处理统一重定向到错误页面
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func error_redirect() {

}

// 通过 Cookie 判断用户是否已登录
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	// 获取请求数据中的cookie
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		// 数据结构
		sess = models.Session{Uuid: cookie.Value}
		// 判断cookie是否可用
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 解析 HTML 模板（应对需要传入多个模板文件的情况，避免重复编写模板代码）
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	// 新建模板
	t = template.New("layout")
	// 获取模板对象文件位置
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	// 渲染模板
	t = template.Must(t.ParseFiles(files...))
	return
}

// 生成响应 HTML
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	// 文件位置
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	// 渲染模板
	templates := template.Must(template.ParseFiles(files...))
	// 绑定模板数据
	templates.ExecuteTemplate(writer, "layout", data)
}

// 返回版本号
func Version() string {
	return "0.1"
}
