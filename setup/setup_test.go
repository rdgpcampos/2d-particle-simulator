package setup

import (
	"testing"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
	"os"
)

func TestRunSetupFile(t *testing.T) {
	RunSetupFile(lib.HomePath+"lib/test-run-file.txt")

	if _, err := os.Stat(lib.HomePath+"lib/test-run-file.txt"); err != nil {
		t.Errorf(err.Error())
	}

}