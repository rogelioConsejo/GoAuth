package auth

import (
	"fmt"
	"testing"
)

func TestGenerarToken(t *testing.T) {
	var i uint8
	for i = 1; i < 10 ; i++  {
		fmt.Printf("%s-%s\n",generarToken(5),generarToken(20))
	}
	generarToken(20)
}
