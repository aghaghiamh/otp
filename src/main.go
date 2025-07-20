package main

import (
	"os"
	"os/signal"
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
	userManagementRepo := adaptor.GetRepoInstance()
	otpRepo := implementation.NewRedisOTPRepository(redisClient)
	otpSvc := service.GetInstanceOfOTPService(otpRepo, userManagementRepo)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	otpHandler := otphandler.New(*otpSvc)
	server := httpserver.New(otpHandler)
	go func() {
		server.Serve()
	}()

	// Graceful Termination
	<-quit
	server.Shutdown()
}
