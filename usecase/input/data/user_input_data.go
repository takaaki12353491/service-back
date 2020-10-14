package inputdata

type User struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type Login struct {
	Identity string
	Password string
}
