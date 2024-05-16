package types

import "github.com/graphql-go/graphql"

var TransactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Transaction",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.Int},
		"bankAccountId": &graphql.Field{Type: graphql.String},
		"amount":        &graphql.Field{Type: graphql.Int},
		"status":        &graphql.Field{Type: graphql.Int},
		"createdAt":     &graphql.Field{Type: graphql.DateTime},
	},
})
