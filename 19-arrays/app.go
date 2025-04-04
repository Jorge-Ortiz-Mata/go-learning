package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Price float64
}

type Products struct {
	data []Product
}

func New(name string, price float64) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}

func NewProducts(products []Product) Products {
	return Products{
		data: products,
	}
}

func main() {
	productOne := New("Mac Pro Max", 99.99)
	productTwo := New("Iphone 15", 34.99)
	productThree := New("Kia Forte", 999.99)
	productFour := New("TV 70", 349.99)
	var productsArray []Product = []Product{productOne, productTwo}

	// Can add multiple objetcs
	productsArray = append(productsArray, productThree, productFour)

	fmt.Println(productsArray) // [{0 Mac Pro Max 99.99} {0 Iphone 15 34.99} {0 Kia Forte 999.99} {0 TV 70 349.99}]

	// Spread operator
	productFive := New("Laptop", 19.99)
	productSix := New("Car", 2.99)
	var otherProducts []Product = []Product{productFive, productSix}
	productsArray = append(productsArray, otherProducts...)
	fmt.Println(productsArray) // [{0 Mac Pro Max 99.99} {0 Iphone 15 34.99} {0 Kia Forte 999.99} {0 TV 70 349.99} {0 Laptop 19.99} {0 Car 2.99}]
}
