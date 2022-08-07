package rest

import "github.com/gin-gonic/gin"

// RunAPIWithHandler func
func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	// get products
	r.GET("/products", h.GetProducts)
	// get promos
	r.GET("/promos", h.GetPromos)

	// r.POST("/user/:id/signout", h.SignOut)
	// r.GET("/user/:id/orders", h.GetOrders)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders")
	}

	// r.POST("/users/charge", h.Charge)
	// r.POST("/users/signin", h.SignIn)
	// r.POST("/users", h.AddUser)
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	// run the server
	return r.Run(address)
}

// RunAPI func
func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}
