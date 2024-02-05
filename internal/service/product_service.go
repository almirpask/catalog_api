package service

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/almirpask/go_api/internal/database"
	"github.com/almirpask/go_api/internal/entity"
	"github.com/almirpask/go_api/rabbitmq"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDb database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: productDb,
	}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) CreateProduct(name string, description string, price float64, categoryID string, imageURL string) (*entity.Product, error) {
	ctx := context.Background()

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	product := entity.NewProduct(name, description, price, categoryID, imageURL)

	createdProduct, err := ps.ProductDB.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	responseJson, err := json.Marshal(createdProduct)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	err = rabbitmq.Publish(ctx, ch, "ProductCreated", string(responseJson), "amq.direct")
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	slog.Info("Product created", "id", createdProduct.ID)

	return product, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(categoryID)

	if err != nil {
		return nil, err
	}

	return products, nil
}
