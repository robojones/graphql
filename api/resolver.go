package api

import (
	"context"
	"github.com/prisma/prisma-client-lib-go"
	"github.com/robojones/graphql/gqlgen"
	"github.com/robojones/graphql/prisma-client"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() gqlgen.PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() gqlgen.UserResolver {
	return &userResolver{r}
}
func (r *Resolver) Comment() gqlgen.CommentResolver {
	return &commentResolver{r}
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) User(ctx context.Context, obj *prisma.Comment) (prisma.User, error) {
	user, err := r.Prisma.Comment(prisma.CommentWhereUniqueInput{ID: &obj.ID}).User().Exec(ctx)
	return *user, err
}

func (r *commentResolver) Post(ctx context.Context, obj *prisma.Comment) (prisma.Post, error) {
	post, err := r.Prisma.Comment(prisma.CommentWhereUniqueInput{ID: &obj.ID}).Post().Exec(ctx)
	return *post, err
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SignupUser(ctx context.Context, email string, name string) (prisma.User, error) {
	user, err := r.Prisma.CreateUser(prisma.UserCreateInput{
		Email: email,
		Name:  name,
	}).Exec(ctx)
	return *user, err
}
func (r *mutationResolver) CreateDraft(ctx context.Context, title string, content string, authorEmail string) (prisma.Post, error) {
	post, err := r.Prisma.CreatePost(prisma.PostCreateInput{
		Title:   title,
		Content: content,
		User: prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{Email: &authorEmail},
		},
	}).Exec(ctx)
	return *post, err
}
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*prisma.Post, error) {
	return r.Prisma.DeletePost(prisma.PostWhereUniqueInput{ID: &id}).Exec(ctx)
}
func (r *mutationResolver) Publish(ctx context.Context, id string) (*prisma.Post, error) {
	published := true
	return r.Prisma.UpdatePost(prisma.PostUpdateParams{
		Where: prisma.PostWhereUniqueInput{ID: &id},
		Data:  prisma.PostUpdateInput{Published: &published},
	}).Exec(ctx)
}

type postResolver struct{ *Resolver }

func (r *postResolver) Comments(ctx context.Context, obj *prisma.Post) ([]prisma.Comment, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &obj.ID}).Comments(nil).Exec(ctx)
}

func (r *postResolver) User(ctx context.Context, obj *prisma.Post) (prisma.User, error) {
	author, err := r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &obj.ID}).User().Exec(ctx)
	return *author, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Feed(ctx context.Context) ([]prisma.Post, error) {
	published := true
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{Published: &published},
	}).Exec(ctx)
}
func (r *queryResolver) FilterPosts(ctx context.Context, searchString string) ([]prisma.Post, error) {
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{
			Or: []prisma.PostWhereInput{{
				TitleContains: &searchString,
			}, {
				TitleContains: &searchString,
			}},
		},
	}).Exec(ctx)
}
func (r *queryResolver) Post(ctx context.Context, id string) (*prisma.Post, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &id}).Exec(ctx)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Comments(ctx context.Context, obj *prisma.User) ([]prisma.Comment, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &obj.ID}).Comments(nil).Exec(ctx)
}

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &obj.ID}).Posts(nil).Exec(ctx)
}
