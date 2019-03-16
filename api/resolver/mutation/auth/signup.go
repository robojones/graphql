package auth

import (
	"context"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/lib/auth"
	"github.com/robojones/graphql/prisma"
	"github.com/vektah/gqlparser/gqlerror"
)

const duplicateEmailErrorMessage = "graphql: A unique constraint would be violated on User. Details: Field name = email"

var DuplicateEmailError = &gqlerror.Error{
	Message: "Email already used for another account",
	Extensions: map[string]interface{}{
		"type": "Auth",
		"code": "DuplicateEmail",
	},
}

func (a *Auth) Signup(ctx context.Context, params gqlgen.SignupInput) (*gqlgen.SignupPayload, error) {
	_, err := a.Prisma.CreateUser(prisma.UserCreateInput{
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: auth.HashPassword(params.Password),
	}).Exec(ctx)

	if err != nil {
		if err.Error() == duplicateEmailErrorMessage {
			return nil, DuplicateEmailError
		}

		panic(err)
	}

	payload, err := a.Login(ctx, gqlgen.LoginInput{
		Email: params.Email,
		Password: params.Password,
	})

	return (*gqlgen.SignupPayload)(payload), err
}
