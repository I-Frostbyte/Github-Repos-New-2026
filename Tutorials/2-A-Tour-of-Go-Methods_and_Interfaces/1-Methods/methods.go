package methods

import "math"

type Vertex struct {
	X, Y float64
}

type User struct {
	name string
	email string
	phone_number string
	age int64
	address string
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func NewUser(
	name, 
	email, 
	phone_number, 
	address string, 
	age int64,
	) User {
	return User{
		name: name,
		email: email,
		phone_number: phone_number,
		age: age,
		address: address,
	}
}

func (u User) GetName(user User) string {
	return user.name
}