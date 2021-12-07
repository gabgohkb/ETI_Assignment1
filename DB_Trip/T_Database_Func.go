package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Trip struct { // map this type to the record in the table
	ID              string //trip ID
	StartPostalCode string // Varchar(6) start postal code
	EndPostalCode   string // Varchar(6) end  postal code
	TimeStamp       string // datetime of trip creation
	DriverID        string // driver id from driver db
	PassengerID     string // passenger id from passenger db
	TripStatus      string //"Waiting", "onTrip" , "Completed" variables
}

//Insert the trip into the database: Initiated by the passenger from trip_Handler_Request
func InsertTrip(db *sql.DB, spc string, epc string, ts string, dID string, pID string) {
	query := fmt.Sprintf("INSERT INTO trips(startPostalCode,endPostalCode,timeStart,driverID,passengerID) VALUES('%s','%s','%s',%s,'%s')", spc, epc, ts, dID, pID)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

//Getting current trip for driver from Database, one driver will always only have 1 tripStatus as 'waiting'
func GetDriverCurrentTrip(db *sql.DB, dID string) Trip {
	query := fmt.Sprintf("SELECT * FROM trips WHERE driverID = '%s' AND (tripStatus = 'Waiting' OR tripStatus = 'OnTrip')", dID)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var trip Trip
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&trip.ID, &trip.StartPostalCode, &trip.EndPostalCode, &trip.TimeStamp, &trip.PassengerID, &trip.TripStatus, &trip.DriverID)
		if err != nil {
			panic(err.Error())
		}
	}

	return trip
}

//check if a Driver exist in Database
func CheckPtripOnGoing(db *sql.DB, pID string) bool {
	query := fmt.Sprintf("SELECT * FROM trips WHERE passengerID = '%s' AND (tripStatus = 'Waiting' OR tripStatus = 'OnTrip')", pID)

	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	NotIn := true
	var trip Trip
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&trip.ID, &trip.StartPostalCode, &trip.EndPostalCode, &trip.TimeStamp, &trip.PassengerID, &trip.TripStatus, &trip.DriverID)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(trip.ID, trip.StartPostalCode, trip.EndPostalCode, trip.TimeStamp, trip.DriverID, trip.PassengerID, trip.TripStatus)
		if trip.PassengerID == pID {

			NotIn = false
		}
	}
	return NotIn
}

//Getting current trip for driver from Database, one driver will always only have 1 tripStatus as 'OnRide'
func GetDriverCurrentTrip_Onride(db *sql.DB, dID string) Trip {
	query := fmt.Sprintf("SELECT * FROM trips WHERE driverID = '%s' AND tripStatus = 'OnTrip'", dID)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var trip Trip
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&trip.ID, &trip.StartPostalCode, &trip.EndPostalCode, &trip.TimeStamp, &trip.PassengerID, &trip.TripStatus, &trip.DriverID)
		if err != nil {
			panic(err.Error())
		}
	}

	return trip
}

//Getting current trip for driver from Database, one driver will always only have 1 tripStatus as 'Waiting'
func GetDriverCurrentTrip_Waiting(db *sql.DB, dID string) Trip {
	query := fmt.Sprintf("SELECT * FROM trips WHERE driverID = '%s' AND tripStatus = 'Waiting'", dID)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var trip Trip
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&trip.ID, &trip.StartPostalCode, &trip.EndPostalCode, &trip.TimeStamp, &trip.PassengerID, &trip.TripStatus, &trip.DriverID)
		if err != nil {
			panic(err.Error())
		}
	}

	return trip
}

//Updating the trip status
func UpdateTripStatus(db *sql.DB, status string, tripid string) {
	query := fmt.Sprintf(
		"UPDATE trips SET tripStatus = '%s' WHERE tripID = '%s'", status, tripid)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Get all of the trips records for a particular passenger, return type: Array of Trip
func GetPassengerTrips(db *sql.DB, pID string) []Trip {
	query := fmt.Sprintf("SELECT * FROM trips WHERE passengerID = '%s' ORDER BY timeStart DESC", pID)

	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var trips []Trip
	for results.Next() {
		var trip Trip
		err = results.Scan(&trip.ID, &trip.StartPostalCode, &trip.EndPostalCode, &trip.TimeStamp, &trip.PassengerID, &trip.TripStatus, &trip.DriverID)
		if err != nil {
			panic(err)
		}
		trips = append(trips, trip)
	}
	fmt.Println(trips)
	return trips
}
