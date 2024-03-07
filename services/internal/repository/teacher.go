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
	stmt := `INSERT INTO users (fullName, email, password, phoneNumber, role)
		VALUES ($1, $2, $3, $4, 'teacher')
		RETURNING ID`

	var teacherID int
	err := r.db.QueryRow(stmt, t.FullName, t.Email, t.Password, t.PhoneNumber).Scan(&teacherID)
	if err != nil {
		return 0, fmt.Errorf("error executing SQL statement: %v", err)
	}

	return teacherID, nil
}

func (r *TeacherRepo) GetOne(id int) (domain.Teacher, error) {
	stmt := `SELECT ID, FullName, Email, Password, PhoneNumber, role
	FROM users
	WHERE ID = $1`

	var teacher domain.Teacher
	var role string
	err := r.db.QueryRow(stmt, id).Scan(&teacher.ID, &teacher.FullName, &teacher.Email, &teacher.Password, &teacher.PhoneNumber, &role)

	if err != nil {
		return domain.Teacher{}, err
	}

	if role != "teacher" {
		return domain.Teacher{}, fmt.Errorf("teacher with ID %d not found", id)
	}

	return teacher, nil
}

func (r *TeacherRepo) GetAll() ([]domain.Teacher, error) {
	stmt := `SELECT ID, FullName, Email, Password, PhoneNumber
	FROM users
	WHERE role='teacher'`

	var teachers []domain.Teacher

	rows, err := r.db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var teacher domain.Teacher
		err := rows.Scan(&teacher.ID, &teacher.FullName, &teacher.Email, &teacher.Password, &teacher.PhoneNumber)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}
