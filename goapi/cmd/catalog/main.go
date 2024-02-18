package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/angelino-valeta/code-commerce/internal/database"
	"github.com/angelino-valeta/code-commerce/internal/service"
	"github.com/angelino-valeta/code-commerce/internal/webserver"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/codecomerce")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProudctHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)

	c.Get("/product", webProudctHandler.GetProducts)
	c.Get("/product/category/{categoryID}", webProudctHandler.GetProductByCategoryID)
	c.Get("/product/{id}", webProudctHandler.GetProduct)
	c.Post("/product", webProudctHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
