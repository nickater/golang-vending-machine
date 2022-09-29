package utils

import (
	"log"
	"strconv"

	"github.com/nickater/golang-vending-machine/models"
)

func DecipherProductCSV(csv [][]string) []models.RawProduct {
	rawProducts := []models.RawProduct{}
	for _, line := range csv {
		product := models.Product{}
		product.Name = line[0]
		intVar, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalln("Number didn\t convert")
		}
		product.Price = int64(intVar)
		quantity, err := strconv.Atoi(line[2])

		rawProducts = append(rawProducts, models.RawProduct{
			Base:     product,
			Quantity: quantity,
		})
	}

	return rawProducts
}
