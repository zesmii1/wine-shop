package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DB_URL")

	// подключение для миграции
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("Ошибка создания драйвера миграций:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Ошибка инициализации мигратора:", err)
	}

	if version, dirty, err := m.Version(); err == nil && dirty {
		log.Printf("Обнаружено грязное состояние версии %d, исправляем...", version)
		if err := m.Force(int(version)); err != nil {
			log.Fatal("Ошибка исправления грязного состояния:", err)
		}
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Ошибка применения миграций:", err)
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка инициализации GORM:", err)
	}

	log.Println("База данных подключена и миграции применены!")
}
