//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

// Customer — покупатель.
type Customer struct {
	Name string
}

func (c *Customer) SetName(name string) {
	c.Name = name
}

func (c *Customer) String() string {
	return fmt.Sprintf("Customer: %s", c.Name)
}

// Seller — продавец.
type Seller struct {
	Name string
}

func (c *Seller) SetName(name string) {
	c.Name = name
}

func (c *Seller) String() string {
	return fmt.Sprintf("Seller: %s", c.Name)
}

// Provider — интерфейс фабрики.
type Provider interface {
	NewCustomer() *Customer
	NewSeller() *Seller
}

type GoogleAuth struct {
}

func (g *GoogleAuth) NewCustomer() *Customer {
	var customer Customer
	// получаем имя из Google-аккаунта
	name := "Google Customer"
	customer.SetName(name)
	return &customer
}

func (g *GoogleAuth) NewSeller() *Seller {
	var seller Seller
	// получаем имя из Google-аккаунта
	name := "Google Seller"
	seller.SetName(name)
	return &seller
}

type YandexAuth struct {
}

func (g *YandexAuth) NewCustomer() *Customer {
	var customer Customer
	// получаем имя из Yandex-аккаунта
	name := "Yandex Customer"
	customer.SetName(name)
	return &customer
}

func (g *YandexAuth) NewSeller() *Seller {
	var seller Seller
	// получаем имя из Yandex-аккаунта
	name := "Yandex Seller"
	seller.SetName(name)
	return &seller
}

// AuthFactory — абстрактная фабрика аутентификации.
func AuthFactory(provider string) Provider {
	switch provider {
	case "google":
		return &GoogleAuth{}
	case "yandex":
		return &YandexAuth{}
	default:
		panic(fmt.Sprintf("unknown provider %s", provider))
	}
}
