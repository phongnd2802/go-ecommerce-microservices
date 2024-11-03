package app

import (
	"context"

	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/errs"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	violations := validateRegisterRequest(req)
	if violations != nil {
		return nil, errs.InvalidArgumentError(violations)
	}

	result, err := server.ua.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	rsp := &dto.RegisterResponse{
		VerifyId: result.VerifyID,
	}
	return rsp, nil
}


func validateRegisterRequest(req *dto.RegisterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateEmail(req.GetVerifyKey()); err != nil {
		violations = append(violations, errs.FieldViolation("verify_key", err))
	}
	return violations
}
