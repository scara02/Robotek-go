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

func (r *TeacherRepo) Delete(id int) (int, error) {
	stmt := `DELETE FROM users
	WHERE id = $1 AND role = 'teacher'
	RETURNING id`

	var deletedID int
	err := r.db.QueryRow(stmt, id).Scan(&deletedID)

	if err != nil {
		return 0, err
	}

	return deletedID, nil
}

func (r *TeacherRepo) AddToGroup(teacherID, groupID int) error {
	stmt := `INSERT INTO teacher_groups (teacherID, groupID)
	VALUES ($1, $2)`

	_, err := r.db.Exec(stmt, teacherID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TeacherRepo) GetGroups(id int) ([]domain.Group, error) {
	stmt := `SELECT g.ID, g.GroupName
	FROM teacher_groups t
	JOIN groups g ON t.groupID=g.id
	WHERE t.teacherID = $1`

	rows, err := r.db.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []domain.Group

	for rows.Next() {
		var group domain.Group
		if err := rows.Scan(&group.ID, &group.GroupName); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *TeacherRepo) DeleteGroup(teacherID, groupID int) error {
	stmt := `DELETE FROM teacher_groups
	WHERE teacherID = $1 and groupID = $2`

	_, err := r.db.Exec(stmt, teacherID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TeacherRepo) Update(id int, updatedTeacher *domain.Teacher) error {
	stmtCheckRole := `SELECT role 
	FROM users 
	WHERE ID=$1`

	var role string

	err := r.db.QueryRow(stmtCheckRole, id).Scan(&role)
	if err != nil {
		return err
	}

	if role != "teacher" {
		return fmt.Errorf("user with ID %d is not a teacher", id)
	}

	stmt := `UPDATE users
	SET 
	FullName=COALESCE(NULLIF($2, ''), FullName), 
	Email=COALESCE(NULLIF($3, ''), Email), 
	Password=COALESCE(NULLIF($4, ''), Password), 
	PhoneNumber=COALESCE(NULLIF($5, ''), PhoneNumber)
	WHERE ID=$1`

	_, err = r.db.Exec(stmt, id, updatedTeacher.FullName, updatedTeacher.Email, updatedTeacher.Password, updatedTeacher.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}