package main

import (
	"otp/src/controller/httpserver"
	otphandler "otp/src/controller/httpserver/otpHandler"
	"otp/src/pkg/config"
	"otp/src/repo/adapter"
	"otp/src/service"
)

func main() {
	config.GetAppConfigInstance()
	otpRepo := adapter.GetReposInstances()
	otpSvc := service.GetInstanceOfOTPService(otpRepo)
	otpHandler := otphandler.New(*otpSvc)
	server := httpserver.New(otpHandler)
	server.Serve()
}
