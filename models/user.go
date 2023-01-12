package models

type User struct {
	Name           string
	Age            int
	CurrentSavings float64
}

func (u *User) Create() (*User, error) {
	return &User{
		Name:           "John",
		Age:            30,
		CurrentSavings: 1000.00,
	}, nil
}
