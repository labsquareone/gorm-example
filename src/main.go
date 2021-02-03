package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitialMigration for project
func InitialMigration() {
	dsn := "host=localhost user=taedongkim password=12345 dbname=gorm port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	defer sqlDB.Close()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection ay be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		panic(err.Error())
	}

	db.Migrator().DropTable(&User{})

	// Migrate the schema
	db.AutoMigrate((&User{}))

	// Create
	user := User{
		Username:  "ted",
		FirstName: "Ted",
		LastName:  "Kim",
	}
	db.Create(&user)
	db.Create(&User{Username: "user2", FirstName: "First Name", LastName: "Last Name"})
	db.Create(&User{Username: "user3", FirstName: "First Name", LastName: "Last Name"})

	// Read
	u := User{}
	u2 := User{}
	db.First(&u)
	db.First(&u2, "first_name = ?", "First Name")

	var newUser User
	db.Last(&newUser)

	fmt.Println(u)
	fmt.Println(u2)
	fmt.Println(newUser)

	u = User{Username: "user3"}
	db.Where(&u).First(&u)
	fmt.Println(u) //user3

	// Update
	db.Model(&u2).Update("last_name", "kkkkk")
	fmt.Println(u2)
	// Update multiple fields
	db.Model(&u2).Updates(User{FirstName: "new", LastName: "new2"})
	fmt.Println(u2)
	db.Model(&u2).Updates(map[string]interface{}{"first_name": "first", "last_name": "llllast"})
	fmt.Println(u2)

	u.FirstName = "aaaaaaaaaaaaa"
	db.Save(&u)
	u = User{FirstName: "aaaaaaaaaaaaa"}
	db.Where(&u).First(&u)
	fmt.Println(u) //aaaaaaaaaaaaa

	// Delete
	fmt.Println("Delete")
	db.Delete(&u2)
	fmt.Println(u2)

	db.Where(&User{FirstName: "aaaaaaaaaaaaa"}).Delete(&User{})

	println("connection to database established")
}

func main() {
	InitialMigration()
}

// User model
type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}
