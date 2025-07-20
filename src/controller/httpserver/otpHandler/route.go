package otphandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/user-management")

	userGroup.POST("/req-otp", h.RequestOTP)
}
