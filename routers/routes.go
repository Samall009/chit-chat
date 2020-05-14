package routers

import (
	"net/http"
	"test/blog/handers"
)

// 基础路由结构体
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

// 路由列表
var webRoutes = WebRoutes{
	{
		"Index",
		"GET",
		"/",
		handers.Index,
	},
	{
		"signup",
		"GET",
		"/signup",
		handers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handers.SignupAccount,
	},
	{
		"login",
		"GET",
		"/login",
		handers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handers.Logout,
	},
	{
		"newThread",
		"GET",
		"/thread/new",
		handers.NewThread,
	},
	{
		"createThread",
		"POST",
		"/thread/create",
		handers.CreateThread,
	},
	{
		"readThread",
		"GET",
		"/thread/read",
		handers.ReadThread,
	},
	{
		"postThread",
		"POST",
		"/thread/post",
		handers.PostThread,
	},
	{
		"error",
		"GET",
		"/err",
		handers.Err,
	},
}
