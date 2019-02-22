package mutation

import (
	"context"
	"github.com/robojones/graphql/lib/auth"
	"github.com/robojones/graphql/lib/session_cookie"
	"github.com/robojones/graphql/prisma"
)

func (m *Mutation) Login(ctx context.Context, email string, password string) (prisma.User, error) {
	user, err := m.Prisma.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err != nil {
		return *user, err
	}

	// TODO: verify password
	// TODO: handle user == nil

	session, err := m.Prisma.CreateSession(prisma.SessionCreateInput{
		User: prisma.UserCreateOneWithoutSessionsInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &user.ID,
			},
		},
		Token: auth.GenerateToken(),
	}).Exec(ctx)

	session_cookie.Set(ctx, session)

	return *user, nil
}
