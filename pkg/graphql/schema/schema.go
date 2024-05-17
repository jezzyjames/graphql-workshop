package schema

import (
	"database/sql"
	"grapql-api/pkg/data/repository"
	"grapql-api/pkg/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

func NewSchema(db *sql.DB) (graphql.Schema, error) {
	contactRepo := repository.NewContactRepo(db)
	transactionRepo := repository.NewTransactionRepo(db)
	rootQuery := resolvers.NewRootQuery(contactRepo)
	mutationQuery := resolvers.NewMutationQuery(transactionRepo)

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: mutationQuery,
	})
}
