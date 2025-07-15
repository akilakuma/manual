package main

import "testing"

func TestGetInfo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "測試",
			want: "plus",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInfo(); got != tt.want {
				t.Errorf("GetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
