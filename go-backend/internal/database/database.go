package database

import (
  "fmt"
  "github.com/fikshun/go-backend/internal/api"
  "github.com/spf13/viper"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
)

func InitDB() *gorm.DB {
  viper.AutomaticEnv()
  viper.SetDefault("POSTGRES_HOST", "localhost")
  viper.SetDefault("POSTGRES_USER", "boilerplate")
  viper.SetDefault("POSTGRES_PASSWORD", "changeme")
  viper.SetDefault("POSTGRES_DB", "boilerplate")
  viper.SetDefault("POSTGRES_PORT", "5432")

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    viper.GetString("POSTGRES_HOST"),
    viper.GetString("POSTGRES_USER"),
    viper.GetString("POSTGRES_PASSWORD"),
    viper.GetString("POSTGRES_DB"),
    viper.GetString("POSTGRES_PORT"))

  postgresConfig := postgres.Config{DSN: dsn}

  db, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{})

  if err != nil {
    log.Fatal(err)
  }

  return db
}

func Migrate(db *gorm.DB) {
  err := db.AutoMigrate(
    &api.User{},
  )
  if err != nil {
    log.Fatal(err)
  }
}
