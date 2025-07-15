package restaurant

import (
	"reflect"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type args struct {
		food string
	}
	tests := []struct {
		name string
		args args
		want *Customer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomer(tt.args.food); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomer_SetID(t *testing.T) {
	type args struct {
		id int32
	}
	tests := []struct {
		name string
		c    *Customer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetID(tt.args.id)
		})
	}
}

func TestCustomer_order(t *testing.T) {
	tests := []struct {
		name string
		c    *Customer
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.order(); got != tt.want {
				t.Errorf("Customer.order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomer_eating(t *testing.T) {
	type args struct {
		food string
	}
	tests := []struct {
		name string
		c    *Customer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.eating(tt.args.food)
		})
	}
}

func TestCustomer_leave(t *testing.T) {
	tests := []struct {
		name string
		c    *Customer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.leave()
		})
	}
}
