-- -- Creating Database for Assignment 1 ETI
CREATE database db_rideshare;

USE db_rideshare;


-- -- Create TABLE for Driver entity
CREATE TABLE driver (

    driverID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    FirstName VARCHAR(30), 
    LastName VARCHAR(30), 
    MobileNo VARCHAR(15),
    EmailAddr VARCHAR(50),
    CarLicenseNo VARCHAR(7),
    identificationNo VARCHAR(9),
    OnRide tinyint(1) DEFAULT 0

); 

-- -- Create TABLE for Passenger entity
CREATE TABLE Passenger (

    passengerID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    FirstName VARCHAR(30), 
    LastName VARCHAR(30), 
    MobileNo VARCHAR(15),
    EmailAddr VARCHAR(50),
    OnRide tinyint(1) DEFAULT 0

); 

-- -- Create TABLE for Trips entity
CREATE TABLE trips (

    tripID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    startPostalCode VARCHAR(6), 
    endPostalCode VARCHAR(6), 
    timeStart datetime,
    passengerID INT,
    tripStatus VARCHAR(15), 
    driverID VARCHAR(5)

); 


