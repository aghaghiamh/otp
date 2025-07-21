package otphandler

import (
	"net/http"
	"otp/src/controller/httpserver/otpHandler/DTO"

	"github.com/labstack/echo/v4"
)

func (h Handler) VerifyOTP(c echo.Context) error {
	var req dto.VerifyOTPInputDTO
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: Validation

	_, err := h.otpSvc.VerifyOTP(*req.MobileNumber, *req.OtpCode)
	if err != nil {
		// TODO: ERROR MSG

		return echo.NewHTTPError(http.StatusBadRequest, "Not proper")
	}

	return c.JSON(http.StatusOK, "resp")
}
