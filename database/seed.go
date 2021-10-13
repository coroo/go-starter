package main

import (
	seeds "github.com/coroo/go-starter/database/seeders"
)

func main() {
	seeds.SeedPaymentMethods()
	seeds.SeedPaymentMethodRates()
	seeds.SeedPaymentMethodLinks()
}