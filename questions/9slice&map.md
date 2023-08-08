# Сколько существует способов задать переменную типа slice или map?

###
```

arr1 := []int{}
arr2 := []int{1,2,3}
var arr3 []int
arr4 := make([]int, 0)
arr5:= new([]string)


var m1 map[string]string
m2 := map[string]string{
    "123":  "456",
    "test": "hello",
}
m3 := make(map[string]string)
m4 := make(map[string]string, 10)
m5 := new(map[string]string)
```