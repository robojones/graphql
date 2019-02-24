package mutation

import (
	"context"
	"github.com/pkg/errors"
	"github.com/robojones/graphql/lib/auth"
	"github.com/robojones/graphql/lib/session_cookie"
	"github.com/robojones/graphql/prisma"
	"golang.org/x/crypto/bcrypt"
)

var UserNotFoundError = errors.New("user not found")
var IncorrectPasswordError = errors.New("password is incorrect")

func (m *Mutation) Login(ctx context.Context, email string, password string) (prisma.User, error) {
	user, err := m.Prisma.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err != nil {
		return prisma.User{}, err
	}

	if user == nil {
		return prisma.User{}, UserNotFoundError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return prisma.User{}, IncorrectPasswordError
	}

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
