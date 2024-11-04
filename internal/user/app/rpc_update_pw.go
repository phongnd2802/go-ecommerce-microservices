package app

import (
	"context"

	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/errs"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) UpdatePasswordRegister(ctx context.Context, req *dto.SetPasswordRequest) (*dto.SetPasswordResponse, error) {
	violations := validateSetPasswordRequest(req)
	if violations != nil {
		return nil, errs.InvalidArgumentError(violations)
	}

	result, err := server.ua.UpdatePasswordRegister(ctx, req)
	if err != nil {
		return nil, err
	}

	rsp := &dto.SetPasswordResponse{
		UserEmail: result.UserEmail,
		UserNickname: result.UserNickname,
	}

	return rsp, nil
}

func validateSetPasswordRequest(req *dto.SetPasswordRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateString(req.GetToken(), 20, 100); err != nil {
		violations = append(violations, errs.FieldViolation("token", err))
	}
	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, errs.FieldViolation("password", err))
	}

	return violations
}