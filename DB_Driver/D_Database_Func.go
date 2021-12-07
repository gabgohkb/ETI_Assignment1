package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Driver struct { // map this type to the record in the table
	ID               string
	FirstName        string
	LastName         string
	MobileNo         string
	EmailAddr        string
	CarLicenseNo     string
	IdentificationNo string
	OnRide           bool
}

/*
//Delete Driver Details in Database
func DeleteDRecord(db *sql.DB, ID string) {
	query := fmt.Sprintf(
		"DELETE FROM driver WHERE driverID='%s'", ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
*/

//Check if there is a change in the driver Identification Number -- return false if identification number is wrong.
func CheckIdentificationNo(db *sql.DB, dID string, identifiNo string) bool {
	query := fmt.Sprintf("SELECT * FROM driver WHERE driverID = '%s'", dID)

	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var driver Driver
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&driver.ID, &driver.FirstName,
			&driver.LastName, &driver.MobileNo, &driver.EmailAddr, &driver.CarLicenseNo, &driver.IdentificationNo, &driver.OnRide)
		if err != nil {
			panic(err.Error())
		}
	}
	return bool(driver.IdentificationNo == identifiNo)
}

//Edit Driver Details in Database
func EditDRecord(db *sql.DB, ID string, FN string, LN string, MN string, EA string, CLN string) {
	query := fmt.Sprintf(
		"UPDATE driver SET FirstName='%s', LastName='%s',MobileNo='%s',EmailAddr='%s',CarLicenseNo='%s'  WHERE driverID=%s",
		FN, LN, MN, EA, CLN, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Set driver status to true or false, determined by the current status
func setDriverStatus(db *sql.DB, dID string, or int) {
	query := fmt.Sprintf("UPDATE driver SET OnRide = '%d' WHERE driverID = '%s'", or, dID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Insert a new Driver in Database
func InsertDRecord(db *sql.DB, FN string, LN string, MN string, EA string, CLN string, IdentificationNo string) {
	query := fmt.Sprintf("INSERT INTO driver(FirstName,LastName,MobileNo,EmailAddr,CarLicenseNo,identificationNo) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')",
		FN, LN, MN, EA, CLN, IdentificationNo)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

//Get Avail Driver from Database with OnRide as false
func GetaAvailDriver(db *sql.DB) string {
	results, err := db.Query("SELECT * FROM driver WHERE OnRide = false ORDER BY RAND() LIMIT 1")
	tempDID := "nil"
	if err != nil {
		panic(err.Error())
	}
	var driver Driver
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&driver.ID, &driver.FirstName,
			&driver.LastName, &driver.MobileNo, &driver.EmailAddr, &driver.CarLicenseNo, &driver.IdentificationNo, &driver.OnRide)
		if err != nil {
			panic(err.Error())
		}
		tempDID = driver.ID
	}
	return tempDID
}

//Get driver from database
func GetDriver(db *sql.DB, ID string) Driver {
	query := fmt.Sprintf("SELECT * FROM driver WHERE driverID = '%s'", ID)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var driver Driver
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&driver.ID, &driver.FirstName,
			&driver.LastName, &driver.MobileNo, &driver.EmailAddr, &driver.CarLicenseNo, &driver.IdentificationNo, &driver.OnRide)
		if err != nil {
			panic(err.Error())
		}
	}

	return driver
}

//check if a Driver exist in Database
func CheckDExist(db *sql.DB, ID string, identifiNo string) bool {
	query := fmt.Sprintf("SELECT * FROM driver WHERE driverID = '%s'OR identificationNo = '%s'", ID, identifiNo)

	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var driver Driver
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&driver.ID, &driver.FirstName,
			&driver.LastName, &driver.MobileNo, &driver.EmailAddr, &driver.CarLicenseNo, &driver.IdentificationNo, &driver.OnRide)
		if err != nil {
			panic(err.Error())
		}
	}
	return bool(driver.ID == ID || driver.IdentificationNo == identifiNo)
}

//Get All Driver from Database
func GetDRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM db_rideshare.driver")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var driver Driver
		err = results.Scan(&driver.ID, &driver.FirstName,
			&driver.LastName, &driver.MobileNo, &driver.EmailAddr, &driver.CarLicenseNo, &driver.IdentificationNo, &driver.OnRide)
		if err != nil {
			panic(err.Error())
		}
	}
}
