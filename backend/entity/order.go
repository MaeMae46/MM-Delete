package entity


import (

   "time"

   "gorm.io/gorm"

)

type Order struct {

   gorm.Model

   ID        uint     `json:"id"`

   OrderDate     time.Time    `json:"order_date"`

   Status string    `json:"status"`

   TotalPrice  float32    `json:"total_price"`

   UserID  uint      `json:"user_id"`

   User    *Users  `gorm:"foreignKey: id" json:"user"`

   PaymentID  uint      `json:"payment_id"`

   //Payment    *Payment  `gorm:"foreignKey: id" json:"payment_id"`

   CodeCollectID  uint      `json:"code_collect_id"`

   //CodeCollect    *CodeCollect  `gorm:"foreignKey: id" json:"code_collect_id"`

   ShippingID  uint      `json:"shipping_id"`

   OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
   //Shipping    *Shipping  `gorm:"foreignKey: id" json:"shipping_id"`
}