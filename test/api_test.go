package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoodTrucks(t *testing.T) {
	url := "http://localhost:9999/foodtrucks"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	assert.Equal(t, err, nil, "Request is created succesfully is expected")

	res, err := client.Do(req)
	assert.Equal(t, err, nil, "Request is sent succesfully is expected")

	defer res.Body.Close()

	assert.Equal(t, 200, res.StatusCode, "OK response is expected")
}
