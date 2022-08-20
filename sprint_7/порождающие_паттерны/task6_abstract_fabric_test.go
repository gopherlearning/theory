//go:build ignore
// +build ignore

package main

import (
	"testing"
)

func TestAuth(t *testing.T) {
	googleAuth := AuthFactory("google")
	yandexAuth := AuthFactory("yandex")

	if googleAuth.NewCustomer().String() != "Customer: Google Customer" {
		t.Errorf("wrong google customer")
	}
	if googleAuth.NewSeller().String() != "Seller: Google Seller" {
		t.Errorf("wrong google seller")
	}
	if yandexAuth.NewCustomer().String() != "Customer: Yandex Customer" {
		t.Errorf("wrong yandex customer")
	}
	if yandexAuth.NewSeller().String() != "Seller: Yandex Seller" {
		t.Errorf("wrong yandex seller")
	}
}
