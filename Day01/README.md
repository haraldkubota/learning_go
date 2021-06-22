# learning_go

Learning Go

## Day01
* Go initializes everything to zero (false, empty string)
* Type conversions are like C, but they have to be explicit
* multiple assignments is common:
** a,b := 3,4
** res,err = function()
* Uses type inference
* "" and \`\` is a string, '' is a rune (=int32)
* complex numbers are a primitive
* THAT is useful! Define and set v, then use in a condition and later on, even in the else block
```
    if v := math.Pow(x, n); v < lim {
        return v
    }
```
* Methods start with a capital (e.g. math.Abs())
* switch cases have no fall-through
* defer runs at the end of the surrounding function, and they get executed in LIFO order
* Pointers just like in C, just more universal
* structs pointers do not need to explicitely dereference: No need for p->member and instead p.member
* arrays are low-level slices with fixed size. Slices is more like lists in Python (same syntax too for slicing) with no fixed size
* sliced just point to arrays
* append(slice, ...) dynamically increases the array if needed
* for k, v := range array_or_slice works
* Function closures work like JS
* Anynymous functions work just like in JS too
* No classes, but methods instead (with receivers as types)
* Pointer receivers to modify data inside the method:
```
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
```
* Value receivers auto-convert pointers to values as needed. Normal functions throw an error instead.
* Act on types:
```
func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}
```
* multithreading is a "go" away
