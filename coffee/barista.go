package coffee

import (
	"fmt"
	"sync"
	"time"
)

type Barista struct {
	Id int
}

func (barista *Barista) PrepareOrder(order Order, orders chan<- Order, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Barista %d starts preparing %s for Customer %d.\n", barista.Id, CoffeeTypeToString(order.CoffeeType), order.CustomerID)
	time.Sleep(CoffeePreparationTimes[order.CoffeeType])
	fmt.Printf("Barista %d finishes %s for Customer %d.\n", barista.Id, CoffeeTypeToString(order.CoffeeType), order.CustomerID)
	orders <- order
}
