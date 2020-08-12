### go utils collection 

![Build Status](https://travis-ci.com/a2htray/go-utils.svg?branch=master)

this package contains a set of generally useful function for common development

#### files

`ReadAllLines`

```go
lines, err := files.ReadAllLines("/filesystem/a.txt")
```

`Remove`

```go
err := files.Remove("/filesystem/a.txt")
// or
err := files.Remove("/filesystem/a.txt", "/filesystem/b.txt")
```

`Extention`

```go
Extension("\\file\\a.txt") // txt
```

`ExtentionWithDot`

```go
ExtensionWithDot("\\file\\a.txt") // .txt
```

#### pretty

`Print`

```go
type V struct {
    A int `json:"a"`
    B string `json:"b"`
}
a := []interface{}{V{
    A: 1,
    B: "a",
}}
Print(a...)

b := []interface{}{1, 2, 3}
Print(b...)
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

fmt.Println(StringMap(ss, func(s string) string {
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
