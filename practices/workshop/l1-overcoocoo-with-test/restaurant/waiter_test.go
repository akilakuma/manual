package restaurant

import (
	"reflect"
	"testing"
)

func TestNewWaiter(t *testing.T) {
	type args struct {
		rt *Restaurant
	}
	tests := []struct {
		name string
		args args
		want *Waiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWaiter(tt.args.rt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWaiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaiter_serveOfCustomerOrder(t *testing.T) {
	tests := []struct {
		name string
		w    *Waiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.serveOfCustomerOrder()
		})
	}
}

func TestWaiter_serveOfOrderToChief(t *testing.T) {
	type args struct {
		seat int
		cID  int32
		food string
	}
	tests := []struct {
		name string
		w    *Waiter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.serveOfOrderToChief(tt.args.seat, tt.args.cID, tt.args.food)
		})
	}
}

func TestWaiter_serveOfDeliverFood(t *testing.T) {
	tests := []struct {
		name string
		w    *Waiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.serveOfDeliverFood()
		})
	}
}

func TestWaiter_checkFoodForCustomer(t *testing.T) {
	type args struct {
		food string
	}
	tests := []struct {
		name string
		w    *Waiter
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.checkFoodForCustomer(tt.args.food); got != tt.want {
				t.Errorf("Waiter.checkFoodForCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaiter_checkCustomerLeave(t *testing.T) {
	tests := []struct {
		name string
		w    *Waiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.checkCustomerLeave()
		})
	}
}

func TestWaiter_working(t *testing.T) {
	tests := []struct {
		name string
		w    *Waiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.working()
		})
	}
}
