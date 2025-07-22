package otphandler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	dto "otp/src/controller/httpserver/otpHandler/DTO"
	"otp/src/pkg/config"
	"otp/src/pkg/errUtils"
	"otp/src/pkg/log"
	"time"

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

	if fieldErrors, vErr := h.validator.ValidateRequestOTP(req); vErr != nil {
		return c.JSON(errutils.GetStatusCode(vErr), echo.Map{
			"message": errutils.GenerateErrorMessage(vErr),
			"errors":  fieldErrors,
		})
	}

	cnf := config.GetAppConfigInstance()
	ctx, cancel := context.WithTimeout(
		c.Request().Context(),
		time.Duration(cnf.RequestTimeoutInSeconds)*time.Second,
	)
	defer cancel()

	err := h.otpSvc.RequestOTP(ctx, req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.GetLoggerInstance().WithError(err).Error("The request timed out.")
			return echo.NewHTTPError(http.StatusGatewayTimeout, errutils.GenerateErrorMessage(err))
		}

		return echo.NewHTTPError(errutils.GetStatusCode(err), errutils.GenerateErrorMessage(err))
	}

	return c.JSON(http.StatusOK, dto.RequestOTPOutput{
		Message: fmt.Sprintf("The OTP code have been sent for %s mobile number.", *req.MobileNumber),
	})
}
