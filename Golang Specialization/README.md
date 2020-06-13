# Week 1 - Getting Started with Go

#### About Go:

* efficient as C
* sharing
* runs fast
* garbage collection
* simpler objects
* essentially OOP (there is a concept of object)
* object is simplified => makes it faster/simple to use
* concurrency is efficient
* machine, assembly, high-level language
* machine language is executed on CPU; very simple add, mul...
* assembly like: almost machine language but is human readable
* compiled language with some features of interpreted languages, like Garbage Collection
* weakly OOP:
    * go does not use term “class” but “struct”s (containing data) with associated methods
    * no inheritance, no generics, no constructors
* its implementation of concurrency

#### Hierarchy

Go recommends directory hierarchy: common organization is good for sharing;
* src, pkg (libraries), bin
* one workspace for many directories

GOPATH: workspace directory; defined during installation
package: group of related source files; can be imported by other packages. first line in source file names the package


file1.go:

package annpkg


file2.go:

package billpkg

file3.go:

import (
   “annpkg”
   “billpkg”
)

* There must be one package called main; that’s where execution starts; compiled main package is an executable; other packages compile into packages
* main package has to have function main

```go
package main
import “fmt” // note: it’s NOT include!!!

func main() {
   fmt.Printf(“Hello, world\n”)
}
```

Import - used to access other packages
Import declarations in Go
Go standard library includes many packages, like fmt
Go searches packages in GOROOT and GOPATH directories first;

#### Go Command line:

go build - to compile the source code; when compiling main package.exe file has the same name as .go file
go doc - prints documentation for package
go format - formats source code files (e.g. indentation...)
go get - downloads and installs new packages
go list - installs all installed packages
go run - compiles and executes
go test - runs tests using files ending in “_test.go”

#### Naming

variables: have to have name and type
```go
var x int
var x, y int
```

type defines what type of data and which operations can be performed on the variable

To improve clarity/readability alias name for type (with type keyword):

```go
type Celsius float64
type IDnum int

var temp Celsius
var pid IDnum 

//initializing variables in the declaration with explicit type specification


var x int = 100 //with implicit type inference

var x = 100 // aise bhi kar sakte hai if that floats your boat lol

var x int
x = 100
// uninitialized variables have value of 0, empty string


x := 100
// short variable declaration (:=): declaration & initialization (type is inferred; can be done only inside a function):

```

# Week 2 - Basic Data Types

#### Pointers

* &(ampersand) - gets address of variable or function
* *(star operator) - returns data at the address; comes before pointer

```go
var x int = 1
var y int
var ip *int
ip = &x
y = *ip
```

new() - another way to create a variable; it returns a pointer to a variable; variable is initialized to 0 by default

```go
ptr := new(int)
*ptr = 3
```

#### Scope

variable scope - where variable can be accessed

in Go, variable scoping is done through blocks(sequence of declarations and statements within matching{})

blocks can be hierarchical

* Explicit blocks: when you write {} yourself; e.g. functions - each function has its own block
* Implicit blocks:
    * universe block - all Go source code
    * package block - all source in a package
    * file block - all source in a file (package can contain many files)
    * if, for, switch statements
    * clause in switch or select

* Lexical scoping - this defines how variable references are resolved; Go is a lexically scoped languageusing blocks.

bi >= bj if bj is defined inside bi ; “defined inside” is transitive
variable is looked in the current/local block and then in the next bigger block
Variable is accessible from block bj if:
* variable is declared in bi and
*  bi >= bj

variables have to be allocated and deallocated (when is not used anymore so memory space is made available) in memory; this has to be done in a timely fashion - before we run out of memory

### stack vs heap
* Stack (traditionally, not in Go): dedicated to function calls; local variables for functionareallocated in stack; they are deallocated automatically when function returns;
* Heap: persistent memory - you have to explicitly deallocate it; in C: malloc - to allocate; free - todeallocate it; error-prone (if you forget to call free) but fast;

### Garbage Collection

you don’t want to deallocate memory prematurely
How to determine when a variable is no longer in use?

```go
func foo() *int {
   x := 1
   return &x
}

func main() {
   var y *int
   y = foo()
   fmt.Printf(”%d”, *y)
}
```
* GC is done by interpreter; GC keeps track of pointers/references to variables; only when all referencesare gone, it deallocates variable; GC keeps track of;
compiled languages like C and C++ can’t do this
* Go is compiled language which has GC built in
* Go compiler determines whether variable will go on stack or on heap; you don’t need to determine that
* GC takes some time - that’s the tradeoff

