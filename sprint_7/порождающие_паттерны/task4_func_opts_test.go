//go:build ignore
// +build ignore

package main

import "testing"

func TestNewCar(t *testing.T) {
	sportsCar := NewCar("sports car", WithSeats(2), WithDoors(2))
	minivan := NewCar("minivan", WithSeats(7), WithDoors(5), WithABS())

	if sportsCar.String() != "sports car [seats:2 / doors: 2 / abs: false]" {
		t.Errorf("wrong %s", sportsCar.String())
	}
	if minivan.String() != "minivan [seats:7 / doors: 5 / abs: true]" {
		t.Errorf("wrong %s", minivan.String())
	}
}
