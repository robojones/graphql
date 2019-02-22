package mutation

import (
	"context"
	"github.com/robojones/graphql/prisma"
	"github.com/robojones/graphql/server/session_context"
	"github.com/robojones/graphql/server/session_cookie"
)

func (m *Mutation) Logout(ctx context.Context) (prisma.User, error) {

	session_cookie.Unset(ctx)

	token, err := session_context.Token(ctx)

	if err != nil {
		return prisma.User{}, err
	}

	_, err = m.Prisma.DeleteSession(prisma.SessionWhereUniqueInput{
		Token: &token,
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	user, err := session_context.User(ctx)

	// needed because the return type cannot be nil
	if err != nil {
		return prisma.User{}, err
	}

	return *user, err
}
