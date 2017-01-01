package main

import   "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"
import   "github.com/julienschmidt/httprouter"
import   "fmt"
import   "net/http"
import   "log"

type Caderno struct{
	gorm.Model
	Code string
	Price uint
}

func TestProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	db, err := gorm.Open("postgres", "host=localhost user=deyvison dbname=learngo sslmode=disable password=88072762")
	if err != nil {
		panic("failed to connect database")
	}
  defer db.Close()

// Migrate the schema
db.AutoMigrate(&Product{})

// Create
db.Create(&Product{Code: "l1212", Price: 1000})

// Read
//var product Product
//db.First(&product, 1) //Find product whitd id = 1
// db.First(&product, "code = ?","L1212") // Find product with code L1212

// Update - update product's price to 2000
// db.Model(&product).Update("Price", 2000)

// Delete - delete product
// db.Delete(&product)

fmt.Fprint(w, "Successfull")
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Feliz Ano Novo!")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/db", TestProduct)

	log.Fatal(http.ListenAndServe(":8080", router))
}
