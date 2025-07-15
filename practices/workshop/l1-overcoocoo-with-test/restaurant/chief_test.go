package restaurant

import (
	"reflect"
	"testing"
)

func TestNewChief(t *testing.T) {
	tests := []struct {
		name string
		want *Chief
	}{
		// {
			// name: "正常的建立一個廚師",
			// want: &Chief{
			// 	job:    make(chan string, 100),
			// 	onDuty: make(chan bool),
			// },
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChief(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChief() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChief_setOutputFoodTable(t *testing.T) {
	type args struct {
		table chan string
	}
	tests := []struct {
		name string
		c    *Chief
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.setOutputFoodTable(tt.args.table)
		})
	}
}

func TestChief_getOrder(t *testing.T) {
	type args struct {
		jobs <-chan string
	}
	tests := []struct {
		name string
		c    *Chief
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.getOrder(tt.args.jobs)
		})
	}
}

func TestChief_cook(t *testing.T) {
	type args struct {
		food string
	}
	tests := []struct {
		name string
		c    *Chief
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.cook(tt.args.food)
		})
	}
}

func TestChief_sendOut(t *testing.T) {
	type args struct {
		food string
	}
	tests := []struct {
		name string
		c    *Chief
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.sendOut(tt.args.food)
		})
	}
}

func TestChief_takeBreak(t *testing.T) {
	tests := []struct {
		name string
		c    *Chief
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.takeBreak()
		})
	}
}

func TestChief_Ｗorking(t *testing.T) {
	tests := []struct {
		name string
		c    *Chief
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Ｗorking()
		})
	}
}
