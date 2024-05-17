package main

import (
	"database/sql"
	"fmt"
	"grapql-api/pkg/auth"
	"grapql-api/pkg/graphql/schema"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database
	os.Remove("./contact.db")

	db, err := sql.Open("sqlite3", "./contact.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the GraphQL schema
	graphqlSchema, err := schema.NewSchema(db)
	if err != nil {
		log.Fatal(err)
	}

	// Create a GraphQL handler for HTTP requests
	graphqlHandler := handler.New(&handler.Config{
		Schema:     &graphqlSchema,
		Pretty:     true,
		Playground: true,
	})

	// Serve GraphQL API at /graphql endpoint
	http.Handle("/graphql", auth.AuthenticationHandler(graphqlHandler))
	http.HandleFunc("/login", auth.LoginHandler)

	// Start the HTTP server
	fmt.Println("Server is running at http://localhost:4002/graphql")
	http.ListenAndServe(":4002", nil)
}

// INSERT
// mutation {
// 	createDepositTransaction(  transaction: {
// 	   status: 1,
// 	   bankAccountId: "test",
// 	   amount: 200,
// 	 })
//    }
