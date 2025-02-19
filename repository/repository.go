package repository

import (
	"context"
	"crud_fire/model"

	"cloud.google.com/go/firestore"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	InsertProduct(ctx context.Context, product model.Product) error
	EditProduct(ctx context.Context, product model.Product) error
	DeleteProduct(ctx context.Context, product model.Product) error
}

type productRepository struct {
	client *firestore.Client
}

func NewProductRepository(client *firestore.Client) ProductRepository {
	return &productRepository{client: client}
}

func (r *productRepository) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	iter := r.client.Collection("coffe_aplication").Doc("data").Collection("product").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var product model.Product
		if err := doc.DataTo(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) InsertProduct(ctx context.Context, product model.Product) error {
	_, _, err := r.client.Collection("coffe_aplication").Doc("data").Collection("product").Add(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) EditProduct(ctx context.Context, product model.Product) error {
	_, err := r.client.Collection("coffe_aplication").Doc("data").Collection("product").Doc(product.Nama).Set(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, product model.Product) error {
	_, err := r.client.Collection("coffe_aplication").Doc("data").Collection("product").Doc(product.Nama).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
