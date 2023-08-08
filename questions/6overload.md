# Есть ли в Go перегрузка методов или операторов?



### Нет, в go нет перегрузки методов и операторов
### Нужно использовать разные имена функций
```
func f(a int)        {}
func f(a int, b int) {} //Dublicate error
func twof(a int, b int) {} // OK

func main() {

}
```
