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
	stmt := `INSERT INTO users (fullName, email, password, phoneNumber, role)
		VALUES ($1, $2, $3, $4, 'student')
		RETURNING ID`

	var studentID int
	err := r.db.QueryRow(stmt, s.FullName, s.Email, s.Password, s.PhoneNumber).Scan(&studentID)
	
	if err != nil {
		return 0, fmt.Errorf("error executing SQL statement: %v", err)
	}

	stmt = `INSERT INTO student_group (studentID, groupID)
	VALUES ($1, $2)`

	_, err = r.db.Exec(stmt, studentID, s.GroupID)

	if err != nil {
		return 0, fmt.Errorf("error executing SQL statement: %v", err)
	}

	return studentID, nil
}

func (r *StudentRepo) GetOne(id int) (domain.Student, error) {
	stmt := `SELECT u.ID, u.FullName, u.Email, u.Password, u.PhoneNumber, s.GroupID
	FROM users u
	JOIN student_group s on u.id=s.studentID
	WHERE u.ID = $1`

	var student domain.Student
	err := r.db.QueryRow(stmt, id).Scan(&student.ID, &student.FullName, &student.Email, &student.Password, &student.PhoneNumber, &student.GroupID)

	if err != nil {
		return domain.Student{}, err
	}

	return student, nil
}

func (r *StudentRepo) GetAll() ([]domain.Student, error) {
	stmt := `SELECT u.ID, u.FullName, u.Email, u.Password, u.PhoneNumber, s.GroupID
	FROM users u
	JOIN student_group s ON u.id=s.studentID`

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
	stmt := `SELECT g.ID, g.GroupName
	FROM student_group s
	JOIN groups g ON s.groupID=g.id
	WHERE s.studentID = $1`

	var group domain.Group
	err := r.db.QueryRow(stmt, id).Scan(&group.ID, &group.GroupName)

	if err != nil {
		return domain.Group{}, err
	}

	return group, nil
}


func (r *StudentRepo) Delete(id int) (int, error) {
	stmt := `DELETE FROM users
	WHERE id = $1 AND role = 'student'
	RETURNING id`

	var deletedID int
	err := r.db.QueryRow(stmt, id).Scan(&deletedID)

	if err != nil {
		return 0, err
	}

	return deletedID, nil
}

func (r *StudentRepo) ChangeGroup(studentID, groupID int) error {
	stmt := `UPDATE student_group
	SET groupID = $1
	WHERE studentID = $2`

	_, err := r.db.Exec(stmt, groupID, studentID)
	if err != nil {
		return err
	}

	return nil
}

func (r *StudentRepo) Update(id int, updatedStudent *domain.Student) error {
	stmtCheckRole := `SELECT role 
	FROM users 
	WHERE ID=$1`

	var role string

	err := r.db.QueryRow(stmtCheckRole, id).Scan(&role)
	if err != nil {
		return err
	}

	if role != "student" {
		return fmt.Errorf("user with ID %d is not a student", id)
	}

	stmt := `UPDATE users
	SET 
	FullName=COALESCE(NULLIF($2, ''), FullName), 
	Email=COALESCE(NULLIF($3, ''), Email), 
	Password=COALESCE(NULLIF($4, ''), Password), 
	PhoneNumber=COALESCE(NULLIF($5, ''), PhoneNumber)
	WHERE ID=$1`

	_, err = r.db.Exec(stmt, id, updatedStudent.FullName, updatedStudent.Email, updatedStudent.Password, updatedStudent.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}