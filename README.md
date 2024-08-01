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
      </ul>
    </li>
    <li>
      <a href="#summary-of-repository">Summary of Repository</a>
      <ul>
        <li><a href="#main.go">main.go</a></li>
        <li><a href="#barista.go">barista.go</a></li>
        <li><a href="#coffee.go">coffee.go</a></li>
        <li><a href="#customer.go">customer.go</a></li>
        <li><a href="#order.go">order.go</a></li>
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

### main.go

To be developed.

### barista.go

To be developed.

### coffee.go

To be developed.

### customer.go

To be developed.

### order.go



