<a name="readme-top"></a>
<h3 align="center">Concurrency using coffee and barista simulation</h3>

  <p align="center">
  Project Description:
  
  This project simulates a simple coffee shop operation where customers arrive, place orders, and baristas prepare and serve the orders. 
    <br />
    <a href="https://github.com/Justinljg/SupersetK8s/"><strong>Explore the docs »</strong></a>

</div>



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
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#working-tree">Working Tree</a></li>
        <li><a href="#running-the-repository-&-testing">Running the repository & Testing</a></li>
      </ul>
    </li>
    <li>
      <a href="#summary-of-repository">Summary of Repository</a></li>
        <ul>
        <li><a href="#main">Main</a></li>
        <li><a href="#sqlc">SQLC</a></li>
        <li><a href="#routes">Routes</a></li>
        <li><a href="#handlers">Handlers</a></li>
      </ul>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

`https://https://github.com/Justinljg/go-assignment`

This Project includes go files to mimic customers and baristas.

`justinljg`, `GO`,`GO concurrency`

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With


`GO`


![image](https://github.com/user-attachments/assets/c2775a5d-dbfc-4b8b-8872-a99ffa11a9ce)


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Open Terminal.

Change the current working directory to the location where you want the cloned directory.

Type git clone, and then paste the URL you copied earlier.

    $ git clone https://github.com/Justinljg/go-assignment

More specific instructions can be seen in https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository if needed.

### Prerequisites

GO, Docker has to be installed.

### Working Tree

The following is the working tree of this repository.

    .
    ├── cafe
    │   ├── barista.go
    │   ├── coffee.go
    │   ├── customer.go
    │   └── order.go
    ├── Dockerfile
    ├── go.mod
    ├── lint.sh
    ├── main.go
    ├── README.md
    └── tests
        ├── barista_test.go
        ├── coffee_test.go
        ├── customer_test.go
        └── order_test.go


<!-- USAGE EXAMPLES -->
## Containerising and running
This repository does not feature a running app but if your would like to run it in a container you can do the following. 

    docker build -t my-app .
Build the docker image.

    docker run -p 4000:80 my-app
Run the docker image.
<br></br>

Alternatively you can just run the go file.

    go run main.go
## Summary of Repository

### Main

<h4>Key Components</h4>
<br>
- Random Number Generator:

  Initializes a random number generator to create random coffee orders.
      
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
</br>
- Channels:

  These channels are used for customer arrivals and order handling. They are buffered channels with a capacity of 10.

    customers := make(chan cafe.Customer, 10)
    orders := make(chan cafe.Order, 10)

- WaitGroup:
  A WaitGroup is used to wait for the goroutines to finish their work before shutting down the program.

      var wg sync.WaitGroup

- Signal Handling:
  This creates a stop signal ctr+c from os to stop the customers.

      stop := make(chan os.Signal, 1)
      signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

- Customer Arrivals:
  A goroutine simulates customer arrivals by sending cafe.Customer objects to the customers channel. The WaitGroup is incremented to track this goroutine.

      go cafe.SimulateCustomerArrivals(customers, &wg)

- Order Handling by Baristas:

  For each barista in the baristas slice, a goroutine is started to handle customer orders:

      go func(b cafe.Barista) {
        for customer := range customers {
          order := cafe.Order{CustomerID: customer.ID, CoffeeType: cafe.CoffeeType(rng.Intn(3))}
          fmt.Printf("Customer %d arrives and places an order for a %s.\n", customer.ID, cafe.CoffeeTypeToString(order.CoffeeType))
          wg.Add(1)
          go b.PrepareOrder(order, orders, &wg)
        }
        wg.Done()
      }(barista)

Each barista listens on the customers channel for incoming customers, creates an order, and starts a new goroutine to prepare the order. The WaitGroup is incremented to track these order preparation goroutines.

- Graceful Shutdown:

  The program waits for a stop signal (<-stop).
  On receiving the stop signal, it closes the customers channel and waits for all goroutines to complete (wg.Wait()).
  After all customers are processed, the orders channel is closed, and any remaining orders are processed and printed.

### barista.go
This file contains structs needed and a PrepareOrder function.

The Barista struct represents a barista with an ID who prepares coffee orders.

The PrepareOrder method simulates a barista preparing a coffee order. It logs the start and end of the preparation, simulates the preparation time, and sends the completed order to the provided channel. 
        
### coffee.go
This file contains structs for the type of coffee, the time needed for each coffee and the type of coffee to string. The type of coffee is simplified by enumeration through ioata.

### customer.go

This file contains structs for customer with an ID and the SimulateCustomerArrivals function.

The SimulateCustomerArrivals function continuously generates customer arrivals and sends them to the customers channel. This function runs in a separate goroutine and increments the customer ID for each new customer. The function continues to run until the parent goroutine signals completion by closing the channel.

### customer.go

This file contains struct for the customer id and the coffee type.


