package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  int
}

type Game struct {
	GameName string
	Price    int
	Genre    string
}

type Product interface {
	ProductName() string
	ProductCost() int
}

func (b Book) ProductName() string {
	return b.Title
}

func (b Book) ProductCost() int {
	return b.Price
}

func (g Game) ProductName() string {
	return g.GameName
}

func (g Game) ProductCost() int {
	return g.Price
}

func shopCal(p Product) {
	fmt.Printf("Product Name: %s, Cost: %dRS\n", p.ProductName(), p.ProductCost())
}

func main() {

	book := Book{
		Title:  "BooK1",
		Author: "F.K Diksha",
		Price:  1500,
	}

	game := Game{
		GameName: "mario",
		Genre:    "Adventure",
		Price:    590,
	}

	shopCal(book)
	shopCal(game)
}
