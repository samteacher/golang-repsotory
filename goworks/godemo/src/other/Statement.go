package other

import "fmt"

//const boilingF = 212.0

func Statement() {

	//var f = boilingF
	//var c = (f - 32) * 5 / 9
	//fmt.Printf("boilingF point  = %g°F or %g°C\n", f, c)

	const freezingF, boilingF = 32.0, 212.0
	// "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	// "212°F = 100°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
