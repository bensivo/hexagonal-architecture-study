package adapters

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
)

// Exposes an OrderService using HTTP endpoints
type HttpOrderService struct {
	svc    orders.OrderService
	logger *zap.SugaredLogger
}

func NewHttpAdapter(svc orders.OrderService, logger *zap.SugaredLogger) *HttpOrderService {
	return &HttpOrderService{
		svc:    svc,
		logger: logger,
	}
}

func (hos *HttpOrderService) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/orders", func(c *gin.Context) {
		res, err := hos.svc.GetOrders()
		if err != nil {
			hos.logger.Errorf("Error getting orders: %v", err)
			http.Error(c.Writer, "Unknown error", http.StatusInternalServerError)
			return
		}

		c.JSON(200, res)
	})

	engine.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := hos.svc.GetOrder(id)
		if err != nil {
			hos.logger.Errorf("Error getting order %s: %v", id, err)
			http.Error(c.Writer, "Unknown error", http.StatusInternalServerError)
			return
		}

		c.JSON(200, res)
	})

	engine.POST("/orders", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			hos.logger.Errorf("Error reading body: %v", err)
			http.Error(c.Writer, "Failed to read body", http.StatusBadRequest)
			return
		}

		dto := &struct {
			Product  string
			Quantity int
		}{}
		err = json.Unmarshal(body, dto)
		if err != nil {
			hos.logger.Errorf("Error unmarshalling JSON: %v", err)
			http.Error(c.Writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		hos.logger.Infof("Creating order with product: %s, quantity: %d", dto.Product, dto.Quantity)
		order, err := hos.svc.CreateOrder(dto.Product, dto.Quantity)
		if err != nil {
			hos.logger.Infof("Error creating order: %v", err)
			http.Error(c.Writer, "Failed to create order", http.StatusBadRequest)
			return
		}

		hos.logger.Infof("Successfully created order: %s", order.ID)
		c.JSON(200, order)
	})
}
