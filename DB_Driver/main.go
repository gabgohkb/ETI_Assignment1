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
type driverInfo struct {
	FirstName    string `json:"fn"`
	LastName     string `json:"ln"`
	MobileNo     string `json:"mn"`
	EmailAddr    string `json:"ea"`
	CarLicenseNo string `json:"cln"`
	IDNO         string `json:"idno"`
}

const baseURL = "http://localhost:"

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
//Make sure that response body is valid
func IsJsonValid(TempDriver driverInfo) bool {
	valid := true

	if TempDriver.EmailAddr == "" || TempDriver.FirstName == "" || TempDriver.LastName == "" || TempDriver.MobileNo == "" || TempDriver.IDNO == "" || TempDriver.CarLicenseNo == "" {
		valid = false
	}
	return valid
}

//Get Trip assigned to driver by driverID -> driver microservice
func getTripAssigned(dID string) string {
	response, err := http.Get(baseURL + "5080/Drivers" + "/" + dID + "/Trip")
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

//Handle driver API's
func driver_Handler(w http.ResponseWriter, r *http.Request) {
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
			var newDriver driverInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &newDriver)
				JsonInfoValid := IsJsonValid(newDriver)
				if !JsonInfoValid {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply driver " +
							"information " + "in JSON format"))
					return
				}
				// check if Driver exists; add only if
				// Driver does not exist
				get_Result := CheckDExist(db, params["driverid"], newDriver.IDNO)
				if !get_Result {
					InsertDRecord(db, newDriver.FirstName,
						newDriver.LastName, newDriver.MobileNo, newDriver.EmailAddr,
						newDriver.CarLicenseNo, newDriver.IDNO)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " +
						params["driverid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Driver Found"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply Driver information " +
					"in JSON format"))
			}
		}
		//---PUT is for creating or updating
		// existing Driver but Identification cannot be edited---
		if r.Method == "PUT" {
			// read the string sent to the service
			var newDriver driverInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newDriver)
				JsonInfoValid := IsJsonValid(newDriver)
				if !JsonInfoValid {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply driver " +
							"information " + "in JSON format"))
					return
				}
				// check if Driver exists; add only if
				// Driver does not exist
				get_Result := CheckDExist(db, params["driverid"], newDriver.IDNO)
				if !get_Result {
					InsertDRecord(db, newDriver.FirstName,
						newDriver.LastName, newDriver.MobileNo, newDriver.EmailAddr,
						newDriver.CarLicenseNo, newDriver.IDNO)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " +
						params["driverid"]))
				} else {
					//update course
					if CheckIdentificationNo(db, params["driverid"], newDriver.IDNO) {
						EditDRecord(db, params["driverid"], newDriver.FirstName,
							newDriver.LastName, newDriver.MobileNo, newDriver.EmailAddr,
							newDriver.CarLicenseNo)
						w.WriteHeader(http.StatusAccepted)
						w.Write([]byte("202 - Driver Updated: " +
							params["driverid"]))
					} else {
						w.WriteHeader(http.StatusConflict)
						w.Write([]byte("409 - Identification Number Not allowed to edit!"))
					}
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply Driver information " +
					"in JSON format"))
			}
		}
	}
	//DELETE is not allowed!
	if r.Method == "DELETE" {
		w.WriteHeader(
			http.StatusForbidden)
		w.Write([]byte("403 - Not Allowed to DELETE for Auditing"))
	}
	//GET getting the information by the ID from database!
	if r.Method == "GET" {
		get_Result := CheckDExist(db, params["driverid"], " ")
		if get_Result {
			json.NewEncoder(w).Encode(GetDriver(db, params["driverid"]))
			//show the trip
			currentTrip := getTripAssigned(params["driverid"])
			if currentTrip == "nil" {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - There are no trips assigned"))
			} else {
				json.NewEncoder(w).Encode(currentTrip)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Driver found"))
		}
	}
	defer db.Close()
	fmt.Println("Database opened")
}

//Driver communication with trips
func Driver_Handler_Trips(w http.ResponseWriter, r *http.Request) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "root:24A48g00@tcp(127.0.0.1:3306)/db_rideshare")

	//handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			driverID := r.URL.Query().Get("dID")
			fmt.Println(driverID)
			status := GetDriver(db, driverID).OnRide
			if !status {
				setDriverStatus(db, driverID, 1)
			} else {
				setDriverStatus(db, driverID, 0)
			}
		}
	}
	if r.Method == "GET" {
		dID := GetaAvailDriver(db)
		if dID == "nil" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Driver found"))
		} else {
			json.NewEncoder(w).Encode(dID)
		}

	}
	defer db.Close()
	fmt.Println("Database opened")
}
func main() {
	//create instance for router
	router := mux.NewRouter()
	//Cors handling methods
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//Handling incoming API request
	router.HandleFunc("/Drivers/{driverid}", driver_Handler).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/Drivers/Trips/trip", Driver_Handler_Trips).Methods("GET", "PUT", "POST", "DELETE")
	fmt.Println("Listening at port 5050")
	log.Fatal(http.ListenAndServe(":5050", handlers.CORS(headers, methods, origins)(router)))
}
