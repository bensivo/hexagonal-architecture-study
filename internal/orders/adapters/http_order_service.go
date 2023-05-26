package adapters

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			res, err := hos.svc.GetOrders()
			if err != nil {
				fmt.Printf("Error getting orders: %v\n", err)
				http.Error(w, "Unknown error", http.StatusInternalServerError)
				return
			}

			resJson, err := json.Marshal(res)
			if err != nil {
				fmt.Printf("Error marshalling orders: %v\n", err)
				http.Error(w, "Unknown error", http.StatusInternalServerError)
				return
			}

			_, err = io.WriteString(w, string(resJson))
			if err != nil {
				fmt.Printf("Error writing response: %v\n", err)
				return
			}
			return
		case http.MethodPost:
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Error reading body: %v\n", err)
				http.Error(w, "Failed to read body", http.StatusBadRequest)
				return
			}

			dto := &struct {
				Product  string
				Quantity int
			}{}
			err = json.Unmarshal(body, dto)
			if err != nil {
				fmt.Printf("Error unmarshalling JSON: %v\n", err)
				http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
				return
			}

			fmt.Printf("Creating order with product: %s, quantity: %d\n", dto.Product, dto.Quantity)
			order, err := hos.svc.CreateOrder(dto.Product, dto.Quantity)
			if err != nil {
				fmt.Printf("Error creating order: %v\n", err)
				http.Error(w, "Failed to create order", http.StatusBadRequest)
				return
			}

			fmt.Printf("Successfully created order: %s\n", order.ID)
			resJson, err := json.Marshal(order)
			if err != nil {
				fmt.Printf("Error marshalling order: %v\n", err)
				http.Error(w, "Unknown error", http.StatusInternalServerError)
				return
			}

			io.WriteString(w, string(resJson))
			return
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Listening at http://localhost:9999")
	err := http.ListenAndServe(":9999", mux)
	if err != nil {
		log.Panicf("Error creating HTTP service: %v\n", err)
	}
}
