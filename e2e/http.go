package e2e

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Make a POST request to an endpoint, and return the response as a map[string]interface{}
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

// Make a GET request to an endpoint, and return the response as a map[string]interface{}
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

// Make a GET request to an endpoint, and return the response as an array of maps
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
