package main

import (
	"fmt"
	"tugas1-golang-Krise/controller"
	"tugas1-golang-Krise/konversiCelcius"
)

func main() {
	var itf konversiCelcius.Hitung = controller.Nilai{
		NilaiCelcius: 20.00,
	}

	fmt.Println("Hasil Konversi 20 Celcius ke Fahreinhet = ", itf.Fahreinhet())
	fmt.Println("Hasil Konversi 20 Celcius ke Reamur = ", itf.Reamur())
	fmt.Println("Hasil Konversi 20 CeLcius ke Kelvin = ", itf.Kelvin())
	fmt.Println("Hasil Konversi 20 Celcius ke Rankin = ", itf.Rankin())
}
