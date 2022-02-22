package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum uint) float64 {
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(sideLen, sidesNum)
	switch sidesNum {
	case 0:
		return math.Pi * math.Pow(sideLen, 2)
	case 3:
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
