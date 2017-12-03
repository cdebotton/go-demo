package schema

import (
	graphql "github.com/neelance/graphql-go"
)

type user struct {
	ID       graphql.ID
	Email    string
	Username string
	Password string
}

var users = []*user{
	{
		ID:       "1",
		Email:    "admin@obscura.dev",
		Username: "admin",
		Password: "root",
	},
	{
		ID:       "2",
		Email:    "editor@obscura.dev",
		Username: "editor",
		Password: "root",
	},
	{
		ID:       "3",
		Email:    "user@obscura.dev",
		Username: "user",
		Password: "root",
	},
	{
		ID:       "4",
		Email:    "demo@obscura.dev",
		Username: "demo",
		Password: "root",
	},
}

type userResolver struct {
	u *user
}

func (r *userResolver) ID() graphql.ID {
	return r.u.ID
}

func (r *userResolver) Email() string {
	return r.u.Email
}

func (r *userResolver) Username() string {
	return r.u.Username
}

func (r *userResolver) Password() string {
	return r.u.Password
}

// Users is the resolver for all user data
func (r *Resolver) Users() []*userResolver {
	var list []*userResolver
	for _, user := range users {
		list = append(list, &userResolver{user})
	}
	return list
}
