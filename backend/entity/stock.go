package entity


import (

   "gorm.io/gorm"

)

type Stock struct {
    gorm.Model
    ID               uint     `json:"id" valid:"required,range(1|10000)"`
    Price            float32  `json:"price" valid:"required,range(0.01|10000)"`
    Quantity         uint     `json:"quantity_in_stock" valid:"required,range(1|10000)"`
    Color            string   `json:"color" valid:"required"`
    ShapeSize        string   `json:"shape_size" valid:"required"`
    Image            string   `json:"image" valid:"required,url"`
    ProductID        uint     `json:"product_id" valid:"required"`
    Product          *Product `gorm:"foreignKey:ProductID" json:"product"`
}
