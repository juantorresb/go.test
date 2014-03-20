package first

import (
	"testing"
)


func TestPairNumber(t *testing.T) {
	
	isPair := IsPairNumber(5)
	if isPair  {
		t.Error("Expected no pair number")
	}
	
	isPair = IsPairNumber(4)
	if !isPair  {
		t.Error("Expected pair number")
	}
}

func TestSquare(t *testing.T) {
	
	v := Square(5)
	if v != 25 {
		t.Error("Expected 25, got ", v)
	}
	
	v = Square(9)
	if v != 81 {
		t.Error("Expected 81, got ", v)
	}
}


func TestAverage(t *testing.T) {
	
	v := Average(5, 3, 7)
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}
/*
func ExampleHello() {
	fmt.Println("hello")
   // Output: hello"
}*/