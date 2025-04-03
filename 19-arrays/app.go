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
	var productsArray []Product = []Product{productOne, productTwo}

	productsArray = append(productsArray, productThree)

	fmt.Println(productsArray)

	// products := NewProducts(productsArray)
	// fmt.Println(products.data)
	// fmt.Println(products.data[0])

	// // Assign value in an empty space of the array
	// productsArray[2] = productThree
	// products = NewProducts(productsArray)
	// fmt.Println(products.data)

	// // Length
	// fmt.Println(len(products.data))
	// fmt.Println(cap(products.data))
}
