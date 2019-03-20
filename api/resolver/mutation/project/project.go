package project

import (
	"github.com/robojones/graphql/prisma"
)

type Project struct {
	Prisma *prisma.Client
}

func New(prisma *prisma.Client) *Project {
	return &Project{
		Prisma: prisma,
	}
}
