package main

import "fmt"

func main() {

	var t1 = &Teacher{
		lession: "english",
	}

	var u1 = User{
		id:   1,
		name: "pikachu",
		t:    t1,
	}

	fmt.Println(&u1.id, u1.GetName, u1.GetName)
	u1.GetName()
	u1.GetName()
	u1.GetName()

	var u2 = User{
		id:   2,
		name: "rachu",
		t:    t1,
	}

	fmt.Println(&u2.id, u2.GetName2, u2.GetName2)
	u2.GetName2()
	u2.GetName2()
	u2.GetName2()
}

// User 使用者
type User struct {
	id   int
	name string
	t    *Teacher
}

// Teacher 老師
type Teacher struct {
	lession string
}

// GetName 取得使用者名稱
func (u User) GetName() string {
	fmt.Println("🐌🐌", &u, &u.id, &u.name, &u.t, &u.t.lession)
	return u.name
}

// GetName2 取得使用者名稱2
func (u *User) GetName2() string {
	fmt.Println("🐛🐛", &u, &u.id, &u.name, &u.t, &u.t.lession)
	return u.name
}
