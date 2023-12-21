Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

проверка if err != nil проверяет, не равно ли значение интерфейса nil. В этом случае, несмотря на то, что test() возвращает nil, интерфейс err фактически не равен nil, потому что он хранит тип *customError
```
