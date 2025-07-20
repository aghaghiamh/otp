package httpserver

import (
	"context"
	"fmt"
	"log"
	otphandler "otp/src/controller/httpserver/otpHandler"
	"otp/src/pkg/config"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpConfig struct {
	Host                    string        `mapstructure:"host"`
	Port                    string        `mapstructure:"port"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type Server struct {
	router     *echo.Echo
	otpHandler otphandler.Handler
}

func New(otpHandler otphandler.Handler) Server {
	return Server{
		router:     echo.New(),
		otpHandler: otpHandler,
	}
}

func (s *Server) Serve() {
	s.router.Use(middleware.RequestID())
	s.router.Use(middleware.Recover())

	s.otpHandler.SetRoutes(s.router)

	cfg := config.GetAppConfigInstance()

	if err := s.router.Start(cfg.Server.Host + ":" + strconv.Itoa(int(cfg.Server.Port))); err != nil {
		log.Fatalf("Couldn't Listen to the %d port: %s", cfg.Server.Port, err.Error())
	}
}

func (s Server) Shutdown() {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("error while shutting down the server: ", err)
	}

	fmt.Println("Gracefully shutdowned!!")
	<-ctxWithTimeout.Done()
}
