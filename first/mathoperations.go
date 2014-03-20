package first

import (
)

func Square(value int) int{
	return value*value
}

func Average(values ... float64)float64 {
	var sum float64
	size := len(values)
	
	for _, val := range values {
		sum += val
	}

	return sum/float64(size)
}

func IsPairNumber(value int) bool {
	isPair := false
	
	val := value %2
	if  val == 0 {
		isPair = true
	}
	return isPair
}