package domain

type User struct {
	ID    string
	Email string
}

func NewUser(email string) *User {
	return &User{
		ID:    generateID(email),
		Email: email,
	}

}

func generateID(email string) string {
	// 這裡可以使用 UUID 或其他方式生成唯一 ID
	return "some-unique-id" + email
}
