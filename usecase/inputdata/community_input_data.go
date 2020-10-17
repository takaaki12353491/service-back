package inputdata

type Community struct {
	UserID      string `validate:"required"`
	Name        string `validate:"required"`
	Description string
}
