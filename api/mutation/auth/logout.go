package auth

import (
	"context"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/lib/session_cookie"
	"github.com/robojones/graphql/prisma"
)

func (a *Auth) Logout(ctx context.Context) (gqlgen.LogoutResult, error) {

	session_cookie.Unset(ctx)

	token, err := session_context.Token(ctx)

	if err != nil {
		return gqlgen.LogoutResult{}, err
	}

	_, err = a.Prisma.DeleteSession(prisma.SessionWhereUniqueInput{
		Token: &token,
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	user, err := session_context.User(ctx)

	if err != nil {
		panic(err)
	}

	return gqlgen.LogoutResult{
		User: *user,
	}, err
}
