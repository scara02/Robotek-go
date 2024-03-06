package repository

import (
	"database/sql"
	"fmt"

	"go-web-robotek/services/internal/domain"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{
		db: db,
	}
}

func (r *StudentRepo) Create(s *domain.Student) (int, error) {
	stmt := `INSERT INTO students (fullName, email, password, phoneNumber, groupID)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING ID`

	var studentID int
	err := r.db.QueryRow(stmt, s.FullName, s.Email, s.Password, s.PhoneNumber, s.GroupID).Scan(&studentID)
	if err != nil {
		return 0, fmt.Errorf("error executing SQL statement: %v", err)
	}

	return studentID, nil
}

func (r *StudentRepo) GetOne(id int) (domain.Student, error) {
	stmt := `SELECT ID, FullName, Email, Password, PhoneNumber, GroupID
	FROM students
	WHERE ID = $1`

	var student domain.Student
	err := r.db.QueryRow(stmt, id).Scan(&student.ID, &student.FullName, &student.Email, &student.Password, &student.PhoneNumber, &student.GroupID)

	if err != nil {
		return domain.Student{}, err
	}

	return student, nil
}

func (r *StudentRepo) GetAll() ([]domain.Student, error) {
	stmt := `SELECT ID, FullName, Email, Password, PhoneNumber, GroupID
	FROM students`

	var students []domain.Student

	rows, err := r.db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student domain.Student
		err := rows.Scan(&student.ID, &student.FullName, &student.Email, &student.Password, &student.PhoneNumber, &student.GroupID)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepo) GetGroup(id int) (domain.Group, error) {
	stmt := `SELECT ID, GroupName
	FROM groups
	WHERE ID = $1`

	var group domain.Group
	err := r.db.QueryRow(stmt, id).Scan(&group.ID, &group.GroupName)

	if err != nil {
		return domain.Group{}, err
	}

	return group, nil
}
