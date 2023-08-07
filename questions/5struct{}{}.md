# Какой размер у структуры struct{}{}?



### 0 байт
```
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := struct{}{}
	fmt.Println(unsafe.Sizeof(a)) // 0
}
```
