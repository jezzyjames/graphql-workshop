package schema

import (
	"database/sql"
	"grapql-api/pkg/data/repository"
	"grapql-api/pkg/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

func NewSchema(db *sql.DB) (graphql.Schema, error) {
	contactRepo := repository.NewContactRepository(db)
	rootQuery := resolvers.NewRootQuery(contactRepo)

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
}
