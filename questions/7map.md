# В какой последовательности будут выведены элементы map[int]int?

### Пример:
#### m[0]=1
#### m[1]=124
#### m[2]=281




### Последовательность может быть разной 

```
package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[0] = 1
	m[1] = 124
	m[2] = 281

	for k, v := range m {
		fmt.Printf("Key: %d Value: %d\n", k, v)
	}

	//Key: 2 Value: 281
	//Key: 0 Value: 1
	//Key: 1 Value: 124

	//or

	//Key: 2 Value: 281
	//Key: 0 Value: 1
	//Key: 1 Value: 124
}

```
