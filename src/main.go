package main

import (
	"otp/src/controller/httpserver"
	otphandler "otp/src/controller/httpserver/otpHandler"
	"otp/src/pkg/config"
	"otp/src/repo/adaptor"
	"otp/src/repo/implementation"
	"otp/src/service"
)

func main() {
	config.GetAppConfigInstance()
	redisClient := adaptor.CreateRedisClient()
	otpRepo := implementation.NewRedisOTPRepository(redisClient)
	otpSvc := service.GetInstanceOfOTPService(otpRepo)
	otpHandler := otphandler.New(*otpSvc)
	server := httpserver.New(otpHandler)
	server.Serve()
}
