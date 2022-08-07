package rest

import (
	"gelato/gin/src/dblayer"
	"gelato/gin/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandlerInterface interface
type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

// Handler struct
type Handler struct {
	db dblayer.DBLayer
}

// NewHandler func
func NewHandler() (*Handler, error) {
	return new(Handler), nil
}

// GetProducts func
func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return
	}

	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetPromos func
func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {
		return
	}

	promos, err := h.db.GetPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promos)
}

// SignIn func
func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.SignInUser(customer.Email, customer.Password)
	if err != nil {
		if err == dblayer.ErrINVALIDPASSWORD {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// AddUser func
func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// SignOut func
func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {
		return
	}
	// get id parameter
	p := c.Param("id")
	// convert p from string to int
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.db.SignOutUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// GetOrders func
func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		return
	}
	// get id paramater
	p := c.Param("id")
	// convert p from string to int
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orders, err := h.db.GetCustomerOrderByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Charge credit card func
func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		return
	}
}
