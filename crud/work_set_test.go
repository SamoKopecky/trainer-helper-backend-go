package crud

import (
	"fmt"
	"testing"
	"trainer-helper/db"
)

// func setupTestDB() *bun.DB {
// 	sqldb, err := sql.Open("sqlite3", "file::memory:?cache=shared")
// 	if err != nil {
// 		panic(err)
// 	}
// 	db := bun.NewDB(sqldb, sqlitedialect.New())
// 	driver, err := sqlite3.WithInstance(sqldb, &sqlite3.Config{})
//
// 	m, err := migrate.NewWithDatabaseInstance("file://../migrations", "sqlite3", driver)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize migrations: %v", err)
// 	}
// 	if err := m.Up(); err != nil {
// 		if err.Error() != "no change" {
// 			log.Fatalf("Failed to apply migrations: %v", err)
// 		} else {
// 			fmt.Println("No new migrations to apply")
// 		}
// 	} else {
// 		fmt.Println("Migrations applied successfully!")
// 	}
// 	return db
// }

func setupTestDb() {
	conn := db.GetDbConnTest(testing.Verbose())
	conn.RunMigrations()
}

func TestInsertMany(t *testing.T) {
	setupTestDb()
	fmt.Print("here")
}
