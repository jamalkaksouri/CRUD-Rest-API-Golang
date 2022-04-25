package main

import (
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/models"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func main() {

	database.Connect("host=localhost user=postgres password=root dbname=crud_rest port=5432 sslmode=disable")

	product := []models.Product{}
	database.Instance.Find(&product)
	if len(product) == 0 {
		log.Println("Adding products to the database...")
		for i := 0; i < 100; i++ {
			store := models.Product{
				Name:        faker.Word(),
				Price:       float64(rand.Intn(9999) + 10),
				Description: faker.Paragraph(),
			}
			database.Instance.Create(&store)
		}
		log.Println("100 Products were added to the database!")
	} else {
		log.Println("Products have already been added in the database!")
	}
}
