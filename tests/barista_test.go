package coffee

import (
	"sync"
	"testing"
	"time"

	"github.com/justinljg/coffee-shop/coffee"
)

func TestBaristaPrepareOrder(t *testing.T) {
	orders := make(chan coffee.Order, 100)
	wg := sync.WaitGroup{}

	barista := coffee.Barista{Id: 1}
	order := coffee.Order{CustomerID: 1, CoffeeType: coffee.Espresso}

	wg.Add(1)
	startTime := time.Now()
	go barista.PrepareOrder(order, orders, &wg)
	wg.Wait()

	duration := time.Since(startTime)
	if duration < coffee.CoffeePreparationTimes[coffee.Espresso] {
		t.Errorf("Expected at least %v for preparation, got %v", coffee.CoffeePreparationTimes[coffee.Espresso], duration)
	}

	select {
	case o := <-orders:
		if o != order {
			t.Errorf("Expected order %v, got %v", order, o)
		}
	default:
		t.Error("Expected order to be completed")
	}
}
