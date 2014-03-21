package log

import (
	"testing"
)


func testLogFile(t *testing.T) {
	isok := Logger()
	if !isok  {
		t.Error("Expected ok return")
	}
}
