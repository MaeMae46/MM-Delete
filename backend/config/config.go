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
   )



   hashedPassword, _ := HashPassword("123456")

   BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")

   User := &entity.Users{
	
       FirstName: "Kunlasatri",

       LastName:  "Kramoncham",

	   ID: 1 ,

       Email:     "May@gmail.com",

       Password:  hashedPassword,

       BirthDay:  BirthDay,
	
	   PhoneNumber: "0615871759",

	   Role: "admin",
	
	   PointID: 1,

   }

   db.FirstOrCreate(User, &entity.Users{

       Email: "May@gmail.com",

   })

   stocks := []entity.Stock{
      {
          ID:         1,
          Price:      850,
          Quantity:   100,
          Color:      "ดำ",
          ShapeSize:  "เหลี่ยม",
          Image:      "/assets/sofa.jpg",
          ProductID:  1,
      },
      {
          ID:         2,
          Price:      1200,
          Quantity:   50,
          Color:      "น้ำตาล",
          ShapeSize:  "กลม",
          Image:      "/assets/table.jpg",
          ProductID:  2,
      },
      {
          ID:         3,
          Price:      750,
          Quantity:   30,
          Color:      "ขาว",
          ShapeSize:  "สามเหลี่ยม",
          Image:      "/assets/chair.jpg",
          ProductID:  3,
      },
  }

  // วนลูปสร้างข้อมูล `Stock`
  for _, stock := range stocks {
      db.FirstOrCreate(&stock, entity.Stock{ID: stock.ID})
  }

  // ตัวอย่างการเพิ่มข้อมูล `Product`
  products := []entity.Product{
      {
          ID:          1,
          Name:        "โซฟา",
          Description: "โซฟา 3 ที่นั่ง รุ่น Junie หุ้มด้วยผ้านำเข้าจากต่างประเทศ นุ่มสบาย ไม่หดตัว อายุการใช้งานทนทาน แข็งแรง ทนทาน ดีไซน์ทันสมัยและสีสันแตกต่าง",
          Image:       "/assets/sofa.jpg",
          UserID:      1,
          CatagoryID:  1,
      },
      {
          ID:          2,
          Name:        "โต๊ะ",
          Description: "โต๊ะไม้ยางพาราสำหรับใช้งานทั่วไป แข็งแรงและใช้งานได้หลากหลาย",
          Image:       "/assets/table.jpg",
          UserID:      1,
          CatagoryID:  2,
      },
      {
          ID:          3,
          Name:        "เก้าอี้",
          Description: "เก้าอี้หวายทอมืออย่างประณีต เหมาะสำหรับการตกแต่งบ้าน",
          Image:       "/assets/chair.jpg",
          UserID:      1,
          CatagoryID:  3,
      },
  }

  // วนลูปสร้างข้อมูล `Product`
  for _, product := range products {
      db.FirstOrCreate(&product, entity.Product{ID: product.ID})
  }
}