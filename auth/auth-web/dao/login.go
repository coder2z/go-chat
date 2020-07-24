package dao

type LoginDao struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
