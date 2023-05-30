package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func postJSON(url string, body map[string]interface{}) (map[string]interface{}, error) {
	post_req_body, _ := json.Marshal(body)
	post_req_buffer := bytes.NewBuffer(post_req_body)
	post_res, err := http.Post(url, "application/json", post_req_buffer)
	if err != nil {
		return nil, err
	}
	defer post_res.Body.Close()

	post_res_bytes, err := ioutil.ReadAll(post_res.Body)
	if err != nil {
		return nil, err
	}
	post_res_body := make(map[string]interface{})
	json.Unmarshal(post_res_bytes, &post_res_body)

	return post_res_body, nil
}

func getJSON(url string) (map[string]interface{}, error) {
	get_res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer get_res.Body.Close()

	get_res_bytes, err := ioutil.ReadAll(get_res.Body)
	if err != nil {
		return nil, err
	}
	get_res_body := make(map[string]interface{})
	json.Unmarshal(get_res_bytes, &get_res_body)

	return get_res_body, nil
}

func getJSONArr(url string) ([]map[string]interface{}, error) {
	get_res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer get_res.Body.Close()

	get_res_bytes, err := ioutil.ReadAll(get_res.Body)
	if err != nil {
		return nil, err
	}
	get_res_body := make([]map[string]interface{}, 0)
	json.Unmarshal(get_res_bytes, &get_res_body)

	return get_res_body, nil
}
