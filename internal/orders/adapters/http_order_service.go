package adapters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
)

// Exposes an OrderService using HTTP endpoints
type HttpOrderService struct {
	svc orders.OrderService
}

func NewHttpAdapter(svc orders.OrderService) *HttpOrderService {
	return &HttpOrderService{
		svc: svc,
	}
}

func (hos *HttpOrderService) Start() {
	r := gin.Default()
	r.GET("/orders", func(c *gin.Context) {
		res, err := hos.svc.GetOrders()
		if err != nil {
			fmt.Printf("Error getting orders: %v\n", err)
			http.Error(c.Writer, "Unknown error", http.StatusInternalServerError)
			return
		}

		c.JSON(200, res)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := hos.svc.GetOrder(id)
		if err != nil {
			fmt.Printf("Error getting order %s: %v\n", id, err)
			http.Error(c.Writer, "Unknown error", http.StatusInternalServerError)
			return
		}

		c.JSON(200, res)
	})

	r.POST("/orders", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v\n", err)
			http.Error(c.Writer, "Failed to read body", http.StatusBadRequest)
			return
		}

		dto := &struct {
			Product  string
			Quantity int
		}{}
		err = json.Unmarshal(body, dto)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			http.Error(c.Writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Creating order with product: %s, quantity: %d\n", dto.Product, dto.Quantity)
		order, err := hos.svc.CreateOrder(dto.Product, dto.Quantity)
		if err != nil {
			fmt.Printf("Error creating order: %v\n", err)
			http.Error(c.Writer, "Failed to create order", http.StatusBadRequest)
			return
		}

		fmt.Printf("Successfully created order: %s\n", order.ID)
		c.JSON(200, order)
	})

	fmt.Println("Listening at http://localhost:9999")
	err := r.Run(":9999")
	if err != nil {
		log.Panicf("Error creating HTTP service: %v\n", err)
	}

}
