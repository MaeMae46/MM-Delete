package entity


import (

   "gorm.io/gorm"

)

type OrderItem struct {
   gorm.Model

   ID       uint    `json:"id"`
   Quantity uint    `json:"quantity"`
   Price    float32 `json:"price"`

   OrderID uint   `json:"order_id"`
   Order   *Order `gorm:"foreignKey:OrderID" json:"order"`

   StockID uint   `json:"stock_id"`
   Stock   *Stock `gorm:"foreignKey:StockID" json:"stock"`
}
