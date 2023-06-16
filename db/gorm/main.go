package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)
  
func main() {
    db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
    defer db.Close()
}
