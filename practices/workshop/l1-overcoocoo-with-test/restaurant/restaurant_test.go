package restaurant

import (
	"reflect"
	"testing"
)

func TestCreateRestaurant(t *testing.T) {
	tests := []struct {
		name string
		want *Restaurant
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRestaurant(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestaurant_NewCustomerID(t *testing.T) {
	tests := []struct {
		name string
		r    *Restaurant
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.NewCustomerID(); got != tt.want {
				t.Errorf("Restaurant.NewCustomerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestaurant_workerOnDuty(t *testing.T) {
	tests := []struct {
		name string
		r    *Restaurant
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.workerOnDuty()
		})
	}
}

func TestRestaurant_CustomerComein(t *testing.T) {
	type args struct {
		c *Customer
	}
	tests := []struct {
		name    string
		r       *Restaurant
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CustomerComein(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Restaurant.CustomerComein() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
