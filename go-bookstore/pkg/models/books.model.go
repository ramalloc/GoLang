package models


import "gorm.io/gorm"
import "github.com/ramalloc/go-bookstore/pkg/config"


var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"book"`
	Author string `gorm:"" json:"author"`
	Publiction string `gorm:"" json:"publiction"`
}


func init (){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}


func (b *Book) CreateBook() *Book{
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var fetchedBook *Book
	db := db.Where("ID=?", Id).Find(&fetchedBook)
	return fetchedBook, db
}

func DeleteBookById(Id int64) *Book{
	var toDeleteBook *Book
	db.Where("ID=?", Id).Delete(&toDeleteBook)
	return toDeleteBook
}