```go
fmt.Printf():
```
* conversion characters %d, %s
* https://stackoverflow.com/questions/53961617/call-has-possible-formatting-directive
```go
fmt.Println()
```
* does NOT support conversion/formatting characters
* adds a new line character at the end
* in the output, adds a SPACE character after each argument
* it's not necessary to add a SPACE explicitly at the end of "a = ", "a =" will work ok:

```go
fmt.Println("a =", a)
```

### Integers
* int
* int8, int16, int32, int64 (2s complement: most significant bit is used for sign)
* unsigned: uint8...etc
*int vs int32 vs in64
* int is the default signed type: it takes 32 bit (4 bytes) on a 32 bit machine and 64 bit (8bytes)on a64 bit machine; the same goes for the unsigned uint. [Why does Go have severaldifferentinteger types?]
* The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits  wideon64-bitsystems. When you need an integer value you should use int unless you have aspecificreason to use asized or unsigned integer type [A Tour of Go - Basic Types]
* Pick the right one: int vs. int64

Integers: binary operators
* arithmetic
* comparison
* boolean
Type conversions
    most binary operations (including assignment) require both operands to be of the same type
```go
var x int32 = 1
var y int16 = 2
x = y
// [go] cannot use y (type int16) as type int32 in assignment
```
##### convert types with T() operation

x = int32(y)

### Floating point
* float32 ~6 digits of precision
* float64 ~ 15 digits of precision
* decimal representation: x float64 = 123.45
* scientific representation: x float64 = 1.2345e2 (e2 = 10^2; base 10)
* complex numbers: var z complex128 = complex(2, 3) // real, imaginary


### ASCII and UNICODE
* each character has to be encoded
* ASCII: 7 or 8 bits: ‘A’ = 0x41
* UNICODE: 32-bit long code (2 to 32 characters can be represented)
* UTF-8: variable length - 8-bit codes same as ASCII; default in Go
* code point (or “rune” in Go): term for a unicode character

### string
* sequence of arbitrary bytes
* array of runes
* read-only (immutable)
* string literal - between “”

### Unicode package
* IsDigit(r rune)
* IsSpace(r rune)
* IsLetter(r rune)
* IsLower(r rune)
* IsPunct(r rune)
* ToUpper(r rune)
* ToLower(r rune)

### Strings package - looks string as a whole
* Compare(a, b) - lexicographical comparison; returns -1, 0, 1
* Contains(s, substr) - returns true/false
* HasPrefix(s, prefix) - returns true/false
* Index(s, substring) - returns the index of the 1st instance
* Replace(s, old, new, n) - returns a copy of s where first n instances of old are replaced by new
* ToLower(s) - returns new string
* ToUpper(s) - returns new string
* TrimSpace(s) - returns new string
* Go Walkthrough: bytes + strings packages

### Strconv package (string conversion)
* Atoi(s) - convert string to int
* Itoa(n) - converts int to string
* FormatFloat()
* ParseFloat()

### Constants
* expressions whose value is known in compile time
* values inferred from right hand side
* iota - function used to generate related but distinct constants (e.g. days of week) - value is not important, we don’t care about value; like an enumeration; each const is assigned to a unique integer


* In the current implementation iota values start from 1 but this is not guaranteed and can change in future
type Grades int
```go
const (
 A Grades = iota
 B
 C
 D
 E
 F
)

fmt.Printf("C = %d", C) // output: C = 2
```
## Control flow - determines the order of execution of statements

### For loops:
    * iterates while condition is true; condition is always required
    * may have initialization (executed only once) and update (executed at the end of each loop):

```go
for ;; { }
```

### Switch 
* contains tags (labels)
* tagless switch
* break - exits the containing loop
* continue - skips the current iteration

### Scan
* reads user input
* waits till user types in something and presses ENTER
* returns number of scanned items (separated by SPACE) and error
* has pointer as an argument
To scan user input which contains SPACE characters we need to use buffer:
* String with spaces
* Golang : Accept input from user with fmt.Scanf skipped white spaces and how to fix it


# Week 3 - Composite Data Types

Composite data types - aggregate other data types

### Arrays

* fixed-length series of elements of a chosen type
* in compile time is known the length of it - compiler knows how much memory is needed
* each element is indexed by using subscript notation - index ([ ])
* indices start from 0
* elements are initialized to zero value of the type (0 for integers, empty string for strings etc...)

