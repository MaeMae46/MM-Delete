package main


import (

   "net/http"


   "github.com/gin-gonic/gin"


   "github.com/MaeMae46/MM-Se-67/backend/config"

   "github.com/MaeMae46/MM-Se-67/backend/controller"

)


const PORT = "8000"


func main() {


   // open connection database

   config.ConnectionDB()


   // Generate databases

   config.SetupDatabase()


   r := gin.Default()


   r.Use(CORSMiddleware())

   router := r.Group("/")

   {
        router.PATCH("/orders/:orderID/status", controller.UpdateOrderStatus)
        r.GET("/history/detail/:userId", controller.GetHistoryDetail)
        r.GET("/orders/history/:userId", controller.GetOrderHistory)
        r.POST("/orders", controller.CreateOrder)

        r.GET("/stocks", controller.GetStocksByProductID)
        r.POST("/stock", controller.Create)
        router.PUT("/stock/:id", controller.Update)
        router.GET("/stock", controller.GetAll)
        router.GET("/stock/:id", controller.Get)
        router.DELETE("/stock/:id", controller.Delete)

        router.GET("/products", controller.GetAllProduct)
		router.GET("/products/:id", controller.GetProductByID)
		router.PUT("/products/:id", controller.UpdateProduct)
		router.DELETE("/products/:id", controller.DeleteProduct)


   }

   r.GET("/", func(c *gin.Context) {

       c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)

   })


   // Run the server


   r.Run("localhost:" + PORT)


}


func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // ตั้งค่าให้อนุญาตการเชื่อมต่อจาก frontend
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, CREATE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
