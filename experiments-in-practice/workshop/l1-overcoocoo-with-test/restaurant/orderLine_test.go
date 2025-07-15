package restaurant

import (
	"reflect"
	"testing"
)

func TestNewOrderLine(t *testing.T) {
	type args struct {
		chiefs []*Chief
	}
	tests := []struct {
		name string
		args args
		want OrderLine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderLine(tt.args.chiefs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderLine_income(t *testing.T) {
	type args struct {
		seat int
		cID  int32
		food string
	}
	tests := []struct {
		name string
		o    *OrderLine
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.income(tt.args.seat, tt.args.cID, tt.args.food)
		})
	}
}
