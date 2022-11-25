package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"zecrud/models"
)

func Food(w http.ResponseWriter, r *http.Request) {
	foodChan := make(chan models.Food, 1)
	reqError := make(chan error, 1)

	// for reals, this is just an exercise, use a decent lib for path parsing
	params := strings.Split(r.URL.Path, "/")

	switch r.Method {

	case http.MethodGet:
		go func() {
			food, _ := models.GetFood(params[2])
			foodChan <- food
		}()

	case http.MethodPost:
		go func() {
			// var somestring interface{}

			var jsonbody struct {
				Name    string "json:name"
				Cuisine string "json:cuisine"
			}
			err := json.NewDecoder(r.Body).Decode(&jsonbody)
			if err != nil {
				fmt.Println("body parse err", err)
			}

			food, _ := models.AddFood(r.Context(), models.AddFoodInput{Name: jsonbody.Name, Cuisine: jsonbody.Cuisine})
			foodChan <- food
		}()

	default:
		go func() {
			reqError <- errors.New("Route Not Found.")
		}()
	}

	select {
	case food := <-foodChan:
		fmt.Printf("Food %+v\n", food)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(food)
	case <-reqError:
		// TODO: handle more error types later
		w.WriteHeader(http.StatusNotFound)
	}

}
