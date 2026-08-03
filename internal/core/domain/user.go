package domain

type User struct {
	ID      int
	Version int

	Name        string
	Email       string
	PhoneNumber *string
}

func NewUser(id int, version int, name string, email string, phoneNumber *string) User {
	return User{
		ID:          id,
		Version:     version,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(name string, email string, phoneNumber *string) User {
	return NewUser(UninitializedID, UninitializedVersion, name, email, phoneNumber)
}
