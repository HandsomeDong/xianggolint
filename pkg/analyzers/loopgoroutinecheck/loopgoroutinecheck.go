package loopgoroutinecheck

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"xianggolint/utils"
)

const name = "loopgoroutinecheck"

var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "the loop variable captured by func literal in go statement might have unexpected values",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.RangeStmt)(nil),
		(*ast.ForStmt)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if utils.IsTestFile(pass, n) {
			return
		}
		var vars []*ast.Ident
		addVar := func(expr ast.Expr) {
			if id, ok := expr.(*ast.Ident); ok {
				vars = append(vars, id)
			}
		}
		var body *ast.BlockStmt
		switch n := n.(type) {
		// e.g. for index, val := range values
		case *ast.RangeStmt:
			body = n.Body
			addVar(n.Key)
			addVar(n.Value)
		case *ast.ForStmt:
			body = n.Body
			switch post := n.Post.(type) {
			// e.g. for p = head; p != nil; p = p.next
			case *ast.AssignStmt:
				for _, lhs := range post.Lhs {
					addVar(lhs)
				}
			// e.g. for i := 0; i < n; i++
			case *ast.IncDecStmt:
				addVar(post.X)
			}
		}
		if len(vars) == 0 || len(body.List) == 0 {
			return
		}

		ast.Inspect(body, func(node ast.Node) bool {
			if goStmt, ok := node.(*ast.GoStmt); ok {
				checkGoStmt(pass, goStmt, vars)
			}
			return true
		})

	})
	return nil, nil
}

func checkGoStmt(pass *analysis.Pass, goStmt *ast.GoStmt, vars []*ast.Ident) {
	fun := goStmt.Call.Fun
	if funcLit, ok := fun.(*ast.FuncLit); ok {
		ast.Inspect(funcLit, func(node ast.Node) bool {
			id, ok := node.(*ast.Ident)
			if !ok || id.Obj == nil {
				return true
			}
			if pass.TypesInfo.Types[id].Type == nil {
				return true
			}
			// 对比是不是循环的key、index或value对象
			for _, v := range vars {
				if v.Obj == id.Obj {
					pass.ReportRangef(id, "%s: loop variable `%s` captured by func literal in go statement might have unexpected values", name, id.Name)
				}
			}
			return true
		})
	}
}
