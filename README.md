# gennest
Generate unlimited nested functions in Go
# Usage
Run:

`go run *.go --level 10  --file test.go --funcName test`

This command generates function with 10 levels of nest and prints "Hello"(subject to change). Also it outputs a function call to stdout for your convinience.
Stdout output:

```
function call bozo :
 test()()()()()()()()()()()
```

test.go:

```
package main
import "fmt"
func test() func() func() func() func() func() func() func() func() func() func() {return func() func() func() func() func() func() func() func() func() func() {return func() func() func() func() func() func() func() func() func() {return func() func() func() func() func() func() func() func() {return func() func() func() func() func() func() func() {return func() func() func() func() func() func() {return func() func() func() func() func() {return func() func() func() func() {return func() func() func() {return func() func() {return func() {fmt.Println("Hello")}}}}}}}}}}}
```
