package restaurant

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewReceptionist(t *testing.T) {
	type args struct {
		rt *Restaurant
	}
	tests := []struct {
		name string
		args args
		want *Receptionist
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReceptionist(tt.args.rt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReceptionist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceptionist_leadToSeat(t *testing.T) {

	mockRestaurant1 := &Restaurant{
		seatRWlocker: new(sync.RWMutex),
		currentSeats: 0,
		maxSeats:     1,
		seats:        make(map[int]Seat, 1),
	}
	mockReceptionist1 := &Receptionist{
		rt: mockRestaurant1,
	}

	mockRestaurant2 := &Restaurant{
		seatRWlocker: new(sync.RWMutex),
		currentSeats: 1,
		maxSeats:     1,
		seats:        make(map[int]Seat, 1),
	}
	mockReceptionist2 := &Receptionist{
		rt: mockRestaurant2,
	}

	mockReceptionist3 := &Receptionist{}

	mockCustomer := &Customer{}

	type args struct {
		c *Customer
	}
	tests := []struct {
		name    string
		r       *Receptionist
		args    args
		wantErr bool
	}{
		{
			name: "有位置可以正常帶位",
			r:    mockReceptionist1,
			args: args{
				c: mockCustomer,
			},
			wantErr: false,
		},
		{
			name: "位置已滿無法帶位",
			r:    mockReceptionist2,
			args: args{
				c: mockCustomer,
			},
			wantErr: true,
		},
		{
			name: "例外狀況，接待人員沒有所在的餐廳",
			r:    mockReceptionist3,
			args: args{
				c: mockCustomer,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.leadToSeat(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Receptionist.leadToSeat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReceptionist_IsHasEmptySeat(t *testing.T) {
	tests := []struct {
		name  string
		r     *Receptionist
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.IsHasEmptySeat()
			if got != tt.want {
				t.Errorf("Receptionist.IsHasEmptySeat() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Receptionist.IsHasEmptySeat() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReceptionist_backToStandBy(t *testing.T) {
	tests := []struct {
		name string
		r    *Receptionist
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.backToStandBy()
		})
	}
}
