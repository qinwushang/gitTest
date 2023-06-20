package models

type Book struct {
	Id          int    //`gorm:"AUTO_INCREMENT"`
	Type        string `form:"type"`
	Name        string `form:"name"`
	Description string `form:"description"`
	Url         string `form:"url"`
}

type Page struct {
	PageNum  int `form:"page"`
	PageSize int `form:"size"`
}
