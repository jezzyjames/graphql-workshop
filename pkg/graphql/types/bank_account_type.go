package types

import "github.com/graphql-go/graphql"

var BankAccountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BankAccount",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"accountId": &graphql.Field{Type: graphql.String},
		"name":      &graphql.Field{Type: graphql.String},
		"balance":   &graphql.Field{Type: graphql.Int},
	},
})
