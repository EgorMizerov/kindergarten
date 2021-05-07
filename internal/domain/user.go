package domain

type User struct {
	ID string

	PasswordHash string
	Email        string
	Phone        string

	FirstName    string
	MiddleName   string
	LastName     string
	DateOfBirth  int64
	IsAdmin      bool
	CreationDate int64
}
