package repo

import (
	"context"

	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
)

type UpdatePassswordParamsTx struct {
	CreateUserBaseParams
}

type UpdatePasswordResultTx struct {
	UserProfile UserUserProfile
}


func (store *sqlStore) UpdatePasswordRegisterTx(ctx context.Context, arg UpdatePassswordParamsTx) (UpdatePasswordResultTx, error) {
	var result UpdatePasswordResultTx

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		
		userBase, err := q.CreateUserBase(ctx, arg.CreateUserBaseParams)
		if err != nil {
			return err
		}

		result.UserProfile, err = q.CreateUserProfile(ctx, CreateUserProfileParams{
			UserID: userBase.UserID,
			UserEmail: userBase.UserEmail,
			UserNickname: user.GetNicknameFromEmail(userBase.UserEmail),
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}