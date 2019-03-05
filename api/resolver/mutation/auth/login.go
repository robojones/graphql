package auth

import (
	"context"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/lib/auth"
	"github.com/robojones/graphql/prisma"
	"github.com/vektah/gqlparser/gqlerror"
)

var UserNotFoundError = &gqlerror.Error{
	Message: "user not found",
	Extensions: map[string]interface{}{
		"type": "Auth",
		"code": "UserNotFound",
	},
}
var IncorrectPasswordError = &gqlerror.Error{
	Message: "password is incorrect",
	Extensions: map[string]interface{}{
		"type": "Auth",
		"code": "PasswordIncorrect",
	},
}

func (a *Auth) Login(ctx context.Context, email string, password string) (*gqlgen.LoginResult, error) {
	user, err := a.Prisma.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err == prisma.ErrNoResult {
		return nil, UserNotFoundError
	}

	if err != nil {
		panic(err)
	}

	err = auth.VerifyPassword(user.PasswordHash, password)
	if err != nil {
		return nil, IncorrectPasswordError
	}

	session, err := a.Prisma.CreateSession(prisma.SessionCreateInput{
		User: prisma.UserCreateOneWithoutSessionsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &user.ID,
			},
		},
		Token: auth.GenerateToken(),
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	auth.SetCookie(ctx, session)

	return &gqlgen.LoginResult{
		Session: *session,
		User:    *user,
	}, nil
}
