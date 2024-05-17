package resolvers

import (
	"grapql-api/pkg/data/repository"

	"github.com/graphql-go/graphql"
)

func NewMutationQuery(transactionRepo *repository.TransactionRepo) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createDepositTransaction": &graphql.Field{
				Type: graphql.Int,
				Args: graphql.FieldConfigArgument{
					"transaction": &graphql.ArgumentConfig{Type: graphql.NewInputObject(graphql.InputObjectConfig{
						Name: "TransactionInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"bankAccountId": &graphql.InputObjectFieldConfig{Type: graphql.String},
							"amount":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
							"status":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
						},
					})},
				},
				Resolve: transactionRepo.InsertTransaction,
			},
		},
	})
}
