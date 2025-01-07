package controller

import (
	"net/http"
    "log"
	"github.com/gin-gonic/gin"
	"github.com/MaeMae46/MM-Se-67/backend/config"
	"github.com/MaeMae46/MM-Se-67/backend/entity"
)

type CreateOrderRequest struct {
	UserID   uint                     `json:"user_id"`
	Items    []CreateOrderItemRequest `json:"items"`
}

type CreateOrderItemRequest struct {
	StockID  uint `json:"stock_id"`
	Quantity uint `json:"quantity"`
}

func CreateOrder(c *gin.Context) {
    var order entity.Order

    // รับข้อมูลคำสั่งซื้อจาก JSON
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // เริ่มต้น Transaction
    tx := config.DB().Begin()

    // บันทึกคำสั่งซื้อ
    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    // Commit การเปลี่ยนแปลง
    tx.Commit()

    c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": order})
}

func GetOrderHistory(c *gin.Context) {
    userId := c.Param("userId")
    var orders []entity.Order

    // ดึงข้อมูล Order จากฐานข้อมูล
    if err := config.DB().Where("user_id = ?", userId).Preload("OrderItems.Stock").Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order history"})
        return
    }

    if len(orders) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No orders found for this user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetHistoryDetail(c *gin.Context) {
    userId := c.Param("userId")
    var historyDetails []entity.History

    if err := config.DB().Where("user_id = ?", userId).Preload("Order").Find(&historyDetails).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch history details"})
        log.Println("Error fetching history details:", err)
        return
    }

    if len(historyDetails) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No history details found for this user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"historyDetails": historyDetails})
}

func UpdateOrderStatus(c *gin.Context) {
    orderID := c.Param("orderId")
    var order entity.Order

    // รับสถานะใหม่จาก request body
    var requestBody struct {
        Status string `json:"status"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // ดึงข้อมูลคำสั่งซื้อจากฐานข้อมูลพร้อมกับ OrderItems
    if err := config.DB().Preload("OrderItems").First(&order, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    // ตรวจสอบสถานะคำสั่งซื้อก่อน
    if order.Status == "Completed" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Order is already completed"})
        return
    }

    // อัพเดตสถานะคำสั่งซื้อ
    order.Status = requestBody.Status

    // หากสถานะเป็น "Completed" ให้หักจำนวนสินค้า
    if order.Status == "Completed" {
        tx := config.DB().Begin() // เริ่มต้น transaction

        // หักจำนวนสินค้าในแต่ละรายการ
        for _, item := range order.OrderItems {
            var stock entity.Stock
            if err := tx.First(&stock, item.StockID).Error; err != nil {
                tx.Rollback()
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stock"})
                return
            }

            // ตรวจสอบจำนวนสินค้าคงเหลือในสต็อก
            if stock.Quantity < item.Quantity {
                tx.Rollback()
                c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock for product"})
                return
            }

            // หักจำนวนสินค้า
            stock.Quantity -= item.Quantity

            // บันทึกการเปลี่ยนแปลงในสต็อก
            if err := tx.Save(&stock).Error; err != nil {
                tx.Rollback()
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
                return
            }
        }

        // บันทึกการเปลี่ยนแปลงสถานะคำสั่งซื้อ
        if err := tx.Save(&order).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
            return
        }

        tx.Commit() // ยืนยันการเปลี่ยนแปลงทั้งหมด
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully", "order": order})
}


