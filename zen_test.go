package zen_test

import (
	"testing"

	"github.com/archmangler/zen"
)

func TestKoan(t *testing.T) {
	if zen.Koan() != "Seijo's soul separated from her being. Which was the real Seijo?" {
		t.Fatal("Wrong Realization. No Satori for you!")
	}
}
