package db

import (
  "log"
  "fmt"
  "runtime"
  "path/filepath"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/Sirupsen/logrus"
  "github.com/DavidHuie/gomigrate"
)

func RunMigrations(db *sql.DB) {
  _, b, _, _ := runtime.Caller(0)
  basepath := filepath.Dir(b)

  migrator, _ := gomigrate.NewMigratorWithLogger(db, gomigrate.Postgres{}, fmt.Sprintf("%s/migrations", basepath), logrus.New())
  err := migrator.Migrate()

  if err != nil {
    log.Fatal(err)
  }
}

func ConnectDB() *sql.DB {
  db, err := sql.Open("postgres", "postgres://user:mdp@localhost:5555/simple-webapp?sslmode=disable")

  if err != nil {
    log.Fatal(err)
  }

  RunMigrations(db)
  return db
}
