package main

import (
	"os"
	"os/signal"
	_ "otp/docs"
	"otp/src/controller/httpserver"
	"otp/src/controller/httpserver/otpHandler"
	"otp/src/pkg/config"
	"otp/src/repo/adaptor"
	"otp/src/repo/implementation"
	"otp/src/service"
)

// @title			OTP UserManagement
// @version		1.0
// @description	This is a UserManagenet which implements OTP
// @termsOfService	http://swagger.io/terms/
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
