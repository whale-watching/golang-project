package dblayer

import (
	"errors"
	"gelato/gin/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// DBORM struct
type DBORM struct {
	*gorm.DB
}

// NewORM func
func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

// GetAllProducts func
func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

// GetPromos func
func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

// GetCustomerByName func
func (db *DBORM) GetCustomerByName(firstname, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}

// GetCustomerByID func
func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

// GetProduct func
func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

// AddUser func
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	// hash the password
	hashPassword(&customer.Password)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	customer.Password = ""
	return customer, err
}

// SignInUser func
func (db *DBORM) SignInUser(email, passowrd string) (customer models.Customer, err error) {

	// find customer row
	result := db.Where(&models.Customer{Email: email}) // TODO check if db.Table("customers") is needed
	// chain to collect customer
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}

	// compare password
	if !checkPassword(customer.Password, passowrd) {
		return customer, ErrINVALIDPASSWORD
	}

	// update loggedin
	err = result.Update("loggedin", true).Error

	if err != nil {
		return customer, err
	}

	// query database with new customer and return the result
	err = db.Find(&customer).Error
	// clear password as it is not needed again
	customer.Password = ""
	return customer, err
}

// SignOutUserByID func
func (db *DBORM) SignOutUserByID(id int) error {
	// create customer
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	// update customer
	return db.Where(&customer).Update("loggedin", false).Error
}

// GetCustomerOrderByID func
func (db *DBORM) GetCustomerOrderByID(id int) (orders []models.Order, err error) {
	// TODO check err = db.Find(&orders, models.Order{CustomerID: id}).Error
	qry := db.Table("orders").Select("*")
	qry = qry.Joins("join customer on customres.id = customer_id")
	qry = qry.Joins("join products on products.id = product_id")
	qry = qry.Where("customer_id=?", id)
	return orders, qry.Scan(&orders).Error
	// return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id=?", id).Scan(&orders).Error
}

//hashPassword func
func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}

	// convert password string to byte slice so that we can use it with the bcrypt package
	sBytes := []byte(*s)

	// obtain hashed password via the GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// update pasword string with hashed version
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingHash)) == nil
}
