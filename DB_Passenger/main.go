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

//Json key-value mapping to current context
type PassengerInfo struct {
	FirstName string `json:"fn"`
	LastName  string `json:"ln"`
	MobileNo  string `json:"mn"`
	EmailAddr string `json:"ea"`
}

//Json key-value mapping to current context
type TripInfo struct {
	StartPostalCode string `json:"StartPostalCode"`
	EndPostalCode   string `json:"EndPostalCode"`
	TimeStamp       string `json:"timestamp"`
	DriverID        string `json:"driverID"`
	PassengerID     string `json:"passengerID"`
	TripStatus      string `json:"tripStatus"`
}

/*
func validKey(r *http.Request) bool {
    v := r.URL.Query()
    if key, ok := v["key"]; ok {
        if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}
*/
const baseURL = "http://localhost:"

//Getting all trips with passengerID from the passenger microservice
func getAllTrips(pID string, w http.ResponseWriter) {
	response, err := http.Get(baseURL + "5080/Passengers/" + pID + "/Trip") // -> trips, trips return using json.newencode
	var trips []TripInfo
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &trips)
		fmt.Println(trips)
		for _, trip := range trips {
			json.NewEncoder(w).Encode(trip)
		}
		/*
			fmt.Println(response.StatusCode)
			fmt.Println(string(data))
		*/
		response.Body.Close()

	}

}

//Get Trip assigned to driver by driverID -> driver microservice
func IsJsonValid(TempPassenger PassengerInfo) bool {
	valid := true

	if TempPassenger.EmailAddr == "" || TempPassenger.FirstName == "" || TempPassenger.LastName == "" || TempPassenger.MobileNo == "" {
		valid = false
	}
	return valid
}

//Handle passenger API's
func Passenger_Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "root:24A48g00@tcp(127.0.0.1:3306)/db_rideshare")

	//handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Header.Get("Content-type") == "application/json" {
		// POST is for creating new Driver
		if r.Method == "POST" {

			// read the string sent to the service
			var newPassenger PassengerInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newPassenger)
				jsonInfoValid := IsJsonValid(newPassenger)
				if !jsonInfoValid {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Passenger " +
							"information " + "in JSON format"))
					return
				}
				// check if Passenger exists; add only if
				// Passenger does not exist
				get_Result := CheckPExist(db, params["passengerid"])
				fmt.Println(get_Result)
				if !get_Result {
					InsertPRecord(db, newPassenger.FirstName,
						newPassenger.LastName, newPassenger.MobileNo, newPassenger.EmailAddr)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " + params["passengerid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Passenger Found"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply Passenger information " +
					"in JSON format"))
			}
		}
		//---PUT is for creating or updating
		// existing Passenger
		if r.Method == "PUT" {
			// read the string sent to the service
			var newPassenger PassengerInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newPassenger)
				jsonInfoValid := IsJsonValid(newPassenger)
				if !jsonInfoValid {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Passenger " +
							"information " + "in JSON format"))
					return
				}
				// check if Passenger exists; add only if
				// Passenger does not exist
				get_Result := CheckPExist(db, params["passengerid"])
				fmt.Printf("This is get result %t", get_Result)
				if !get_Result {
					InsertPRecord(db, newPassenger.FirstName,
						newPassenger.LastName, newPassenger.MobileNo, newPassenger.EmailAddr)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " +
						params["passengerid"]))
				} else {
					//update course
					EditPRecord(db, params["passengerid"], newPassenger.FirstName,
						newPassenger.LastName, newPassenger.MobileNo, newPassenger.EmailAddr)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Passenger Updated: " +
						params["passengerid"]))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply Passenger information " +
					"in JSON format"))
			}
		}
	}
	//GET getting the information by the ID from database!
	if r.Method == "GET" {
		get_Result := CheckPExist(db, params["passengerid"])
		if get_Result {
			json.NewEncoder(w).Encode(GetPassenger(db, params["passengerid"]))
			getAllTrips(params["passengerid"], w)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Passenger found"))
		}
	}
	//DELETE is not allowed!
	if r.Method == "DELETE" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 - Not Allowed to DELETE for Auditing"))
	}

	defer db.Close()
	fmt.Println("Database opened")
}

func Passenger_Handler_Trips(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "root:24A48g00@tcp(127.0.0.1:3306)/db_rideshare")

	//handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			status := GetPassenger(db, params["passengerid"]).OnRide
			if !status {
				setPassengerStatus(db, params["passengerid"], 1)
			} else {
				setPassengerStatus(db, params["passengerid"], 0)
			}
		}
	}
	defer db.Close()
	fmt.Println("Database opened")
}

func main() {
	router := mux.NewRouter()

	//Cors handling methods
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//Handling incoming API request
	router.HandleFunc("/Passengers/{passengerid}", Passenger_Handler).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/Passengers/{passengerid}/trips/OnRideStatus", Passenger_Handler_Trips).Methods("POST")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(router)))
}
