
package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "strings"

    "gopkg.in/yaml.v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
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

type UserData struct {
    Name  string   `yaml:"name"`
    Email string   `yaml:"email"`
    Note  string   `yaml:"note"`
    Tags  []string `yaml:"tags"`
}

func main() {
    name := flag.String("name", "", "User's name")
    email := flag.String("email", "", "User's email")
    note := flag.String("note", "", "User's note")
    tags := flag.String("tags", "", "Comma-separated list of tags")
    fromYaml := flag.String("from-yaml", "", "Path to YAML file with user data")
    flag.Parse()

    var userData UserData

    if *fromYaml != "" {
        data, err := ioutil.ReadFile(*fromYaml)
        if err != nil {
            log.Fatalf("failed to read YAML file: %v", err)
        }
        err = yaml.Unmarshal(data, &userData)
        if err != nil {
            log.Fatalf("failed to unmarshal YAML: %v", err)
        }
    } else {
        userData = UserData{
            Name:  *name,
            Email: *email,
            Note:  *note,
            Tags:  strings.Split(*tags, ","),
        }
    }

    if userData.Name == "" || userData.Email == "" {
        log.Fatalln("Name and email are required")
    }

    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{}, &Tag{})

    var tagModels []Tag
    for _, tagName := range userData.Tags {
        tagName = strings.TrimSpace(tagName)
        var tag Tag
        db.FirstOrCreate(&tag, Tag{Name: tagName})
        tagModels = append(tagModels, tag)
    }

    user := User{Name: userData.Name, Email: userData.Email, Note: userData.Note, Tags: tagModels}
    db.Create(&user)

    fmt.Printf("Created user: %+v\n", user)
}
