package coffee

import (
	"math/rand"
	"sync"
	"time"
)

type Customer struct {
	Id int
}

func SimulateCustomerArrivals(customers chan<- Customer, wg *sync.WaitGroup) {
	defer wg.Done()
	customerID := 1
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		customers <- Customer{Id: customerID}
		customerID++
	}
}
