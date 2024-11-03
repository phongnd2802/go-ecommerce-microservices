package repo

import "context"

type RegisterTxParams struct {
	CreateUserVerifyParams
	AfterCreate []func(UserUserVerify) error
}

type RegisterTxResult struct {
	UserVerify UserUserVerify
}

func (store *sqlStore) RegisterTx(ctx context.Context, arg RegisterTxParams) (RegisterTxResult, error) {
	var result RegisterTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.UserVerify, err = q.CreateUserVerify(ctx, arg.CreateUserVerifyParams)
		if err != nil {
			return err
		}
		for _, callback := range arg.AfterCreate {
			if err := callback(result.UserVerify); err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}