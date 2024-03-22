package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Permissions []string

type PermissionsModel struct {
	DB *sql.DB
}

func (p Permissions) Include(code string) bool {
	for _, permission := range p {
		if permission == code {
			return true
		}
	}
	return false
}

func (m PermissionsModel) AddForUser(userID int64, codes ...string) error {
	query := `
INSERT INTO users_permissions
SELECT $1, permissions.id FROM permissions WHERE permissions.code = ANY($2)`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, query, pq.Array(codes))
	return err
}

func (m PermissionsModel) GetAllForUser(userID int64) (Permissions, error) {
	query :=
		`SELECT permissions.code
		FROM permissions
		INNER JOIN user_permissions ON user_permissions.permission_id = permissions.id
		INNER JOIN users ON users.id = user_permissions.user_id
		WHERE users.id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions Permissions

	for rows.Next() {
		var permission string
		err = rows.Scan(&permission)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
