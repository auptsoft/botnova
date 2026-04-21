package entities

type Project struct {
	Id          int `gorm:"primaryKey"`
	UserID      int
	Name        string
	Description string
	Options     string //Extra details
}
