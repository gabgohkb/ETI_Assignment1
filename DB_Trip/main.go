package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const baseURL = "http://localhost:"

//Trip data structure
type TripInfo struct {
	StartPostalCode string `json:"startpc"`
	EndPostalCode   string `json:"endpc"`
	TimeStamp       string `json:"timestamp"`
	DriverID        string `json:"dID"`
	PassengerID     string `json:"pID"`
	TripStatus      string `json:"tripStatus"`
}

//Check if Json from response body is valid
func IsJsonValid(tempTrip TripInfo) bool {
	valid := true

	if tempTrip.StartPostalCode == "" || tempTrip.EndPostalCode == "" || tempTrip.TimeStamp == "" {
		valid = false
	}
	return valid
}

//After the passenger request for a trip, this func will get passenger OnRide status from Passenger microservice
func getAdriver() string {
	response, err := http.Get(baseURL + "5050/Drivers/Trips/trip")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		/*
			fmt.Println(response.StatusCode)
			fmt.Println(string(data))
		*/
		response.Body.Close()
		return string(data)
	}
	return "nil"
}

//Set passanger OnRide status
func setDriverOnRideStatus(dID string) {
	fmt.Println(dID)
	response, err := http.Post(baseURL+"5050/Drivers/Trips/trip?dID="+dID, "application/json", nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//Set passanger OnRide status
func setPassengerOnRideStatus(pID string) {
	fmt.Println(pID)
	response, err := http.Post(baseURL+"5000/Passengers/"+pID+"/trips/OnRideStatus", "application/json", nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//Function to handle API for Passenger trips
func trip_Handler_Request(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "root:24A48g00@tcp(127.0.0.1:3306)/db_rideshare")

	//handle error
	if err != nil {
		panic(err.Error())
	}

	// POST is for creating new trip
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			// read the string sent to the service
			var newTrip TripInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &newTrip)
				fmt.Println(newTrip)
				jsonInfoValid := IsJsonValid(newTrip)
				if !jsonInfoValid {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Trip " +
							"information " + "in JSON format"))
					return
				}
				IsPOnRide := CheckPtripOnGoing(db, params["passengerid"])
				fmt.Println(IsPOnRide)
				if IsPOnRide {
					//find a driver
					driverFound := getAdriver()
					if driverFound != "nil" {
						//create a trip
						InsertTrip(db, newTrip.StartPostalCode, newTrip.EndPostalCode, newTrip.TimeStamp, driverFound, params["passengerid"])
						w.WriteHeader(http.StatusCreated)
						w.Write([]byte("201 - Trip Requested by: " +
							params["passengerid"] + " Your Driver's ID is: " + driverFound))
					} else {
						w.WriteHeader(http.StatusNotFound)
						w.Write([]byte("404 - No Driver found"))
					}
				} else {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte("202 - Passenger is currently on a ride Unable to create another trip"))
				}
			}

		}
	}
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(GetPassengerTrips(db, params["passengerid"]))
	}
}

//Function to handle API for driver trips
func trip_Handler_driver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "root:24A48g00@tcp(127.0.0.1:3306)/db_rideshare")

	//handle error
	if err != nil {
		panic(err.Error())
	}
	// PUT is for starting a trip
	// and ending of the trip
	if r.Header.Get("Content-type") == "application/json" {
		// returns the key/value pairs in the query string as a map object
		actionFromDriver := r.URL.Query().Get("action")
		if r.Method == "PUT" {
			if actionFromDriver == "startTrip" {
				if GetDriverCurrentTrip_Waiting(db, params["driverid"]).TripStatus != "" {
					pID := GetDriverCurrentTrip_Waiting(db, params["driverid"]).PassengerID
					tripID := GetDriverCurrentTrip_Waiting(db, params["driverid"]).ID
					//set driver OnRide == true
					setDriverOnRideStatus(params["driverid"])
					//set passenger OnRide == true
					setPassengerOnRideStatus(pID)
					//update trip status
					UpdateTripStatus(db, "onTrip", tripID)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Trip Started by: " +
						params["driverid"]))
				} else {
					w.WriteHeader(http.StatusNotAcceptable)
					w.Write([]byte("406 - There are no trips to start"))
				}
			}
			fmt.Println(actionFromDriver)
			if actionFromDriver == "endTrip" {
				if GetDriverCurrentTrip_Onride(db, params["driverid"]).TripStatus != "" {
					pID := GetDriverCurrentTrip_Onride(db, params["driverid"]).PassengerID
					tripID := GetDriverCurrentTrip_Onride(db, params["driverid"]).ID
					//set driver OnRide == false
					setDriverOnRideStatus(params["driverid"])
					//set passenger OnRide == false
					setPassengerOnRideStatus(pID)
					//update trip status
					UpdateTripStatus(db, "Completed", tripID)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Trip end by: " +
						params["driverid"]))
				} else {
					w.WriteHeader(http.StatusNotAcceptable)
					w.Write([]byte("406 - There are no trips to end"))
				}
			}
		}
	}
	if r.Method == "GET" {
		//Getting driver's current trip assigned by the system
		trip := GetDriverCurrentTrip(db, params["driverid"])
		if trip.DriverID == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Trip found"))
		} else {
			//Send to the driver after get request for trip
			json.NewEncoder(w).Encode(trip)
		}
	}
}

func main() {
	router := mux.NewRouter()
	//Cors handling methods
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//routing api calls to functions
	router.HandleFunc("/Passengers/{passengerid}/Trip", trip_Handler_Request).Methods("POST", "GET")
	router.HandleFunc("/Drivers/{driverid}/Trip", trip_Handler_driver).Methods("GET", "PUT", "POST", "DELETE")
	fmt.Println("Listening at port 5080")
	log.Fatal(http.ListenAndServe(":5080", handlers.CORS(headers, methods, origins)(router)))
}
