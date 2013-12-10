// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tLogin struct {}
var Login tLogin


func (_ tLogin) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Login.Login", args).Url
}

func (_ tLogin) Enter(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Login.Enter", args).Url
}

func (_ tLogin) OAuth(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("Login.OAuth", args).Url
}


type tApplication struct {}
var Application tApplication


func (_ tApplication) Profile(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Profile", args).Url
}


type tWebSocket struct {}
var WebSocket tWebSocket


func (_ tWebSocket) Room(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("WebSocket.Room", args).Url
}

func (_ tWebSocket) RoomSocket(
		user string,
		ws interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "ws", ws)
	return revel.MainRouter.Reverse("WebSocket.RoomSocket", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