```go
var x[5] int // array of 5 integers
x[0] = 2
fmt.Printf(x[1]) // 0
```

### Array literal
* predefined set of values that make up an array
* used to initialize an array
* length of literal must be length of array

```go
var x[5] int = [5] {1, 2, 3, 4, 5}
```
... (three dots) - keyword; used to express the size of an array literal inferred from the list of values; in the previous example:

```go
x := [...] {1, 2, 3, 4, 5} //  size of the array literal is 5 
```
iterating through array: use a for loop with range keyword; 
range returns two values: index and value at that index
```go
x := [3] int {1, 2, 3}
for i, v := range x {
   fmt.Printf("index = %d; element value = %d", i, v)
}
```

Only slices, channels and arrays can be used in for-range.

### Slices

* lots of time are used instead of arrays
* there must be some underlying array which is a basis of the slice
* slice is a window on the underlying array (larger or of the same size)
* can change their size; have variable size, up to the size of array
* has 3 properties:
    * Pointer - indicates the start of the slice; points to the element of the array where slice starts
    * Length - number of element in the slice
    * Capacity - max number of elements; from start of slice to the end of array; e.g. if slice starts at 10th element of the array which has 100 elements => capacity is 90
* colon (:) is used between m and n; 
* m - index of (pointer to) the 1st element in the array which is IN the slice 
* n - index of (pointer to) the 1st element in the array which is NOT in the slice

```go
arr := [...]{"a", "b", "c", "d", "e", "f", "g", }
slice1 := arr[1:3] // {"b", "c",}
slice2 := arr[2:5] // {"c", "d", "e",}
```

* len() - returns the length of the slice
* cap() - returns the capacity of the slice (number of elements in the underlying array from the beginning of the slice to the end of the array)

a2 := [3]string{"a", "b", "c"}
slice21 := a2[0:1] // { "a" }
fmt.Println(len(slice21), cap(slice21)) // 1, 3

    writing in the element of slice writes actually to the underlying array
    overlapping slices can refer to the same element of the array; writing in one slice changes another slice 
    slice literals
        can be used to initialize a slice
        creates the underlying array and creates a slice to reference it
        slice points to the beginning of the array and length is  capacity

slice := [] int {1, 2, 3}

## Variable Slices

    3rd way to make slice: make() function
        2-argument version: type (elements initialized to 0, empty string etc...), length (equal to capacity): 

slice :=  make([]int, 10)

        3-argument version: specify length and capacity separately: e.g. 10 is length and 15 a capacity

slice := make([]int, 10, 15) 

    size of the slice can be increased by append()
        adds elements at the end of the slice
        inserts into underlying array
        increases the size of the array if necessary
        we can keep appending as we wish

slice := make([]int, 0, 3) // length slice is 0 but the size of underlying array is 3

slice := append(slice, 100) // 100 is the number (element) we want to add to the slice

## Hash Tables

    contains key-value pairs
    each value is associated with the unique key
    hash function
        used to compute a slot for a key
        its argument is the key, the output is the slot in the "array"
        it's never called explicitly, it's part of the hash table implementation
        e.g. {"Joe", "x" }, {"Jane", "y" }, {"Pat", "z" }. Hash function takes keys (names) and outputs e.g. slots 3, 1, 5 (indexes in the array/slice where values are stored)
        we never use indexes/slot numbers but only keys e.g. hashTable["Jane"] would return "y"
    access is in constant time
    naming is useful as it's easier to use keys (e.g. names) rather than indexes (numbers)
    advantages: 
        faster lookup than list: hash table has constant time - it's always the same, no matter how many key-value pairs is in the table; list has linear lookup time - the more elements in list, the more time it takes to find some particular element
        arbitrary keys, not integers like in slices or arrays; keys have meaning
    disadvantages:
        may have collisions: when hash function maps two different keys into the same slot
            they are rare; hash functions are made in such way to minimize them

## Maps

    Golang's implementation of hash table

var mapId map[string]int // string is the key type and int is the value type
idMap = make(map[string]int) // creating a (empty) map

    we can define map literal used to initialize a new map:

idMap := map[string]int {"joe": 123}

    accessing maps
        referencing a value with [key]: idMap["Joe"]
            returns 0 if key is not present
        to add a new key-value pair (or modifying the value of the existing key): idMap["Jane"] = 456
        to delete a key-value pair: delete(idMap, "Joe")
    map functions
        two-value assignment tests for existence of the key; if "Joe" key exists, p (presence of the key) will be true and id will be its value; if we don't care for the value, we can use "_" as the name instead of "id"

