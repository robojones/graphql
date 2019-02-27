package auth

import (
	"context"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/lib/auth"
	"github.com/robojones/graphql/lib/session_cookie"
	"github.com/robojones/graphql/prisma"
	"github.com/vektah/gqlparser/gqlerror"
	"golang.org/x/crypto/bcrypt"
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

func (a *Auth) Login(ctx context.Context, email string, password string) (gqlgen.LoginResult, error) {
	user, err := a.Prisma.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err == prisma.ErrNoResult {
		return gqlgen.LoginResult{}, UserNotFoundError
	}

	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return gqlgen.LoginResult{}, IncorrectPasswordError
	}

	session, err := a.Prisma.CreateSession(prisma.SessionCreateInput{
		User: prisma.UserCreateOneWithoutSessionsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &user.ID,
			},
		},
		Token: auth.GenerateToken(),
	}).Exec(ctx)

	session_cookie.Set(ctx, session)

	return gqlgen.LoginResult{
		Session: *session,
		User:    *user,
	}, nil
}
