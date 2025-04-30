package repo

import (
	"context"
	"database/sql"
	"gowiki/internal/db"
	"gowiki/internal/logger"
	model "gowiki/internal/model"
)

func SaveUser(ctx context.Context, tx *sql.Tx, user *model.User) error {
	query := `INSERT INTO users (username, email, password_hash)
	          VALUES ($1, $2, $3)
	          RETURNING id, create_at`
	err := db.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.PasswordHash).
		Scan(&user.ID, &user.CreateAt)
	if err != nil {
		logger.Error.Printf("Ошибка при сохранении пользователя (email=%s: %v", user.Email, err)
		return err
	}
	return nil
}

func GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (*model.User, error) {
	row := db.DB.QueryRowContext(ctx, "SELECT id, username, email, password_hash, create_at FROM users WHERE email = $1", email)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreateAt)
	if err != nil {
		logger.Error.Printf("Ошибка при получении пользователя по почте (email=%s): %v", email, err)
		return nil, err
	}
	return user, nil
}
