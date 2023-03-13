package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"strconv"

	// package yang digunakan untuk mengubah string menjadi tipe int
	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"go-postgres-crud/models" //models package dimana Buku didefinisikan

	// digunakan untuk mendapatkan parameter dari router
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []models.Employee `json:"data"`
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	//create an empty employee
	var employee models.Employee

	//decode data json req to employee
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		log.Fatalf("cannot decode request body. %v", err)
	}

	//call models then insert
	insertID := models.addEmployee(employee)

	//format response object
	res := response{
		ID:      insertID,
		Message: "Employee data inserted",
	}

	json.NewEncoder(w).Encode(res)
}

func GetEmployee(w http.ResponseWriter, r *http.Request)  {
	//set header
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//get id employee from param request with key param "id"
	params := mux.Vars(r)

	//convert string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("cannot change string to int. %v", err)
	}

	//call models getoneemployee with param "id" for one row
	employee, err:models.GetOneEmployee(int64(id))

	if err != nil {
		log.Fatalf("cannot get employee data. %v", err)
	}

	json.NewEncoder(w).Encode(employee)
}