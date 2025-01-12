package entity


import (

   "time"

   "gorm.io/gorm"

)

type Users struct {

   gorm.Model

   ID        uint     `json:"id"`

   Email     string    `json:"email"`

   FirstName string    `json:"first_name"`

   LastName  string    `json:"last_name"`

   Password  string    `json:"-"`

   PhoneNumber string  `json:"phone_number"`

   Role      string     `json:"role"`

   BirthDay  time.Time `json:"birthday"`

   PointID  uint      `json:"point_id"`

   Point    *Point  `gorm:"foreignKey: id" json:"point_id"`
}