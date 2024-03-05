package domain

type Teacher struct {
	ID          int
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
}

func NewTeacher(id int, fullName, email, password, phoneNumber string) Teacher {
	return Teacher{
		ID:          id,
		FullName:    fullName,
		Email: email,
		Password: password,
		PhoneNumber: phoneNumber,
	}
}
