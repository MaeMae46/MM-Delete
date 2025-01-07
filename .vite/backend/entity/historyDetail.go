package entity


import (

   "gorm.io/gorm"

)

type HistoryDetail struct {

   gorm.Model

   ID        uint     `json:"id"`

   ProductName     string    `json:"product_name"`

   Quantity  uint    `json:"quantity"`

   PricePerUnit  float32    `json:"price_per_unit"`

   SubTotal  float32    `json:"sub_total"`

   StockID  uint      `json:"stock_id"`

   Stock    *Stock  `gorm:"foreignKey: id" json:"stock_id"`

   HistoryID  uint      `json:"history_id"`

   History    *History  `gorm:"foreignKey: id" json:"history_id"`
}