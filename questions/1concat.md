# Какой самый эффективный способ конкатенации строк?

## Использовать strings.Builder

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder //Создаём объект builder
	builder.WriteString("Hello") //Добавляем строки
	builder.WriteString(" ")
	builder.WriteString("world!")

	result := builder.String() // Получаем итоговую строку
	fmt.Println(result)
}
```
## string.Builder позволяет избежать создания промежуточных строк и копирования памяти