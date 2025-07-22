package otphandler

import (
	"net/http"
	dto "otp/src/controller/httpserver/otpHandler/DTO"
	errutils "otp/src/pkg/errUtils"

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

	// TODO: other functionalities of context like timeouts better to be implemented.
	access_token, err := h.otpSvc.VerifyOTP(c.Request().Context(), verifyReq)
	if err != nil {

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
