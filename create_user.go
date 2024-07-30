
package main

import (
    "flag"
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "strings"
)

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"index"`
    Email string `gorm:"uniqueIndex"`
    Note  string
    Tags  []Tag `gorm:"many2many:user_tags;"`
}

type Tag struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"uniqueIndex"`
}

func main() {
    name := flag.String("name", "", "User's name")
    email := flag.String("email", "", "User's email")
    note := flag.String("note", "", "User's note")
    tags := flag.String("tags", "", "Comma-separated list of tags")

    flag.Parse()

    if *name == "" || *email == "" {
        fmt.Println("Name and email are required")
        return
    }

    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{}, &Tag{})

    var tagModels []Tag
    for _, tagName := range strings.Split(*tags, ",") {
        tagName = strings.TrimSpace(tagName)
        var tag Tag
        db.FirstOrCreate(&tag, Tag{Name: tagName})
        tagModels = append(tagModels, tag)
    }

    user := User{Name: *name, Email: *email, Note: *note, Tags: tagModels}
    db.Create(&user)

    fmt.Printf("Created user: %+v\n", user)
}
