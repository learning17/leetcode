package main
import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	im := big.NewInt(math.MaxInt64)
	io := big.NewInt(1956)
	im.Mul(im, io)
	fmt.Println(im)
}
