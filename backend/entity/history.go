package entity


import (

   "time"

   "gorm.io/gorm"

)

type History struct {

   gorm.Model

   ID        uint     `json:"id"`

   OrderDate     time.Time    `json:"order_date"`

   PointsEarned float32    `json:"points_earned"`

   PointsRedeemed  float32    `json:"points_redeemed"`

   TotalAmount  float32    `json:"total_amount"`

   UserID  uint      `json:"user_id"`

   User    *Users  `gorm:"foreignKey: id" json:"user_id"`

   OrderID  uint      `json:"order_id"`

   Order    *Order  `gorm:"foreignKey: id" json:"order_id"`
}