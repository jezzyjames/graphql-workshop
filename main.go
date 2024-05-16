package main

import (
	"database/sql"
	"fmt"
	"grapql-api/pkg/data/models"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	_ "github.com/mattn/go-sqlite3"
)

func createRootQuery(db *sql.DB) *graphql.Object {

	contactField := graphql.NewObject(graphql.ObjectConfig{
		Name: "Contact",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.Int},
			"name":      &graphql.Field{Type: graphql.String},
			"firstName": &graphql.Field{Type: graphql.String},
			"lastName":  &graphql.Field{Type: graphql.String},
			"genderId":  &graphql.Field{Type: graphql.Int},
			"dob":       &graphql.Field{Type: graphql.String},
			"email":     &graphql.Field{Type: graphql.String},
			"phone":     &graphql.Field{Type: graphql.String},
			"address":   &graphql.Field{Type: graphql.String}}})

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				// Type can be a structure
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello, GraphQL!", nil
				},
			},
			"contact": &graphql.Field{
				Type: graphql.NewList(contactField),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, _ := db.Query("SELECT * FROM contact")
					var contacts []models.Contact
					for rows.Next() {
						var contact models.Contact
						err := rows.Scan(&contact.ID, &contact.Name, &contact.FirstName, &contact.LastName, &contact.GenderID, &contact.DOB, &contact.Email, &contact.Phone, &contact.Address, &contact.PhotoPath, &contact.CreatedAt, &contact.CreatedBy)

						if err != nil {
							return nil, err
						}
						contacts = append(contacts, contact)
					}

					return contacts, nil
				},
			},
		},
	})

	return rootQuery
}

func main() {
	// Create DB
	os.Remove("./contact.db")

	db, err := sql.Open("sqlite3", "./contact.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// End Create Database

	// Define GraphQL schema
	rootQuery := createRootQuery(db)
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		// Get
		Query: rootQuery,
		// Insert, Update, Delete
		// Mutation: ,
		// Subsription: ,
	})

	// Create a GraphQL handler for HTTP requests
	graphqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
	})

	// Serve GraphQL API at /graphql endpoint
	http.Handle("/graphql", graphqlHandler)

	// Start the HTTP server
	fmt.Println("Server is running at http://localhost:4000/graphql")
	http.ListenAndServe(":4000", nil)
}

// type Contact struct {
// 	ID        int64       `db:"contact_id" json:"contactId"`
// 	Name      string      `db:"name" json:"name"`
// 	FirstName string      `db:"first_name" json:"firstName"`
// 	LastName  string      `db:"last_name" json:"lastName"`
// 	GenderID  int         `db:"gender_id" json:"genderId"`
// 	DOB       interface{} `db:"dob" json:"dob"`
// 	Email     string      `db:"email" json:"email"`
// 	Phone     string      `db:"phone" json:"phone"`
// 	Address   string      `db:"address" json:"address"`
// 	PhotoPath string      `db:"photo_path" json:"photoPath"`
// 	CreatedAt string      `db:"created_at" json:"createdAt"`
// 	CreatedBy string      `db:"created_by" json:"createdBy"`
// }
