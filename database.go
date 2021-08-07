package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&Product{})

  db.Create(&Product{Code: "D42", Price: 100})

  
  var product Product
  db.First(&product, 1) 
  db.First(&product, "code = ?", "D42") 

  
  db.Model(&product).Update("Price", 200)
  
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})


  db.Delete(&product, 1)
  type User struct {
	ID           1
	Name         shop_db
	Email        A@gmail.com
	Age          19
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
  }
}