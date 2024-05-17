package repository

import (
	"database/sql"
	"grapql-api/pkg/data/models"
)

type ContactRepo struct {
	DB *sql.DB
}

func NewContactRepo(db *sql.DB) *ContactRepo {
	return &ContactRepo{DB: db}
}

func (r *ContactRepo) GetAllContacts() ([]models.Contact, error) {
	rows, err := r.DB.Query("SELECT * FROM contact")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
}
