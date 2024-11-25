package main

func HitungDiskon(hargaAwal float64, persenDiskon float64) float64 {
	return hargaAwal * (1 - persenDiskon/100)
}
