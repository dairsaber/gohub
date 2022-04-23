package users

import "gohub/app/models"

// User 用户模型
type User struct {
	models.BaseModel

	Name  string `json:"name,omitempty"`
	Email string `json:"_"`
}
