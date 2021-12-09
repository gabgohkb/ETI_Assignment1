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

This assignment's requirements are to create 2 or more microservices which will be ran in the background while a frontend such as website can be used to interact with the functions in the microservices. 

There are a total of 3 microservices:
(the functions will be shown in each microservice.)
1. Passenger
  * Log in
  * Update Passenger Details 
  * Create Passenger
  * Get Passenger Details 
  * Delete will not be allowed (Error message will occur)
  * View all trips (in reverse chronological order) 
2. Trip
  * Updates OnRide status (Driver & Passenger)
  * Request trip (Passenger)
  * Retrieves trip details (Driver)
  * Start ride (Driver)
  * End ride (Driver)
3. Driver
  * Log in
  * Update Driver Details ('identificationNo' will not be allowed)
  * Create Driver
  * Get Driver Details 
  * Delete will not be allowed (Error message will occur)
  * View Current Trip 

### Design Consideration of Microservices

Why 3 Microservice: 
  * With Driver, Passenger and Trips seperated it allows lesser dependencies between how they will communicate. 
  * Driver and Passenger will focus on functions like Creating, Updating and Getting their personnal details. 
  * Trips will focus on Requesting of trip, Assigning trip and both start and end trips. 
  * This allows each area of focus to be seperated in case if one microservice is required to be shut down for enhancement or errors the other microservices can still operate thus ensuring user experience from the running microservices. 
  * Example: If Trips microservice is shutdown, Driver and Passenger microservice will still allow the users to create  profiles, update and get personnal details. 

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
- [ ] Tidy up both backend and frontend

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


