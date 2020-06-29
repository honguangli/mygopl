package tempconv

import (
	"fmt"
	"testing"
)

func TestCToK(t *testing.T) {
	fmt.Println(CToK(Celsius(AbsoluteZeroC)))
	fmt.Println(CToK(Celsius(15)))
	fmt.Println(CToK(Celsius(BoilingC)))
}

func TestKToC(t *testing.T) {
	fmt.Println(KToC(Kelvin(0)))
	fmt.Println(KToC(Kelvin(288.15)))
	fmt.Println(KToC(Kelvin(373.15)))
}
