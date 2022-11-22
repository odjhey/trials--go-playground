package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zeordermore/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	order, _ := models.CreateOrder(nil)
	fmt.Printf("Order %+v\n", order)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
	return
}
