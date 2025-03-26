package repository_test

import (
    "testing"
    "auth-service/internal/domain"
    "auth-service/internal/infrastructure/repository"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func TestUserRepository_Create(t *testing.T) {
    // Configurar mock de DB
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock DB: %v", err)
    }
    defer db.Close()

    gormDB, err := gorm.Open(postgres.New(postgres.Config{
        Conn: db,
    }), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }

    repo := repository.NewUserRepository(gormDB)

    t.Run("creación exitosa", func(t *testing.T) {
        mock.ExpectBegin()
        mock.ExpectQuery(`INSERT INTO "users"`).
            WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("123"))
        mock.ExpectCommit()

        user := &domain.User{Email: "test@example.com"}
        err := repo.Create(user)
        
        assert.NoError(t, err)
        assert.Equal(t, "123", user.ID)
    })
}