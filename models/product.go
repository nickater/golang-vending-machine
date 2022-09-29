package models

type Product struct {
	Name  string
	Price int64
}

type RawProduct struct {
	Base     Product
	Quantity int
}
