package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Passenger struct { // map this type to the record in the table
	ID        string
	FirstName string
	LastName  string
	MobileNo  string
	EmailAddr string
	OnRide    bool
}

//Delete Passenger Details in Database
func DeletePRecord(db *sql.DB, ID string) {
	query := fmt.Sprintf(
		"DELETE FROM passenger WHERE passengerID='%s'", ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Edit Passenger Details in Database
func EditPRecord(db *sql.DB, ID string, FN string, LN string, MN string, EA string) {
	query := fmt.Sprintf(
		"UPDATE passenger SET FirstName='%s', LastName='%s',MobileNo='%s',EmailAddr='%s'  WHERE passengerID=%s",
		FN, LN, MN, EA, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Insert a new Passenger in Database
func InsertPRecord(db *sql.DB, FN string, LN string, MN string, EA string) {
	query := fmt.Sprintf("INSERT INTO passenger(FirstName,LastName,MobileNo,EmailAddr) VALUES ('%s', '%s', '%s', '%s')",
		FN, LN, MN, EA)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

//check if a passenger exist in Database
func CheckPExist(db *sql.DB, ID string) bool {
	query := fmt.Sprintf("SELECT * FROM passenger WHERE passengerID = '%s'", ID)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var passenger Passenger
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&passenger.ID, &passenger.FirstName,
			&passenger.LastName, &passenger.MobileNo, &passenger.EmailAddr, &passenger.OnRide)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(passenger.ID, passenger.FirstName,
			passenger.LastName, passenger.MobileNo, passenger.EmailAddr)
	}

	return bool(passenger.ID == ID)
}

//Set passenger status to true or false, determined by the current status
func setPassengerStatus(db *sql.DB, pID string, or int) {
	query := fmt.Sprintf("UPDATE passenger SET OnRide = '%d' WHERE passengerID = '%s'", or, pID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Get a Passenger from database
func GetPassenger(db *sql.DB, ID string) Passenger {
	query := fmt.Sprintf("SELECT * FROM passenger WHERE passengerID = '%s'", ID)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var passenger Passenger
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&passenger.ID, &passenger.FirstName,
			&passenger.LastName, &passenger.MobileNo, &passenger.EmailAddr, &passenger.OnRide)
		if err != nil {
			panic(err.Error())
		}
	}

	return passenger
}

//Get All Passenger from Database
func GetPRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM db_rideshare.passenger")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var passenger Passenger
		err = results.Scan(&passenger.ID, &passenger.FirstName,
			&passenger.LastName, &passenger.MobileNo, &passenger.EmailAddr, &passenger.OnRide)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(passenger.ID, passenger.FirstName,
			passenger.LastName, passenger.MobileNo, passenger.EmailAddr)
	}
}
