package otphandler

import (
	"net/http"
	"otp/src/controller/httpserver/otpHandler/DTO"

	"github.com/labstack/echo/v4"
)

func (h Handler) RequestOTP(c echo.Context) error {
	var req dto.RequestOTPInputDTO
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: Validation

	err := h.otpSvc.RequestOTP(*req.MobileNumber)
	if err != nil {
		// TODO: ERROR MSG

		return echo.NewHTTPError(http.StatusBadRequest, "Not proper")
	}

	return c.JSON(http.StatusOK, "resp")
}
