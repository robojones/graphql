package mutation

import (
	"context"
	"github.com/robojones/graphql/prisma"
	"github.com/robojones/graphql/server/auth"
)

func (m *Mutation) LoginUser(ctx context.Context, email string, password string) (prisma.User, error) {
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

	auth.ApplySession(ctx, session)

	return *user, nil
}
