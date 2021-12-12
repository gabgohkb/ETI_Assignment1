<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#usage">Usage</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>


<!-- ABOUT THE PROJECT -->
## About The Project

Passenger page for getting current status, viewing trips and requesting for trips.
[![Passenger_Screen_Shot][passenger-screenshot]]

Driver page for getting current status, starting and ending ride.
[![Driver_Screen_Shot][driver-screenshot]]

The Ride Share Assignment requires at least 2 microservices that is derived from applying the techniques of Decomposition using Strategic DDD and Tactical DDD. It also has an optional requirement for a frontend to be implemented. The following will explain what each of the microservices currently do, how is the services derieved and what frontend is being used together with a visual representation of the Microservices diagram. 

There are a total of 3 microservices:
(the functions and http method will be shown in each microservice.)
1. Passenger
  * GET - Log in, Get Passenger Details 
    * Example: "http://localhost:5000/Passengers/1"
  * PUT - Update Passenger Details 
    * Example: Update Passenger Act: curl -H "Content-Type: application/json" -X PUT http://localhost:5000/Passengers/{passengerid} -d "{\"fn\":\"jacob\",\"ln\":\"lim\",\"mn\":\"92323812\",\"ea\":\"jl@gmail.com\"}"
  * POST - Create Passenger
    * Example: Create Passenger Acct: curl -H "Content-Type: application/json" -X POST http://localhost:5000/Passengers/{passengerid} -d "{\"fn\":\"jacob\",\"ln\":\"lim\",\"mn\":\"92323812\",\"ea\":\"jl@gmail.com\"}"
  * DELETE - Delete will not be allowed (Error message will occur)
    * Example: curl -X DELETE http://localhost:5000/Passengers/1
  * GET - View all trips (in reverse chronological order) 
    * Example: http://localhost:5080/Passengers/1/Trip
2. Trip
  * PUT (But Post was used to accomodate to golang) - Updates OnRide status (Driver & Passenger)
    * "http://localhost:5050/Drivers/Trips/trip?dID={driverID}"
    * "http://localhost:5000/Passengers/{passengerID}/trips/OnRideStatus"
  * POST - Request trip (Passenger)
    * Example: curl -H "Content-Type: application/json" -X POST http://localhost:5080/Passengers/1/Trip -d "{\"startpc\":\"670210\",\"endpc\":\"210999\",\"timestamp\":\"2021-11-23 17:59:59\",\"dID\":\"\",\"pID\":\"\",\"tripStatus\":\"\"}"
  * GET - Retrieve Available driver
    * "http://localhost:5050/Drivers/Trips/trip"
  * PUT - Start ride (Driver)
    * curl -H "Content-Type: application/json" -X PUT http://localhost:5080/Drivers/1/Trip?action=startTrip
  * PUT - End ride (Driver)
    * curl -H "Content-Type: application/json" -X PUT http://localhost:5080/Drivers/1/Trip?action=endTrip
3. Driver
  * GET - Log in, Get Driver Details 
    * Example: curl http://localhost:5050/Drivers/1
  * PUT - Update Driver Details ('identificationNo' will not be allowed)
    * Example: curl -H "Content-Type: application/json" -X PUT http://localhost:5050/Drivers/{driverid} -d "{\"fn\":\"jacob\",\"ln\":\"lim\",\"mn\":\"92323812\",\"ea\":\"jl@gmail.com\",\"cln\":\"s111fx\",\"idno\":\"s1010201F\"}"
  * POST - Create Driver
    * Example: curl -H "Content-Type: application/json" -X POST http://localhost:5050/Drivers/{driverid} -d "{\"fn\":\"jacob\",\"ln\":\"lim\",\"mn\":\"92323812\",\"ea\":\"jl@gmail.com\",\"cln\":\"sx234f\",\”idno\”:\”t0029201F\”}"
  * DELETE - Delete will not be allowed (Error message will occur)
    * Example: curl -X DELETE http://localhost:5050/Drivers/1
  * GET - View Current Trip 
    * Example: curl http://localhost:5050/Drivers/1

