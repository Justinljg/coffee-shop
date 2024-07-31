package cafe

import (
	"sync"
	"testing"
	"time"

	"github.com/justinljg/coffee-shop/cafe"
)

func TestBaristaPrepareOrder(t *testing.T) {
	orders := make(chan cafe.Order, 10)
	wg := sync.WaitGroup{}

	barista := cafe.Barista{ID: 1}
	order := cafe.Order{CustomerID: 1, CoffeeType: cafe.Espresso}

	wg.Add(1)
	startTime := time.Now()
	go barista.PrepareOrder(order, orders, &wg)
	wg.Wait()

	duration := time.Since(startTime)
	if duration < cafe.CoffeePreparationTimes[cafe.Espresso] {
		t.Errorf("Expected at least %v for preparation, got %v", cafe.CoffeePreparationTimes[cafe.Espresso], duration)
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
