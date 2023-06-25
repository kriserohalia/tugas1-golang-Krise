package controller

import (
	"tugas1-golang-Krise/konversiCelcius"
)

func (a Nilai) Fahreinhet() float64 {
	hasil := (1.8 * a.NilaiCelcius) + 32
	konversiCelcius.Connect.PostFahreinhet(konversiCelcius.Connect{}, hasil)
	return hasil
}

func (a Nilai) Reamur() float64 {
	hasil := (0.8 * a.NilaiCelcius)
	konversiCelcius.Connect.PostReamur(konversiCelcius.Connect{}, hasil)
	return hasil
}

func (a Nilai) Kelvin() float64 {
	hasil := a.NilaiCelcius + 273.15
	konversiCelcius.Connect.PostKelvin(konversiCelcius.Connect{}, hasil)
	return hasil
}

func (a Nilai) Rankin() float64 {
	hasil := (a.NilaiCelcius + 273.15) * 1.8
	konversiCelcius.Connect.PostRankin(konversiCelcius.Connect{}, hasil)
	return hasil
}
