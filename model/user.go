package model

type UserType string

const (
    USER UserType = "user"
)

type User struct {
    Id   string   `gorm:"primary_key" json:"id"`
    Name string   `json:"name"`
    Type UserType `json:"type"`
}

func (User) TableName() string {
    return "qianshou.user"
}
