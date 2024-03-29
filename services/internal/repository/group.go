package repository

import (
	"database/sql"
	"fmt"

	"go-web-robotek/services/internal/domain"
)

type GroupRepo struct {
	db *sql.DB
}

func NewGroupRepo(db *sql.DB) *GroupRepo {
	return &GroupRepo{
		db: db,
	}
}

func (r *GroupRepo) Create(g *domain.Group) (int, error) {
	stmt := `INSERT INTO groups (groupName)
		VALUES ($1)
		RETURNING ID`

	var groupID int
	err := r.db.QueryRow(stmt, g.GroupName).Scan(&groupID)
	if err != nil {
		return 0, fmt.Errorf("error executing SQL statement: %v", err)
	}

	return groupID, nil
}

func (r *GroupRepo) GetOne(id int) (domain.Group, error) {
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

func (r *GroupRepo) GetAll() ([]domain.Group, error) {
	stmt := `SELECT ID, GroupName
	FROM groups`

	var groups []domain.Group

	rows, err := r.db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group domain.Group
		err := rows.Scan(&group.ID, &group.GroupName)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *GroupRepo) Delete(id int) (int, error) {
	stmt := `DELETE FROM groups
	WHERE id = $1
	RETURNING id`

	var deletedID int
	err := r.db.QueryRow(stmt, id).Scan(&deletedID)

	if err != nil {
		return 0, err
	}

	return deletedID, nil
}

func (r *GroupRepo) Update(id int, updatedGroup *domain.Group) error {
	stmt := `UPDATE groups
	SET GroupName=COALESCE(NULLIF($2, ''), GroupName)
	WHERE id = $1`

	_, err := r.db.Exec(stmt, id, updatedGroup.GroupName)

	if err != nil {
		return err
	}

	return nil
}