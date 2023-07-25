package main

func generateWithReturns(level int, expr string) string {
	if level == 0 {
		return expr
	}
	res := "return "
	for i := 0; i < level; i++ {
		res += "func() "
	}
	res += "{" + generateWithReturns(level-1, expr)
	return res
}

func generateRest(level int) string {
	res := ""
	for i := 0; i < level; i++ {
		res += "}"
	}
	return res
}

func Generate(level int, expr string, funcName string) string {
	return generateHeader(level, funcName) + generateWithReturns(level, expr) + generateRest(level+1)
}

func generateHeader(level int, funcName string) string {
	res := "func " + funcName + "() "
	for i := 0; i < level; i++ {
		res += "func() "
	}
	return res + "{"
}

func GenerateFuncCall(level int, funcName string) string {
	res := funcName
	for i := 0; i < level+1; i++ {
		res += "()"
	}
	return res
}
