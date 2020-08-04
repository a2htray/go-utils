### go utils collection

this package contains a set of generally useful function for common development

#### files

`ReadAllLines`

```go
lines, err := files.ReadAllLines("/filesystem/a.txt")
```

`Reomve`

```go
err := files.Remove("/filesystem/a.txt")
// or
err := files.Remove("/filesystem/a.txt", "/filesystem/b.txt")
```

#### slices

`InterfaceForeach`

```go
ss := []interface{}{
    "1",
    "2",
    "3",
}
newSS := make([]int, 0, len(ss))
InterfaceForeach(ss, func(i interface{}) {
    o, _ := strconv.Atoi(i.(string))
    newSS = append(newSS, o)
})
fmt.Println(newSS)
```

`StringIn`

```go
StringIn([]string{"1", "2", "3"}, "3")
```

`StringMap`

```go
ss := []string{"1", "2", "3"}

fmt.Println(StringMap(ss, func(s string) interface{} {
    return s + "@"
}))
```

`StringForeach`

```go
ss := []string{"1", "2", "3"}
StringForeach(ss, func(s string) {
    fmt.Println(s)
})
```

`StringUnique`

```go
ss := []string{"1", "2", "2"}
fmt.Println(StringUnique(ss))
```

`StringExtend`

```go
ss := []string{"1", "2", "2"}
s1 := []string{"4"}
s2 := []string{"5"}
fmt.Println(StringExtend(ss, s1, s2))
```

#### strings

`RegexSplit`

```go
s := "a2htray@.#yuen"
for _, v := range RegexSplit(s, `[@.#]+`, -1)  {
    fmt.Println(v)
}
```
