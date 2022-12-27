package utils

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"strings"
)

// IsTestFile 判断当前文件是不是单元测试文件，文件名为 xxx_test.go 的就是
func IsTestFile(pass *analysis.Pass, n ast.Node) bool {
	if file := pass.Fset.File(n.Pos()); file != nil && strings.HasSuffix(file.Name(), "_test.go") {
		return true
	}
	return false
}
