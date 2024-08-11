<a name="readme-top"></a>
<h3 align="center">Concurrency using coffee and barista simulation</h3>

  <p align="center">
  Project Description:
  
  This project simulates a simple coffee shop operation where customers arrive, place orders, and baristas prepare and serve the orders. 
    <br />
    <a href="https://github.com/Justinljg/cofee-shop"><strong>Explore the docs »</strong></a>

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
        <li><a href="#containerising-and-running">Containerising and running</a></li>
        <li><a href="#testing">Testing</a></li>
      </ul>
    </li>
    <li>
      <a href="#summary-of-repository">Summary of Repository</a>
      <ul>
        <li><a href="#main">main</a></li>
        <li><a href="#barista">barista</a></li>
        <li><a href="#coffee">coffee.go</a></li>
        <li><a href="#customer">customer</a></li>
        <li><a href="#order">order</a></li>
      </ul>
    </li>
  </ol>
</details>


<!-- ABOUT THE PROJECT -->
## About The Project

`https://https://github.com/Justinljg/coffee-shop`

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

    $ git clone https://github.com/Justinljg/coffee-shop

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


## Containerising and running
This repository does not feature a running app but if your would like to run it in a container you can do the following. 

    docker build -t my-app .
Build the docker image.

    docker run -p 4000:80 my-app
Run the docker image.
<br></br>

Alternatively you can just run the go file.

    go run -race main.go

## Testing
The tests can be run through the command.

    go test -race -v ./...


## Summary of Repository
I am going to use a lot of my personal simplified terms here according to my understanding for my own reference which may differ from the standards.

The code tries to run the a structure where the operations is seperate from the logic. What this means in this repository is operational stuff like {e.g. goroutines to ensure concurrency} operations is seperate from {e.g. coffetype, waiting times} logic. This allows me to run mosts test without integrations or mocks for my unit tests. This is based on this [article](https://blog.boot.dev/clean-code/writing-good-unit-tests-dont-mock-database-connections/). The context of this article is more on database connections but I felt that the practice is also applicable here. 

The picture below is used to depict what the setup is like.


From my personal experience, it is important to take note where you implement your wait groups and where you send or close your channels. This can cause race conditions.

Race conditions happens in very simplified terms is when multiple functions access the same "thing"(can be variables or channels) concurrently. Your code may still be able to run in low loads but in high loads it may lead to crashes or data may not be written properly. So it is important to sync your channels or ensure variables are independent.

This repository uses go routines, buffered channels, range and close , waitgroups and select.

### main

Parsing the number of customer from a str to an int var to be used. 

Creates the channels and the number of Baristas.

Create the waitgroup, goroutines for things to run concurrently and in parallel with a cancel context to shutdown gracefully.

Parallelism is achieved through the two number of workers/baristas.

### barista

Loops through the customers.

Check for cancellation.

Logs customer orders.

Prepare the drinks and send to orders channel.

### coffee

Contains Coffee type, conversion/expression of coffee type through numbers and Coffee waiting time.

### customer

Loops through through a specified customer number with each customer having a random wiating time.

### order

Contains the struct for an order.

