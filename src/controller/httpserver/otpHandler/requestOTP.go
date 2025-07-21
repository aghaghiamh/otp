package otphandler

import (
	"net/http"
	dto "otp/src/controller/httpserver/otpHandler/DTO"

	_ "otp/src/docs"

	"github.com/labstack/echo/v4"
)

// RequestOTP Request an OTP
//
//	@Summary		Request an OTP
//	@Description	Request an OTP
//	@Tags			OTP
//	@Accept			json
//	@Produce		json
//	@Param			payload body dto.RequestOTPInputDTO true "Request OTP Payload"
//	@Success		200			string		model.Account
//	@Router			/user-management/req-otp [post]
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
