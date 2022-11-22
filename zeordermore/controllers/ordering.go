package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"zeordermore/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {

	orderChan := make(chan models.Order, 1)
	reqError := make(chan error, 1)

	switch r.Method {

	case http.MethodPost:
		go func() {
			ord, _ := models.CreateOrder(nil)
			orderChan <- ord
		}()

	default:
		go func() {
			reqError <- errors.New("Route Not Found.")
		}()
	}

	select {
	case ord := <-orderChan:
		fmt.Printf("Order %+v\n", ord)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ord)
	case <-reqError:
		// TODO: handle more error types later
		w.WriteHeader(http.StatusNotFound)
	}

}
