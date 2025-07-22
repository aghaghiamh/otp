package otphandler

import (
	"context"
	"errors"
	"net/http"
	dto "otp/src/controller/httpserver/otpHandler/DTO"
	"otp/src/pkg/config"
	errutils "otp/src/pkg/errUtils"
	"otp/src/pkg/log"
	"time"

	"github.com/labstack/echo/v4"
)

// VerifyOTP Verify an OTP
//
//	@Summary		Verify an OTP
//	@Description	Verify an OTP
//	@Tags			OTP
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	dto.VerifyOTPInput	true	"Verify OTP Payload"
//	@Success		200		string	dto.VerifyOTPOutput
//	@Router			/user-management/verify-otp [post]
func (h Handler) VerifyOTP(c echo.Context) error {
	var verifyReq dto.VerifyOTPInput
	if err := c.Bind(&verifyReq); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: Validation is a must here, to pass it to the service layer

	cnf := config.GetAppConfigInstance()
	ctx, cancel := context.WithTimeout(
		c.Request().Context(),
		time.Duration(cnf.RequestTimeoutInSeconds)*time.Second,
	)
	defer cancel()

	access_token, err := h.otpSvc.VerifyOTP(ctx, verifyReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.GetLoggerInstance().WithError(err).Error("The request timed out.")
			return echo.NewHTTPError(http.StatusGatewayTimeout, errutils.GenerateErrorMessage(err))
		}

		return echo.NewHTTPError(errutils.GetStatusCode(err), errutils.GenerateErrorMessage(err))
	}

	return c.JSON(http.StatusOK, dto.VerifyOTPOutput{
		AuthTokens: dto.AuthTokens{
			AccessToken: access_token,
		},
		UserInfo: dto.UserOTPResponseInfo{
			MobileNumber: *verifyReq.MobileNumber,
		},
	})
}
