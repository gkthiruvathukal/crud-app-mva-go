
package main

import (
    "flag"
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"index"`
    Email string `gorm:"uniqueIndex"`
    Note  string
}

func main() {
    id := flag.Uint("id", 0, "User's ID")
    flag.Parse()

    if *id == 0 {
        fmt.Println("ID is required")
        return
    }

    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.Delete(&User{}, *id)
    fmt.Printf("Deleted user with ID: %d\n", *id)
}
