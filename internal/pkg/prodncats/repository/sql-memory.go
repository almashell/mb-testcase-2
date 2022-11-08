package repository

//go:generate mockgen -destination=../mocks/mock_repository.go -package=mocks mb_api/internal/pkg/prodncats Repository

import "C"
import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"log"
	"mb_api/internal/pkg/models"
	"mb_api/internal/pkg/prodncats"
)

type sqlProdncatsRepository struct {
	db *pgx.ConnPool
}

func NewSqlProdncatsRepository(db *pgx.ConnPool) prodncats.Repository {
	return &sqlProdncatsRepository{db: db}
}

func (er *sqlProdncatsRepository) GetProducts() ([]models.Products, error) {
	sql := "SELECT P.name, C.name FROM products P LEFT JOIN productcategories PC ON p.id = PC.product_id LEFT JOIN categories C ON PC.category_id = C.id"
	rows, err := er.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var products []models.Products
	for rows.Next() {
		prodInfo := models.Products{}
		catn := pgtype.Varchar{}
		err = rows.Scan(
			&prodInfo.ProductName,
			&catn)
		prodInfo.ProductCategory = catn.String
		if err != nil {
			return nil, err
		}
		products = append(products, prodInfo)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (er *sqlProdncatsRepository) GetCategories() ([]models.Categories, error) {
	sql := "SELECT C.name, P.name FROM products P RIGHT JOIN productcategories PC ON p.id = PC.product_id RIGHT JOIN categories C ON PC.category_id = C.id"
	rows, err := er.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var categories []models.Categories
	for rows.Next() {
		log.Println(rows)
		catInfo := models.Categories{}
		prodn := pgtype.Varchar{}
		err = rows.Scan(
			&catInfo.CategoryName,
			&prodn)
		catInfo.CategoryProduct = prodn.String
		if err != nil {
			return nil, err
		}
		categories = append(categories, catInfo)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (er *sqlProdncatsRepository) GetPairs() ([]models.Pairs, error) {
	sql := "SELECT P.name, C.name FROM products P FULL JOIN productcategories PC ON p.id = PC.product_id FULL JOIN categories C ON PC.category_id = C.id"
	rows, err := er.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var pairs []models.Pairs
	for rows.Next() {
		pairInfo := models.Pairs{}
		catn := pgtype.Varchar{}
		prodn := pgtype.Varchar{}
		err = rows.Scan(
			&prodn,
			&catn)
		pairInfo.ProductName = prodn.String
		pairInfo.CategoryName = catn.String
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pairInfo)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return pairs, nil
}