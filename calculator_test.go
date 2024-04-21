package main

import (
	"testing"
)

func TestHitungKelilingPersegi(t *testing.T) {
	sisi := 6.0
	Hasil := 24.0
	totalHitung := HitungKelilingPersegi(sisi)

	if totalHitung != Hasil {
		t.Errorf("Maaf hasil perhitungan dari keliling persegi salah!. Hasil yang diharapkan: %f, Hasil: %f", Hasil, totalHitung)
	}
}

func TestHitungKelilingSegiLima(t *testing.T) {
	sisi := 3.0
	Hasil := 15.0
	totalHitung := HitungKelilingSegiLima(sisi)

	if totalHitung != Hasil {
		t.Errorf("Maaf hasil perhitungan dari keliling segi lima salah!. Hasil yang diharpkan: %f, Hasil: %f", Hasil, totalHitung)
	}
}
