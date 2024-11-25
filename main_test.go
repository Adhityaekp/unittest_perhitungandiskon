package main

import "testing"

type TestCase struct {
	hargaAwal    float64
	persenDiskon float64
	expected     float64
}

func TestHitungDiskon(t *testing.T) {
	var listTest = []TestCase{
		{hargaAwal: 1000, persenDiskon: 20, expected: 800},
		{hargaAwal: 500, persenDiskon: 10, expected: 450},
		{hargaAwal: 200, persenDiskon: 0, expected: 200},
		{hargaAwal: 100, persenDiskon: 50, expected: 50},
		{hargaAwal: 1500, persenDiskon: 0, expected: 1500},
	}

	for _, test := range listTest {
		t.Run("Testing HitungDiskon", func(t *testing.T) {
			result := HitungDiskon(test.hargaAwal, test.persenDiskon)
			if result != test.expected {
				t.Errorf("Expected %.2f, but got %.2f", test.expected, result)
			}
		})
	}
}
