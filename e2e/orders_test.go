package e2e

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestOrders_PostAndGet(t *testing.T) {
	product := gofakeit.Fruit()
	quantity := gofakeit.Number(1, 10)

	// GIVEN: I post an order
	post_res_body, _ := postJSON("http://localhost:9999/orders", map[string]interface{}{
		"Product":  product,
		"Quantity": quantity,
	})
	id := post_res_body["ID"]

	// WHEN: I get the order by id
	get_res_body, _ := getJSON(fmt.Sprintf("http://localhost:9999/orders/%s", id))

	// THEN: I get should get it back
	assert.Equal(t, get_res_body, map[string]interface{}{
		"Product":  product,
		"Quantity": float64(quantity), // Parsing JSON always returns float64
		"Status":   "RECEIVED",
		"ID":       id,
	})
}

func TestOrders_PostAndGetAll(t *testing.T) {
	product := gofakeit.Fruit()
	quantity := gofakeit.Number(1, 10)

	// GIVEN: I posted 2 orders
	post_res_body, _ := postJSON("http://localhost:9999/orders", map[string]interface{}{
		"Product":  product,
		"Quantity": quantity,
	})
	id_1 := post_res_body["ID"]

	post_res_body, _ = postJSON("http://localhost:9999/orders", map[string]interface{}{
		"Product":  product,
		"Quantity": quantity,
	})
	id_2 := post_res_body["ID"]

	// GIVEN: I call get all
	get_res_body, _ := getJSONArr("http://localhost:9999/orders")
	ids := []string{}
	for _, v := range get_res_body {
		ids = append(ids, v["ID"].(string))
	}

	// THEN: Both orders should be in the response
	assert.Contains(t, ids, id_1)
	assert.Contains(t, ids, id_2)
}

// func TestOrders_StressTest(t *testing.T) {
// 	for i := 0; i < 10000; i++ {
// 		product := gofakeit.Fruit()
// 		quantity := gofakeit.Number(1, 10)
// 		postJSON("http://localhost:9999/orders", map[string]interface{}{
// 			"Product":  product,
// 			"Quantity": quantity,
// 		})
// 	}
// }
