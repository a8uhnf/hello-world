package model


type User struct {
	ID   int
	Name string
	Role string
}

type Users []User

func (u Users) Exists(id int) bool {
	return true
}

func (u Users) FindByName(name string) (User, error) {
	return User{}, nil
}

