package main

import (
	"database/sql"

	"github.com/almirpask/go_api/internal/database"
	"github.com/almirpask/go_api/internal/entity"
	"github.com/almirpask/go_api/internal/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(catalog_db:3306)/catalog")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)
	products := []*entity.Product{
		{Name: "Product 1", Description: "Description 1", Price: 90, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/1.png"},
		{Name: "Product 2", Description: "Description 2", Price: 200, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/2.png"},
		{Name: "Product 3", Description: "Description 3", Price: 300, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/3.png"},
		{Name: "Product 4", Description: "Description 4", Price: 400, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/4.png"},
		{Name: "Product 5", Description: "Description 5", Price: 500, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/5.png"},
		{Name: "Product 6", Description: "Description 6", Price: 600, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/6.png"},
		{Name: "Product 7", Description: "Description 7", Price: 700, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/7.png"},
		{Name: "Product 8", Description: "Description 8", Price: 27, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/8.png"},
		{Name: "Product 9", Description: "Description 9", Price: 900, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/9.png"},
		{Name: "Product 10", Description: "Description 10", Price: 1000, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/10.png"},
		{Name: "Product 11", Description: "Description 11", Price: 1100, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/11.png"},
		{Name: "Product 12", Description: "Description 12", Price: 1200, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/12.png"},
		{Name: "Product 13", Description: "Description 13", Price: 1300, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/13.png"},
		{Name: "Product 14", Description: "Description 14", Price: 1400, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/14.png"},
		{Name: "Product 15", Description: "Description 15", Price: 1500, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/15.png"},
		{Name: "Product 16", Description: "Description 16", Price: 1600, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/16.png"},
		{Name: "Product 17", Description: "Description 17", Price: 1700, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/17.png"},
		{Name: "Product 18", Description: "Description 18", Price: 1800, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/18.png"},
		{Name: "Product 19", Description: "Description 19", Price: 1900, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/19.png"},
		{Name: "Product 20", Description: "Description 20", Price: 2000, CategoryID: "6b4c28f4-6831-495a-9444-19c93452faa3", ImageURL: "http://image_server:9000/products/20.png"},
	}

	for _, product := range products {
		_, err := productService.CreateProduct(product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
		if err != nil {
			return
		}
	}

}
