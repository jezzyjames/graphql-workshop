package resolvers

import (
	"grapql-api/pkg/data/repository"
	"grapql-api/pkg/graphql/types"

	"github.com/graphql-go/graphql"
)

func NewRootQuery(contactRepo *repository.ContactRepo) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello, GraphQL!", nil
				},
			},
			"contact": &graphql.Field{
				Type: graphql.NewList(types.ContactType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return contactRepo.GetAllContacts()
				},
			},
		},
	})
}
