Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

test() возвращает 2, так как defer работает с именованной возвращаемой переменной. anotherTest() возвращает 1, так как defer изменяет переменную x после того, как её значение уже было скопировано для возврата из функции.
```
