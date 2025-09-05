package entity

type Video struct {
	Title       string `json:"title" binding:"min=2,max=20,required,is-cool"`
	Description string `json:"description" binding:"required,max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      User   `json:"author" binding:"required" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ID     uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint64 `json:"userId" binding:"required"`
}
