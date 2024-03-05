package repository

import (
	"database/sql"
	"fmt"

	"go-web-robotek/services/internal/domain"
)

type TeacherRepo struct {
	db *sql.DB
}

func NewTeacherRepo(db *sql.DB) *TeacherRepo {
	return &TeacherRepo{
		db: db,
	}
}

func (r *TeacherRepo) Create(t *domain.Teacher) (int, error) {
	stmt := `INSERT INTO teachers (fullName, email, password, phoneNumber)
		VALUES ($1, $2, $3, $4)
		RETURNING ID`

	var teacherID int
	err := r.db.QueryRow(stmt, t.FullName, t.Email, t.Password, t.PhoneNumber).Scan(&teacherID)
	if err != nil {
		return 0, fmt.Errorf("error executing SQL statement: %v", err)
	}

	return teacherID, nil
}

func (r *TeacherRepo) GetByID(id int) (domain.Teacher, error) {
	stmt := `SELECT ID, FullName, Email, Password, PhoneNumber
	FROM teachers
	WHERE ID = $1`

	var teacher domain.Teacher
	err := r.db.QueryRow(stmt, id).Scan(&teacher.ID, &teacher.FullName, &teacher.Email, &teacher.Password, &teacher.PhoneNumber)

	if err != nil {
        return domain.Teacher{}, err
    }

    return teacher, nil
}
