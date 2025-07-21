package otphandler

import (
	"fmt"
	"net/http"
	dto "otp/src/controller/httpserver/otpHandler/DTO"
	"otp/src/pkg/errUtils"

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
//	@Param			payload	body		dto.RequestOTPInput	true	"Request OTP Payload"
//	@Success		200		{object}	dto.RequestOTPOutput
//	@Router			/user-management/req-otp [post]
func (h Handler) RequestOTP(c echo.Context) error {
	var req dto.RequestOTPInput
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: Validation

	err := h.otpSvc.RequestOTP(*req.MobileNumber)
	if err != nil {

		return echo.NewHTTPError(errutils.GetStatusCode(err), errutils.GenerateErrorMessage(err))
	}

	return c.JSON(http.StatusOK, dto.RequestOTPOutput{
		Message: fmt.Sprintf("The OTP code have been sent for %s mobile number.", *req.MobileNumber),
	})
}
