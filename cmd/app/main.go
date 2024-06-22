package main

import (
	"api/internal/infra/repository"
	"api/internal/infra/web"
	"api/internal/usecase"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi"
)

func main() {
	dbServer := ""
	dbPort := 1433
	dbUser := "sa"
	dbPassword := ""
	dbDatabase := ""

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		dbServer, dbUser, dbPassword, dbPort, dbDatabase)

	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepositoryMssql(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)
	listProductsUsecase := usecase.NewListProductUseCase(repository)
	productHandlers := web.NewProductHandlers(createProductUsecase, listProductsUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	http.ListenAndServe(":8080", r)

	/*
		msgChan := make(chan *kafka.Message)
		go akafka.Consume([]string{"products"}, "host", msgChan)

		for msg := range msgChan {
			dto := usecase.CreateProductInputDTO{}
			err := json.Unmarshal(msg.Value, &dto)

			if err != nil {

			}

			_, err = createProductUsecase.Execute(dto)
		}*/

}
