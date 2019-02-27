package mutation

import (
	"context"
	"github.com/pkg/errors"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/prisma"
	"golang.org/x/crypto/bcrypt"
)

const duplicateEmailErrorMessage = "graphql: A unique constraint would be violated on User. Details: Field name = email"

var DuplicateEmailError = errors.New("Email already used for another account")

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func (m *Mutation) Signup(ctx context.Context, email string, name string, password string) (gqlgen.LoginResult, error) {
	_, err := m.Prisma.CreateUser(prisma.UserCreateInput{
		Name:         name,
		Email:        email,
		PasswordHash: hashPassword(password),
	}).Exec(ctx)

	if err.Error() == duplicateEmailErrorMessage {
		return gqlgen.LoginResult{}, DuplicateEmailError
	}

	if err != nil {
		panic(err)
	}

	return m.Login(ctx, email, password)
}
