package otphandler

import (
	"net/http"
	dto "otp/src/controller/httpserver/otpHandler/DTO"

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
	var req dto.VerifyOTPInput
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: Validation

	access_token, err := h.otpSvc.VerifyOTP(*req.MobileNumber, *req.OtpCode)
	if err != nil {
		// TODO: ERROR MSG

		return echo.NewHTTPError(http.StatusBadRequest, "Not proper")
	}

	return c.JSON(http.StatusOK, dto.VerifyOTPOutput{
		AuthTokens: dto.AuthTokens{
			AccessToken: access_token,
		},
		UserInfo: dto.UserOTPResponseInfo{
			MobileNumber: *req.MobileNumber,
		},
	})
}
