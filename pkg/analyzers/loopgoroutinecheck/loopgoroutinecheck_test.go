package loopgoroutinecheck

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %s", err)
	}
	testdata := filepath.Join(wd, "testdata")
	tests := []string{"data"}
	analysistest.Run(t, testdata, Analyzer, tests...)
}
