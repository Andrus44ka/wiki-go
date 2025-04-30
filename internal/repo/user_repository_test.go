package repo_test

import (
	"context"
	"gowiki/internal/db"
	"gowiki/internal/logger"
	"gowiki/internal/model"
	"gowiki/internal/repo"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logger.Init() // ← обязательно ДО db.Init()
	db.Init()
	code := m.Run()
	os.Exit(code)
}

func TestSaveUser(t *testing.T) {
	ctx := context.Background()

	// Начинаем транзакцию
	tx, err := db.DB.Begin()
	if err != nil {
		t.Fatalf("не удалось начать транзакцию: %v", err)
	}
	defer tx.Rollback() // откатить транзакцию в конце теста

	// Создаем пользователя
	user := &model.User{
		Username:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
	}

	// Сохраняем пользователя в транзакции
	err = repo.SaveUser(ctx, tx, user)
	if err != nil {
		t.Fatalf("ошибка при сохранении пользователя: %v", err)
	}

	// Проверяем, что ID был присвоен
	if user.ID == 0 {
		t.Fatalf("ожидался ID после сохранения, но он остался 0")
	}

	// Проверяем, что пользователь существует в базе данных (например, с помощью GetUserByEmail)
	savedUser, err := repo.GetUserByEmail(ctx, tx, user.Email)
	if err != nil {
		t.Fatalf("не удалось получить пользователя из базы данных: %v", err)
	}
	if savedUser.ID != user.ID {
		t.Fatalf("ожидался ID: %v, но получили ID: %v", user.ID, savedUser.ID)
	}
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	tx, err := db.DB.Begin()
	if err != nil {
		t.Fatalf("не удалось начать транзакцию: %v", err)
	}
	defer tx.Rollback()

	email := "test@example.com"

	user, err := repo.GetUserByEmail(ctx, tx, email)
	if err != nil {
		t.Fatalf("ошибка при получении пользователя: %v", err)
	}

	if user == nil {
		t.Fatal("пользователь не найден (user == nil)")
	}

	if user.Email != email {
		t.Fatalf("ожидался email %s, но получен %s", email, user.Email)
	}
}