### Design Consideration of Microservices

Why 3 Microservice: 
  * With Driver, Passenger and Trips seperated it allows lesser dependencies between how they will communicate. 
  * Driver and Passenger will focus on functions like Creating, Updating and Getting their personnal details. 
  * Trips will focus on Requesting of trip, Assigning trip and both start and end trips. 
  * This allows each area of focus to be seperated in case if one microservice is required to be shut down for enhancement or errors the other microservices can still operate thus ensuring user experience from the running microservices. 
  * Example: If Trips microservice is shutdown, Driver and Passenger microservice will still allow the users to create  profiles, update and get personnal details. 

Frontend: 
  * Monolith Frontend:
    * Above is chosen because this assignment's implementations will not be growing and it also shows how technologically agnostic the frontend, microservice and backend can be. 

Design Pattern used: 
  * Decomposition: 
    * Strategic Domain-Driven Design Result: 
      * Accounts 
      * Trips Management 
    * Tactical Domain-Driven Design Identified Microservice: 
      * Driver Microservice (Handle Functions: Create, Retrieve, Update Driver Account, Retrieve Current Trip)
      * Passenger Microservice (Handle Functions: Create, Retrieve, Update Passenger Account, Retrieve List of Trips Taken)
      * Trip Microservice (Handle Functions: Request Trip, Start Trip, End Trip, Update Account On Ride Status)

<p align="left">(<a href="#top">back to top</a>)</p>

To better understand, below is a diagram of the assignment's structure and how communications are made.
[![ArchitectureDiagram-screenshot][architecturediagram-screenshot]]
[![ArchitectureDiagramEach-screenshot][architecturediagrameach-screenshot]]


<p align="left">(<a href="#top">back to top</a>)</p>


### Built With

The main objective of this assignment is to put the knowledge and skills learnt about Golang to use. 
Vanilla Javascript and HTML CSS is used for the frontend.

* [Golang](https://go.dev/)
* [HTML](https://html.com/)
* [JavaScript](https://www.javascript.com/)
* [JQuery](https://jquery.com)
<p align="left">(<a href="#top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

Make sure that MySQL and Golang is downloaded on your device.

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/gabgohkb/ETI_Assignment1.git
   ```
2. Install necessary libraries
   ```sh
   go get -u github.com/go-sql-driver/mysql
   go get -u github.com/gorilla/mux
   go get -u github.com/gorilla/handlers
   ```
3. Execute database start script in `/DB/ExecuteDBSQL.sql`

<p align="left">(<a href="#top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

To start using the ride-hailing platform, follow the below steps:
1. Run all microservices in each directory
 ```sh
 cd Assignment1\DB_Passenger
 go run main.go P_Database_Func.go
 ```
 ```sh
 cd Assignment1\DB_Driver
  go run main.go D_Database_Func.go
 ```
 ```sh
 cd Assignment1\DB_Trip
go run main.go T_Database_Func.go
 ```
2. Open frontend by opening `Home.html` in `Assignment1\Templates`

<p align="left">(<a href="#top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

- [x] Backend using Golang
- [x] Frontend using HTML JavaScript
- [x] Tidy up both backend and frontend

<p align="left">(<a href="#top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="left">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Gabriel Goh - [School Email](mailto:s10198258@connect.np.edu.sg) 

Project Link: [https://github.com/gabgohkb/ETI_Assignment1.git](https://github.com/gabgohkb/ETI_Assignment1.git)

<p align="left">(<a href="#top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Driver-screenshot]: images/driver.PNG
[Passenger-screenshot]: images/passenger.PNG
[TripPassenger-screenshot]: images/Ptrip.PNG
[TripDriver-screenshot]: images/Dtrip.PNG
[ArchitectureDiagram-screenshot]: images/ArchitectureDiagram.PNG
[ArchitectureDiagramEach-screenshot]: images/ArchitectureDiagramEach.PNG


