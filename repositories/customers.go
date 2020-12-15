package repositories

import (
	"github.com/edwinvautier/go-boilerplate/database"
	"github.com/edwinvautier/go-boilerplate/models"
)

// PersistCustomer is used to persist user objects to database
func PersistCustomer(customer *models.Customer) error {
	var err error

	err = database.Db.Debug().Create(&customer).Error
	if err != nil {
		return err
	}

	return nil
}