id, p := idMap["Joe"]

        len() returns number of key-value pairs in the map
    Iterating through map
        use for loop with the range keyword
        two-value assignment with range

for key, value := range idMap {
   fmt.Println(key, value)
}

## Structs

    aggregate/composite data type
    groups data of various types

type Person struct{
   name string
   address string
   phone string
}

var p1 Person

    each property is called a field
    to access fields (read/write) - use dot notation:

p1.name = "Joe"
x1 = p1.address

    initializing structs
        can use new()
            initializes fields to 0 

p1 := new (Person)

        by using a struct literal

p1 := Person { 
   name : "Joe", 
   address : "A street", 
   phone : "123"
}

# Week 4 - Protocols and Formats
## RFCs

    definitions of Internet protocols and formats
    examples:
        HTML
        URI 
        HTTP
    Go provides protocol packages
        contain functions that encode/decode protocol format
    "net/http"
        web communication protocol
        e.g. http.Get(www.uci.edu)
    "net"
        TCP/IP  and socket programming
        e.g. net.Dial("tcp", "uci.edu:80")
    JSON
        JavaScript Object Notation
        Format to represent structured information
        attribute-value pairs
            naturally correlate to structs/maps
        basic value types: bool, number, string, array, object
            can be combined hierarchically
        Go struct:

p1 :=  Person(name: "Joe", addr: "a street", phone: "123")

        Equivalent JSON object:

{ "name": "Joe", "addr": "a street", "phone": "123" }

## JSON

    all is UNICODE
    human readable
    fairly compact
    types can be combined recursively
        arrays of structs, structs of structs...
    Go package: "encoding/json"
    JSON marshalling
        generating JSON representation from an object 


type Person struct{
   name string
   addr string
   phone string
}

p1 := Person(name: "Joe", addr: "a st.", phone: "123")
barr, err := json.Marshal(p1)


    Marshal() returns JSON representation as []byte
        err is nil if there was no error
        barr is byte array that contains JSON representation
    Unmarshal() converts a JSON []byte into Go object
        second argument is the address of the object into which JSON will be deserialized; its fields have to match JSON object properties (and the whole structure); object must "fit" JSON 

var p2 Person
err := json.Unmarshal(barr, &p2)


## File Access, ioutil

    File access is liner, not random access
    files used to be stored on physical tapes
        mechanical delay when reading as file beginning is on one end of the tape and end on the other
    physical disks still have linear access
    RAM (flash memory) has random access
    basic file operations:
        Open - get handle for access
        Read - read bytes into []byte
        Write - write []byte into file
        Close - release handle
        Seek - move read/write head

    there is more than one package in Go which handles file access, ioutil [1] is one of them
    ioutil.ReadFile()
        dat is []byte filled with the content of the entire file
        explicit open/close are not needed
        large files can cause a problem: if their size is same or larger than available RAM memory e.g. loading 8GB file to 8GB RAM

dat, e := ioutil.ReadFile("text.txt")

    ioutil.WriteFile()
        creates a new file
        3rd arg is a read/write permission, 0777 = every user can read/write


dat = "Hello, world"
err := ioutil.WriteFile("outfile.txt", dat, 0777)

File Access, os

    ioutil provides simple utility
    for more control over file access, use os package
    os.Open()
        opens a file
        returns file descriptor
    os.Close()
        closes a file
    os.Read()
        reads from a file into a byte array ([]byte)
        we can control how much to read
    os.Write()
        writes bytes into the file
        writes as much as byte array is long
    Opening and reading:
        os.Read() returns number of bytes read; can be smaller than the size of barr

f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.Close()

    File Create / Write
        os.Write() - writes byte array
        os.WriteString() - writes Unicode sequence

f, err := os.Create("outfile.txt")
barr := []byte {1, 2, 3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")

---

# Bonus 

* Project/Code Organization

* How should I structure packages for a multiple-binary web application?
* Go: Best Practices for Production Environments

* Naming Conventions
* Which way to name a function in Go, CamelCase or Semi-CamelCase?

* How to return both value and error from a function
* Go by Example - Errors
 
* Unit Testing
* Package testing
* How to test a unexported (private) function in go (golang)?


References:
https://golang.org/doc/
