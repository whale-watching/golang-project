package dblayer

import (
	"errors"
	"gelato/gin/src/models"
)

// ErrINVALIDPASSWORD variable
var ErrINVALIDPASSWORD = errors.New("Invalid password")

// DBLayer interface
type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrderByID(int) ([]models.Order, error)
}
