package config


import (

   "fmt"

   "time"

   "github.com/MaeMae46/MM-Se-67/backend/entity"

   "gorm.io/driver/sqlite"

   "gorm.io/gorm"

)


var db *gorm.DB


func DB() *gorm.DB {

   return db

}


func ConnectionDB() {

   database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})

   if err != nil {

       panic("failed to connect database")

   }

   fmt.Println("connected database")

   db = database

}


func SetupDatabase() {

    db.AutoMigrate(
        &entity.Catagory{},
        &entity.History{},
        &entity.HistoryDetail{},
        &entity.Order{},
        &entity.OrderItem{},
        &entity.Product{},
        &entity.Stock{},
        &entity.Users{},
        &entity.Point{},
        &entity.PointPolicy{},
        &entity.PointTran{},
    )

    hashedPassword, _ := HashPassword("123456")
    BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")

    User := &entity.Users{
        FirstName:   "Kunlasatri",
        LastName:    "Kramoncham",
        ID:          1,
        Email:       "May@gmail.com",
        Password:    hashedPassword,
        BirthDay:    BirthDay,
        PhoneNumber: "0615871759",
        Role:        "admin",
        PointID:     1,
    }

    db.FirstOrCreate(User, &entity.Users{
        Email: "May@gmail.com",
    })

    Point := &entity.Point{
        ID:          1,
        TotalPoint: 100,
    }

    db.FirstOrCreate(&Point, entity.Point{ID: Point.ID})

    PointPolicy := &entity.PointPolicy{
        ID:          1,
        EarnRate:    4,  // 4 แต้ม ต่อการใช้จ่าย 1 บาท
        RedeemRate:  100, // ใช้ 100 แต้ม เพื่อลดราคา 1 บาท
        Description: "สะสม 4 แต้ม ต่อการใช้จ่าย 1 บาท และใช้ 100 แต้ม เพื่อลดราคา 1 บาท",
    }

    db.FirstOrCreate(&PointPolicy, entity.PointPolicy{ID: PointPolicy.ID})
    

    // ตัวอย่างข้อมูล Stock
    stocks := []entity.Stock{
        {
            ID:        1,
            Price:     850,
            Quantity:  100,
            Color:     "ดำ",
            ShapeSize: "เหลี่ยม",
            Image:     "/assets/sofa.jpg",
            ProductID: 1,
        },
        {
            ID:        2,
            Price:     1200,
            Quantity:  50,
            Color:     "น้ำตาล",
            ShapeSize: "กลม",
            Image:     "/assets/table.jpg",
            ProductID: 2,
        },
    }

    for _, stock := range stocks {
        db.FirstOrCreate(&stock, entity.Stock{ID: stock.ID})
    }

    // ตัวอย่างข้อมูล Product
    products := []entity.Product{
        {
            ID:          1,
            Name:        "โซฟา",
            Description: "โซฟา 3 ที่นั่ง รุ่น Junie หุ้มด้วยผ้านำเข้าจากต่างประเทศ",
            Image:       "/assets/sofa.jpg",
            UserID:      1,
            CatagoryID:  1,
        },
        {
            ID:          2,
            Name:        "โต๊ะ",
            Description: "โต๊ะไม้ยางพาราสำหรับใช้งานทั่วไป",
            Image:       "/assets/table.jpg",
            UserID:      1,
            CatagoryID:  2,
        },
    }

    for _, product := range products {
        db.FirstOrCreate(&product, entity.Product{ID: product.ID})
    }

    // ตัวอย่างข้อมูล Order และ OrderItem
    orders := []entity.Order{
        {
            ID:          1,
            OrderDate:   time.Now(),
            Status:      "pending",
            TotalPrice:  1700,
            UserID:      1,
        },
    }

    for _, order := range orders {
        db.FirstOrCreate(&order, entity.Order{ID: order.ID})
    }

    // ตัวอย่างข้อมูล OrderItem
    orderItems := []entity.OrderItem{
        {
            ID:       1,
            Quantity: 2,
            Price:    1700,
            OrderID:  1,
            StockID:  1,
        },
    }

    for _, item := range orderItems {
        db.FirstOrCreate(&item, entity.OrderItem{ID: item.ID})
    }

    // ตัวอย่างข้อมูล History และ HistoryDetail
    histories := []entity.History{
        {
            ID:             1,
            OrderDate:      time.Now(),
            PointsEarned:   10,
            PointsRedeemed: 0,
            TotalAmount:    1700,
            UserID:         1,
            OrderID:        1,
        },
    }

    for _, history := range histories {
        db.FirstOrCreate(&history, entity.History{ID: history.ID})
    }

    historyDetails := []entity.HistoryDetail{
        {
            ID:          1,
            ProductName: "โซฟา",
            Quantity:    2,
            PricePerUnit: 850,
            SubTotal:    1700,
            StockID:     1,
            HistoryID:   1,
        },
    }

    for _, detail := range historyDetails {
        db.FirstOrCreate(&detail, entity.HistoryDetail{ID: detail.ID})
    }
}
