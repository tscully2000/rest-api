package main

import (
		"encoding/json"
		"log"
		"net/http"

		"github.com/gorilla/mux"
)

// Car constructor
type Car struct {
		ID    string `json:"id,omitempty"`
		Year  int    `json:"year"`
		Make  string `json:"make"`
		Model string `json:"model"`
}

var cars []Car

func main() {
		router, port, path, pathID := mux.NewRouter(), ":8000", "/api/cars", "/api/cars/{id}"
		cars = append(cars,
				Car{ID: "1", Year: 2017, Make: "Subaru", Model: "Impreza"},
				Car{ID: "2", Year: 2001, Make: "Volvo", Model: "S60"},
				Car{ID: "3", Year: 1998, Make: "Ford", Model: "Escort"},
		)
		router.HandleFunc(path, getCars).Methods("GET")
		router.HandleFunc(pathID, getCar).Methods("GET")
		router.HandleFunc(pathID, createCar).Methods("POST")
		router.HandleFunc(pathID, updateCar).Methods("PUT")
		router.HandleFunc(pathID, deleteCar).Methods("DELETE")
		log.Fatal(http.ListenAndServeTLS(port, "cert.pem", "key.pem", router))
}

func getCars(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cars)
}

func getCar(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, vehicle := range cars {
				if vehicle.ID == params["id"] {
						json.NewEncoder(w).Encode(vehicle)
						return
				}
		}
		json.NewEncoder(w).Encode(&Car{})
}

func createCar(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var car Car
		_ = json.NewDecoder(r.Body).Decode(&car)
		car.ID = params["id"]
		cars = append(cars, car)
		json.NewEncoder(w).Encode(car)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for i, vehicle := range cars {
				if vehicle.ID == params["id"] {
						cars = append(cars[:i], cars[i+1:]...)
						var car Car
						_ = json.NewDecoder(r.Body).Decode(&car)
						car.ID = params["id"]
						cars = append(cars, car)
						json.NewEncoder(w).Encode(car)
						return
				}
		}
		json.NewEncoder(w).Encode(cars)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for i, vehicle := range cars {
				if vehicle.ID == params["id"] {
						cars = append(cars[:i], cars[i+1:]...)
						break
				}
		}
		json.NewEncoder(w).Encode(cars)
}
