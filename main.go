package main

import (
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

// IceCream(model) of type Struct
// FlavourId : ID of type int
// Name:       Flavour name of type string
// Price:      Price of type float64
// Serving:    Serving should be cone/cup of type string
type IceCream struct {
	FlavourId     int  `json:"flavourid"`
	Name   		  string  `json:"name"`
	Price  		  float64  `json:"price"`
	Serving 	  string `json:"serving"`
}
var flavours []IceCream

// Function to get all data items in the Icecream Struct 
// Args: 
//      http.ResponseWriter: http RestposeWrite for writing/displaying the data
//		http.Request :       http Request to respond to
func getAllflavours(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavours)
}

// Function to get a single IceCream flavour
// Args:
//      http.ResponseWriter: http RestposeWrite for writing/displaying the data
//		http.Request :       http Request to respond to  
func getIceCream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	flavourid_int,_ := strconv.Atoi(params["flavourid"])
	// Loop through flavours and find one with the flavourId from the params
	for _, item := range flavours {
		if item.FlavourId == flavourid_int {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&IceCream{})
}

// Function to add new IceCream flavour
// Args:
//      http.ResponseWriter: http RestposeWrite for writing/displaying the data
//		http.Request :       http Request to respond to
func createIceCream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var IceCream IceCream
	_ = json.NewDecoder(r.Body).Decode(&IceCream)
	IceCream.FlavourId = genNextId()
	flavours = append(flavours, IceCream)
	json.NewEncoder(w).Encode(IceCream)
}

// Function to return the flavourId of the last data item
func genNextId() int {
	lp := flavours[len(flavours) - 1]
	return lp.FlavourId + 1
}

// Function to update an IceCream flavour
// Args:
//      http.ResponseWriter: http RestposeWrite for writing/displaying the data
//		http.Request :       http Request to respond to
func updateIceCream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flavourid_int,_ := strconv.Atoi(params["flavourid"])
	for index, item := range flavours {
		if item.FlavourId == flavourid_int {
			flavours = append(flavours[:index], flavours[index+1:]...)
			var IceCream IceCream
			_ = json.NewDecoder(r.Body).Decode(&IceCream)
			IceCream.FlavourId = flavourid_int
			flavours = append(flavours, IceCream)
			json.NewEncoder(w).Encode(IceCream)
			return
		}
	}
}

// Function to delete an IceCream flavour
// Args:
//      http.ResponseWriter: http RestposeWrite for writing/displaying the data
//		http.Request :       http Request to respond to
func deleteIceCream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flavourid_int,_ := strconv.Atoi(params["flavourid"])
	for index, item := range flavours {
		if item.FlavourId == flavourid_int {
			flavours = append(flavours[:index], flavours[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(flavours)
}

// Main function
func main() {
	// Initiate a Gorillamux Router 
	r := mux.NewRouter()

	// Hardcoded data
	flavours = append(flavours, IceCream{FlavourId: 1, Name: "vanilla", Price: 65.50, Serving: "cup"})
	flavours = append(flavours, IceCream{FlavourId: 2, Name: "strawberry", Price: 50.00, Serving: "cone"})

	// Route handles & endpoints using various RESTful APIs
	r.HandleFunc("/flavours", getAllflavours).Methods("GET")
	r.HandleFunc("/flavours/{flavourid}", getIceCream).Methods("GET")
	r.HandleFunc("/flavours", createIceCream).Methods("POST")
	r.HandleFunc("/flavours/{flavourid}", updateIceCream).Methods("PUT")
	r.HandleFunc("/flavours/{flavourid}", deleteIceCream).Methods("DELETE")

	// Start server and listen to port localhost:8000
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Request sample
// {
// 	"flavourid": 4545454, //automatically assigns the flavour id according to the last item in the data base
// 	"name":"choclate",
// 	"price": 35.75,
//  "serving": "cup"
// }
