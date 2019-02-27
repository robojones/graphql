package auth

import (
	"context"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/prisma"
	"github.com/vektah/gqlparser/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

const duplicateEmailErrorMessage = "graphql: A unique constraint would be violated on User. Details: Field name = email"

var DuplicateEmailError = &gqlerror.Error{
	Message: "Email already used for another account",
	Extensions: map[string]interface{}{
		"type": "Auth",
		"code": "DuplicateEmail",
	},
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func (a *Auth) Signup(ctx context.Context, email string, name string, password string) (gqlgen.LoginResult, error) {
	_, err := a.Prisma.CreateUser(prisma.UserCreateInput{
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

	return a.Login(ctx, email, password)
}
