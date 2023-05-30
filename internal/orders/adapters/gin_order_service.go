package adapters

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
)

// Adapter for the OrderService, exposes the service's endpoints as HTTP routes on a gin server
type GinOrderService struct {
	svc    orders.OrderService
	logger *zap.SugaredLogger
}

func NewGinAdapter(svc orders.OrderService, logger *zap.SugaredLogger) *GinOrderService {
	return &GinOrderService{
		svc:    svc,
		logger: logger,
	}
}

func (gos *GinOrderService) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/orders", func(c *gin.Context) {
		res, err := gos.svc.GetOrders()
		if err != nil {
			gos.logger.Errorf("Error getting orders: %v", err)
			http.Error(c.Writer, "Unknown error", http.StatusInternalServerError)
			return
		}

		c.JSON(200, res)
	})

	engine.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := gos.svc.GetOrder(id)
		if err != nil {
			gos.logger.Errorf("Error getting order %s: %v", id, err)
			http.Error(c.Writer, "Unknown error", http.StatusInternalServerError)
			return
		}

		c.JSON(200, res)
	})

	engine.POST("/orders", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			gos.logger.Errorf("Error reading body: %v", err)
			http.Error(c.Writer, "Failed to read body", http.StatusBadRequest)
			return
		}

		dto := &struct {
			Product  string
			Quantity int
		}{}
		err = json.Unmarshal(body, dto)
		if err != nil {
			gos.logger.Errorf("Error unmarshalling JSON: %v", err)
			http.Error(c.Writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		gos.logger.Infof("Creating order with product: %s, quantity: %d", dto.Product, dto.Quantity)
		order, err := gos.svc.CreateOrder(dto.Product, dto.Quantity)
		if err != nil {
			gos.logger.Infof("Error creating order: %v", err)
			http.Error(c.Writer, "Failed to create order", http.StatusBadRequest)
			return
		}

		gos.logger.Infof("Successfully created order: %s", order.ID)
		c.JSON(200, order)
	})
}
