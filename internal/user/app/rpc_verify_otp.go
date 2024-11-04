package app

import (
	"context"

	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/errs"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	VerifyOTPRegister = 0
	VerifyOTPResetPassword = 1
)


func (server *Server) VerifyOTP(ctx context.Context, req *dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, error) {
	violations := validateVerifyOTPRequest(req)
	if violations != nil {
		return nil, errs.InvalidArgumentError(violations)
	}

	result, err := server.ua.VerifyOTP(ctx, req)
	if err != nil {
		return nil, err
	}
	
	rsp := &dto.VerifyOTPResponse{
		Token: result.VerifyKeyHash,
	}
	return rsp, nil
}


func validateVerifyOTPRequest(req *dto.VerifyOTPRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateEmail(req.GetVerifyKey()); err != nil {
		violations = append(violations, errs.FieldViolation("verify_key", err))
	}

	if err := validator.ValidateString(req.GetVerifyCode(), 6, 6); err != nil {
		violations = append(violations, errs.FieldViolation("verify_code", err))
	}

	return violations
}