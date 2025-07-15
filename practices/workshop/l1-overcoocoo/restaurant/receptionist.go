package restaurant

import (
	"errors"
	"strconv"

	"github.com/fatih/color"
)

//  接待人員要接待剛進門的客人
//  一次只能接待一位

// Receptionist 接待人員
type Receptionist struct {
	rt        *Restaurant // 餐廳
	isServing bool        // 是否正在服務
}

// NewReceptionist 新接待人員
func NewReceptionist(rt *Restaurant) *Receptionist {
	return &Receptionist{
		rt: rt,
	}
}

// leadToSeat 帶位
func (r *Receptionist) leadToSeat(c *Customer) error {
	r.isServing = true
	defer r.backToStandBy()

	index, hasEmpty := r.IsHasEmptySeat()
	if hasEmpty {
		r.rt.currentSeats++
		r.rt.seats[index] = Seat{
			cID:      c.id,
			customer: c,
		}
		color.Magenta("接待人員帶位成功,客人:" + strconv.Itoa(int(c.id)))
		return nil
	}
	return errors.New("沒有座位了")
}

// IsHasEmptySeat 確認空座位
func (r *Receptionist) IsHasEmptySeat() (int, bool) {

	r.rt.seatRWlocker.Lock()
	defer r.rt.seatRWlocker.Unlock()
	if r.rt.currentSeats < r.rt.maxSeats {
		for i := 0; i < r.rt.maxSeats; i++ {
			if _, exists := r.rt.seats[i]; !exists {
				return i, true
			}
		}
	}
	return -1, false
}

// backToStandBy 帶位完畢
func (r *Receptionist) backToStandBy() {
	r.isServing = false
}
