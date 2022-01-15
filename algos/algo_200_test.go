package algos_test

import (
	"testing"

	"github.com/keithyin/lc-go/algos"
)

func TestAlgo200(t *testing.T) {
	grip := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}

	val := algos.NumIslands(grip)
	if val != 3 {
		t.Errorf("expected 1, but got %d\n", val)
	}
}
