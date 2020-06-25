# Week 1 - FUNCTIONS & ORGANIZATION

#### Why Use Functions?

function - set of instructions grouped together, usually with a name
all programs in Go have to have function main() - where program's execution begins
```go
func main() {
   fmt.Printf("Hello, world!")
}
```

main() function is a special in a sense we don't call it explicitly which is not the case for other functions
```go
func PrintHello() {
   fmt.Printf("Hello, world!")
}

func main() {
   PrintHello()
}
```
* function declaration starts with keyword func, function name, arguments, return type and body
* Why using functions?
    * reusability (within the same project or across multiple projects via libraries)
    * abstraction
    * hiding details of implementation; only input-output behaviour is what we need to know (we look at function as a black box)
    * improves understandability
    * naming 
    * grouping of function calls


#### Function Parameters and Return Values

* functions need some data to work on - they can be passed via function parameters
* parameters are listed in parentheses after function name
* arguments are supplied in the call
```go
func foo(x int, y int) {
   return x * y
}

func main() {
   foo(2, 3)
}
```
#### Parameter Options

* if no parameters are needed, put nothing in parentheses; you still need parentheses
* list arguments of same type
```go
func foo(x, y int) {
}
```

#### Return Values
* Functions can return value as result
* type of return value after parameters in declaration
* function call is on the right side of an assignment
```go
func foo(x int) int {
   return x + 1
} 

y := foo(1) // y gets assigned value 2
```

* Functions can have multiple return values
* their types must be listed in the declaration
```go
func foo2(x int) (int, int) {
   return x, x + 1
}

a, b := foo2(3) // a is assigned 3, b 4
```
#### Call by Value, Reference
* how are arguments passed to parameters during the function call

##### Call by Value
* arguments are copied to parameters
* data used is the copy of the original;
* called function can't interfere with the original variables in the calling function
* modifying parameters has no effect outside the function
```go
func foo(y int) {
   return y + 1
}

func main() {
   x := 2
   foo(x)
   fmt.Print(x) // still 2
}
```

* tradeoffs of call by value
* advantage: 
    * data encapsulation
    * function variables are changed only inside the function
* disadvantage:
    * copying time
    * large object may take a long time to copy
#### Call by Reference
* programmer can pass a pointer as an argument
* called function has direct access to caller variable in memory
* we don't pass data but a pointer to data (the address of data)
```go
func foo(y *int) {
   *y = *y + 1
}

func main() {
   x := 2
   foo(&x) // foo gets a copy of the address where x is in memory; foo is modifying x
   fmt.Print(x) // 3
}
```

* tradeoffs of Call by Reference
* advantages:
    * copying times; only the address is copied, not data
* disadvantage:
    * (no) data encapsulation
    * called function can change data/variables in the calling function

##### Passing Arrays and Slices

*how to pass an array to a function?
    *array arguments are copied
    *array can be big so this can be a problem
```go
func foo(x [3]int) int { // NOTE size of array => argument is of type array (not slice)
   return x[0]
}

func main() {
   a := [3]int{1, 2, 3}
   fmt.Print(foo(a)) // entire array is copied!
}
```
* instead, we can pass a reference - a pointer to the array (array pointer)
* this can be messy for having to reference and dereference
```go
func foo(x *[3]int) {
   (*x)[0] = (*x)[0] + 1
}

func main() {
   a := [3]int{1, 2, 3}
   foo(&a)
   fmt.Print(a) // [2 2 3]
}
```
* proper Go approach is to pass slices instead of arrays
* in Go: get used to use slices instead of arrays! you can almost always use slices instead of arrays
* slice is a structure that contains 3 things:
    * pointer to an underlying array
    * length
    * capacity
* if we pass a slice, we actually pass a copy of the pointer to an array
* function can use that pointer directly, without the need to dereference/reference explicitly
```go
func foo(sli int) int {
   sli[0] = sli[0] + 1
}

func main() {
   a := []int {1, 2, 3} // NOTE no size is specified between [] => this is a SLICE declaration!!!
   foo(a)
   fmt.Print(a) [2 2 3]
}
```
* instead of arrays or pointers to arrays pass slices to functions!!!
* Well-Written Functions
* how you should write functions so your code is well-organized and understandable
* Understandability
* code = functions + data
* if you are asked to find a feature, you should be able find it quickly; also, your peer reviewers should also be able to find it quickly
* if you are asked where data is used (defined, accessed, modified), you should be able to find it quickly
* Debugging Principles
    * e.g. code crashes inside a function
* two options for a cause:
* function is written incorrectly
    * e.g. sorts a slice in a wrong order
* data that function uses is incorrect
    * e.g. sorts slice correctly but slice has wrong elements in it
* Supporting Debugging
* functions need to be understandable
* determine if actual bahaviour matches desired behaviour - this shall be easy to do
* data needs to be traceable
* we should be able to trace where did data come from
* global variables complicate this
* anybody can write into them
* otherwise we know that calling function passed data...

#### Guidelines for Functions
##### Function Naming
* behaviour should be understandable at a glance
* parameter naming counts too
```go
func ProcessArray(a []int) float {}  // BAD! What kind or processing? What is the meaning of a?

func ComputeRMS(samples []float) float {} // GOOD. We know what f-n does, and what is the argument
```
##### Functional Cohesion
* function should perform only one "operation"
* an "operation" depends on context
* even from a function name we can see that it does only one thing
* if you need to put two or more actions in a function name, that should raise an alert
* merging behaviours makes code complicated
* Reduce number of parameters
* use few parameters
* debugging requires tracing function input data
* it's more difficult with large number of parameters
* if you function has many parameters it may have bad functional cohesion - it does too many things; each functionality ("operation") needs its own set of inputs; 
* group related arguments into a structure
    * e.g. if we want to pass 3 points in space to a function we might need to pass 2 x 3 = 9 arguments

* improvement: define a struct point and pass 3 points: 
* type point struct {x, y, z float}
* best: define a struct triangle and pass a triangle as a single argument:
* type triangle struct {p1, p2, p3 point}


#### Function Guidelines
##### Function Complexity
* function length is the most obvious measure
* functions should not be complex; easier to debug
* short functions can be complex too
##### Function length
* write complicated code with simple functions
* function call hierarchy
* Control-flow Complexity
* how many control-flow paths are there in a function?
* paths from the start to the end of function
* if there are no if statements: one control-flow path
* if there is if statement: two control-flow paths
* control-flow describes conditional paths
* functional hierarchy can reduce control-flow complexity: separate conditional code into separate functions

---

# Week 2 - FUNCTION TYPES
#### First-Class Values
*  Go treats functions as a first-class values
*  functions are First-class
*  Go implements some features of functional programming
*  Go treats functions as any other type like int, float...
*  variables can be declared to be a function type and then assigned a function
*  functions can be created dynamically, on the fly
*  so far we've been creating them dynamically, in the global space we'd use func
*  but they can be created dynamically, inside other functions
*  functions can be passed as arguments to functions
*  functions can be returned from functions
*  functions can be stored in structs
####  Variables as Functions 
*  declare variable as function
*  variable becomes an alias (another name) for that function

```go
var funcVar func(int) int // "func(int) int" is function signature

func incFn(x int) int {
   return x + 1
}

func main() {
   funcVar = incFn // in assignment just use function name, without () as we're not calling function here
   fmt.Print(funcVar(1))
}
```

#### Functions as Arguments
* functions can be passed to other functions as arguments
* we have to use keyword func
```go
func applyIt(afunc func(int) int, val int) int {
   return afunc(val)
}


func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }

func main() {
   fmt.Println(applyIt(incFn, 2)) // 3
   fmt.Println(applyIt(decFn, 2)) // 1
}
```

#### Anonymous Functions
* functions don't need to have names
* functions with no name are called anonymous 
* when passing function to another function you usually don't need to name passed function
* function is created right there at the call 
* this comes from lambda calculus
```go
func main() {
   v := applyIt(func (x int) int { return x + 1}, 2);
   fmt.Println(v) // 3
}
```

#### Returning Functions
Functions as Return Values

* functions can create functions and return them
* new function can have a different set of parameters; controllable parameters
   * example: Distance to Origin function
* takes a point (x, y coordinates)
* returns distance to origin
* what if I want to change the origin?
   * option1: origin becomes a parameter
   * option 2: we create a function for each origin (o_x, o_y)
* origin is build in the returned function
```go
func (float64, float64) float64 is the type of the function returned
func makeDistOrigin(o_x, o_y float64) func (float64, float64) float64 {
   fn := func (x, y float64) float64 {
      return math.Sqrt(math.Pow(o_x, 2) + math.Pow(o_y, 2))
  }
  return fn
}
```

#### Special-Purpose Functions
* we make special-purpose functions by giving them parameters (e.g. Dist1 and Dist2 have different origins)
```go
func main() {
   Dist1 := MakeDistOrigin(0, 0)
   Dist2 := MakeDistOrigin(2, 2)
   fmt.Println(Dist1(2, 2))
   fmt.Println(Dist2(2, 2))
}
```
#### Environment (Scope) of a Function
* every function has an environment ("scope")
* set of all names that are valid inside a function; that you can refer inside the function
* environment includes names defined locally, in the function
* Lexical Scoping: Go is lexically scoped 
* environment includes names defined in block where the function is defined
* BK: this is called variable capturing 
* when you start passing around functions as arguments, the environment goes along with functions 
```go
var x int
func foo(y int) {
   z := 1
   ...
}
```

#### Closure
* function + its environment, together
* in Go, it is implemented as a structure which contains pointer to function and pointer to environment
* when you pass function as an argument to another function, you pass its environment with it
* at the place where this function is executed, it still has an access to variables from the place where it was defined
   * e.g. o_x and o_y are carried with returned function, and are accessible when its called later, wherever and whenever is called
* variables are coming from the closure, from the environment where function was defined
```go
func makeDistOrigin(o_x, o_y float64) func (float64, float64) float64 {
   fn := func (x, y float64) float64 {
      return math.Sqrt(math.Pow(o_x, 2) + math.Pow(o_y, 2))
   }
   return fn
}
```

#### Variadic and Deferred
Variable Argument Number
* it is possible to pass variable number or arguments to function; such function is called variadic
* to specify this use ellipsis character: ...
* such argument is treated as a slice inside the function
```go
func getMax(vals ...int) int {
   maxV := -1
   for _, v := range vals {
      if v > maxV {
         maxV = v 
      }
   }
   return maxV
} 
```

#### How to pass list of arguments to variadic function?
* you can pass a comma-separated list of arguments
* you can pass a slice
* need a ... suffix
```  go
func main() {
   fmt.Println(getMax(1, 2, 6, 4))
   vslice := []int {1, 3, 6, 4}
   fmt.Println(getMax(vslice...))
}
```

#### Deferred Function Calls

* call can be deferred until surrounding function completes
* they don't get executed where they are explicitly called but after the surrounding function is done
* typically used for cleanup activities
* use keyword defer
```go
func main() {
   defer fmt.Println("Bye!") // "Bye!" printed after "Hello!"
   fmt.Println("Hello!")
}
```
* the arguments are NOT evaluated in a deferred way, they are evaluated immediately but the call is deferred
* if you pass an argument, it is evaluated right there where defer statement is 
```go
func main() {
   i := 1
   defer fmt.Println(i + 1) // 2 is printed second time
   i++ // 2
   fmt.Println(i) // 2 is printed first time
}
```

---

# Week 3 - OBJECT ORIENTATION IN GO
### Classes and Encapsulation
#### Classes 

* What is OOP?
   * Go supports OOP
   * It does not have classes but something equivalent (structs)
* What is class?
   * Collection of data fields and functions that share a well-defined responsibility (they are all related to the same concept)
* function in a class is called a method
* Example:
   * Point class
   * used in geometry program
   * data:
      * x and y coordinate
   * functions:
      * DistToOrigin(), Quadrant()
      * AddXOffset(), AddYOffset()
      * SetX(), SetY()
* class is a template; contains fields, not data

#### Object
* instance of the class
* contains data
* Example:
   * instances of Point class

#### Encapsulation
* associated with OOP (and generally, with abstraction)
* if there is a program using your class, you want to hide details
* you want to prevent someone changing internal data; therefore we provide public methods that shall be used to modify the state of the object from the outside
* Example:
   * double distance to origin (double x and y)
   * option 1 (safe): expose method DoubleDist() which doubles x and y internally
   * option 2 (not safe): allow programmer to access x and y directly; but programmer can make mistake if for example they double x but forget to double y
* by exposing methods we prevent such mistakes and object will always be in good state

#### Support for Classes in Go (Part 1)
* No "class" keyword
* there is no "class" keyword in Go
* most Object Oriented languages have class keyword
* Data fields and methods are defined inside a class block

example in Python:
```python
class Point:
   def _init_(self, xval, yval):
      self.x = xval
      self.y = yval
```

#### Associating Methods with Data

* Go has different way of associating methods with data
* Go is using "receiver types"
* data is some type 
* method has a receiver type that it is associated with
* BK: the approach is the same as in C: we have some struct type and if some function has to deal with it, we just pass to it a pointer to that struct. The * terminology reminds me of Objective-C
* type and function have to be defined in the same package
* when we call a method we use dot notation
* example:
   * we want to associate function Double with our custom type MyInt
   * MyInt is the receiver type - it is specified before the name of the function
   * mi is the receiver object (instance of the receiver type) that double would be called on
```go
type MyInt int

func (mi MyInt) Double () int { 
   return int(mi *2) 
}

func main() {
   v := MyInt(3)
   fmt.Println(v.Double())
}
```

* Double() could be defined for multiple types (to have multiple receiver types); Go looks what's the type left of the dot (.) operator to find out what is the receiver type (and so which Double() function to call)
* in the example above mi becomes an implicit argument of function Double() (just like this - pointer to the current instance of the class - is an implicit argument in C++ classes' methods)

#### Implicit Method Arguments

* although it seems that Double() takes no arguments, there is one implicit (hidden) argument: and instance (object) of the receiver type
* object v is an implicit argument to the method
* call by value (that's how argument passing is done in Go)
* a copy of v is made and passed to the function

#### Support for Classes in Go (Part 2)
* in a normal OOP language lots of different data (fields in class) is associated with any number of methods
* the same can be done in Go: we'll just make a struct with lots of various data and make it to be a receiver type
* in struct you can group together arbitrary number of arbitrary data 

#### Structs, again
   * struct types compose data fields 
   * this is traditional feature of classes

```go
type Point struct {
   x float64
   y float64
}
```

#### Structs with Methods

* structs and methods together allow arbitrary data and functions to be composed
```go
func (p Point) DistToOrig() {
   t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
   return math.Sqrt(t)
}

func main() {
   p1 := Point(3, 4)
   fmt.Println(p1.DistToOrig())
}
```

#### Encapsulation

* Controlling Access
   * Go provides lots of different support of encapsulation and how to keep data private
   * we want to be able to control data access
   * we want people to use data in a way we define - via functions/methods
   * we can define a set of public functions that allow another/external package to access the data

* package data

```go
var x int = 1

func PrintX() {
   fmt.Println(x)
}
```
-------------------
```go
package main

import "data"

func main() {
   data.PrintX()
}
```

* PrintX function starts with capital leter => it gets exported
* package main can access (see) x only through that exported function 
* x can't be modified externally in the example above but we can allow that if we export another method which allows that

* Controlling Access to Structs
   * we can do something similar to struct members
   * hide fields of structs by starting field name with lower-case letter
   * define public methods which access hidden data
   * example:
      * need InitMe() to assign hidden data fields

```go
package data

type Point struct {
   x float64
   y float64
}

func (p *Point) InitMe(xn, yn float64) {
   p.x = xn
   p.y = yn
}

func (p *Point) Scale (v float64) {
   p.x = p.x * v
   p.y = p.y * v
}

func (p *Point) PrintMe() {
   fmt.Println(p.x, p.y)
}
```
-----------------------------
```go
package main
import "data"

func main() {
   var p data.Point
   p.InitMe(3, 4)
   p.Scale(2)
   p.PrintMe()
}
```
* access to hidden fields only through public methods

##### Point Receivers
* Limitations of Methods
* there are some limitations to the process of associating methods with receiver types we described above
   * Method cannot change the state of receiver object as it's passed by value 
   * receiver type/object is implicitly passed to the method
   * it is passed by value (like any function argument in Go) => method receives only its copy => method can't change receiver object (method can't modify  the data inside the receiver)
   * example: OffsetX() should increase coordinate x in object p1
```go
func main() {
   p1 := Point(3, 4)
   p1.OffsetX(5)  // only temp copy of p1 inside OffsetX() is changed
}
```

* Large Receivers

   * if receiver object is large, lots of copying is required when you make a call
   * all object is copied onto the stack
```go
type image [100][100]int
func main(){
   i1 := GrabImage()
   i1.BlurImage() // 10000 bytes gets copied on stack - this can be slow
}
```

##### Solution:

* Pointer Receivers
   * instead of passing objects by value we can pass by reference (pointer)
   * instead of using regular types for receiver types, we can use pointer to those types as receiver type
```go
func (p *Point) OffsetX(v float64) {
   p.x = p.x + v
}
```

* Point Receivers, Referencing, Dereferencing
   * No Need to Dereference
   * when using a pointer receiver there is no need to perform explicit dereferencing (as in the previous example where Point is referenced as p, not *p)
   * dereferencing is automatic with . operator
   * No Need to Reference
```go
func main() {
   p := Point(3, 4)
   p.OffsetX(5)  // no need to do something like (&p).OffsetX()
   fmt.Println(p.x)
}
```

* Using Pointer Receivers
   * Good programming practice
   * all methods for a type have pointer receivers, or 
   * all methods for a type have non-pointer receivers
   * This is for mixing pointer/non-pointer receivers for a type will get confusing
   * pointer receiver allows modification

# Week 4 - INTERFACES FOR ABSTRACTION
#### Polymorphism
* one of OOP properties
* ability of an object to have different "forms" depending on the context
* example: 
   * Area() function - the function with the same name can do the same thing but in a different way, depending on the context
   * rectangle: area = base * height
   * triangle: area = 0.5 * base * height
* these two functions:
   * at high level of abstraction, they are identical in a way what they do; they do the same thing - compute the area
   * at low level of abstraction, they are different, in a way how do they compute the area
* We need Go's support for polymorphism to achieve this

* How is polymorphism implemented in traditional OOP languages?

#### Inheritance
* Go does NOT have inheritance
* there are parent-child (base-derived; superclass-subsclass) relations between classes
* superclass is a top level class
* subclass extends from superclass, subclass inherits data and methods from a superclass
* Example: 
   * Speaker superclass - represents anything that can make noise/speak
   * Speak() method prints "<noise>"
   * Subclasses Cat and Dog
   * Also have the Speak() method, inherited from the Speaker superclass
   * Cat and Dog are different forms of Speaker

#### Overriding
* subclass redefines a method inherited from the superclass 
* example: 
   * Speaker, Cat, Dog
   * Speaker Speak() prints "<noise>"
   * Cat Speak() prints "meow"
   * Dog Speak() prints "woof"
   * without overriding Cat and Dog Speak() methods would print "<noise>" but overriding allows them to redefine Speak() methods to print what they want
   * Speak() is polymorphic: in context of Cat it prints "meow" and in context of Dog, it prints "woof"
* although overriden in subclasses, the method keeps the same signature

#### Interfaces
* interface is a concept used in Go to help us get polymorphism
* we don't need inheritance or overriding to get polymorphism, we can get it with interfaces
* interface is a set of method signatures (name, parameters, return values)
* implementation is not defined
* it is used to express conceptual similarity between types
* example:
   * Shape2D interface has two methods: Area() and Perimeter()
   * all 2D shapes must have Area() and Perimeter()
   * any type that has these two methods can be considered to be 2D shape
* Satisfying the Interface
* types satisfies the interface if type defines all methods specified in the interface
* same method signatures (names, args, return values)
* example:
   * Rectangle and Triangle types satisfy the Shape2D interface if:
   * have Area() and Perimeter() methods
   * Additional methods are OK
   * similar to inheritance with overriding
* Defining an Interface Type
* use keyword interface

```go
type Shape2D interface {
   Area() float64
   Perimeter() float64
} 

type Triangle {
   ...
}

func (t Triangle) Area() float 64 {
   ...
}

func (t Triangle) Perimeter() float 64 {
   ...
}
```

* in the example above, we can say that Triangle implements (satisfies) interface Shape2D
* we don't state explicitly that Triangle implements interface; we just present to compiler the interface and methods and it infers which types can are satisfying which interfaces
* we don't care what data is in Triangle, which fields/properties it has; it only matters that that type is set as receiver type for functions with the same signature as defined in the interface
#### Interface vs. Concrete Types
#### Concrete vs Interface Types
* concrete and interface types are fundamentally different
##### Concrete Type
* a regular type
* Specifies the exact representation of the data and methods 
* fully specified
* complete method implementation included
* it has data which is associated with it

##### Interface type:
* just specifies some method signatures
* not data
* implementations are abstracted
* interface type eventually gets mapped to a concrete type

##### Interface Values
* can be treated like other values
* assigned to variables
* passed, returned
* interface values have two components

##### Dynamic Type: 
* concrete type which it is assigned to (like a class which implements an interface in classic OOP languages)
##### Dynamic Value: 
* value of the dynamic type (like an instance of the class which implements an interface in classic OOP languages)
* interface value is actually a pair (dynamic type, dynamic value)

##### Defining an Interface Type

```go
type Speaker interface {
   Speak()
}

type Dog struct {
   name string
}

func (d Dog) Speak() {
   fmt.Println(d.name)
}

func main() {
   var s1 Speaker // interface value
   var d1 Dog{"Brian"} // 
   s1 = d1 // legal as Dog satisfies interface Speaker; 
   s1.Speak()
}
```

* interface type is Speaker
* interface value is s1
* dynamic type is Dog
* dynamic value is d1
* pair is (Dog, d1)

##### Interface with Nil Dynamic Value
* an interface can have a nil dynamic value (no dynamic value)

```go
var s1 Speaker
var d1 *Dog // pointer
s1 = d1 // legal
```

* d1 is pointer to Dog, it is not concrete object, has no data in it; d1 has no concrete value yet
* s1 has dynamic type - Dog, but has NO dynamic value

#### Nil Dynamic Value
* interface with dynamic type but not dynamic value; it is legal to call interface methods on nil dynamic value
* can still call the Speak() method of s1
* doesn't need a dynamic value to call interface methods
* need to check inside the method

```go
func(d *Dog) Speak() {
   if d == nil { // does it have dynamic value or not?
      fmt.Println("<noise>")
   } else {
      fmt.Println(d.name)
   }
}

var s1 Speaker
var d1 *Dog
s1 = d1
s1.Speak() // it is legal to call function on a non-assigned pointer!
```

#### Nil Interface Value
* interface with nil dynamic type
* very different from an interface with a nil dynamic value
* we can't call interface methods as there is no underlying concrete type with methods to call, there is no method implementations
* Nil dynamic value and valid dynamic type
* can call a method since type is known
example:
```go
var s1 Speaker
var d1 *Dog
s1 = d1
s1.Speak()
Nil dynamic type
cannot call a method, runtime error
var s1 Speaker
s1.Speak() // error!
```

#### Using Interfaces
* why would we need interfaces? why are they used?
* Ways to Use an Interface
* need a function which takes multiple types of parameter
* e.g. 
   * we want function to take either type int or type float
   * we want function foo() to accept parameter of type X or type Y
* we can:
   * define interface Z
   * make types X and Y satisfy Z (BK: this is something like class extensions...as any type can at any time be extended to satisfy any interface; this can happen after some concrete type is defined)
   * make Z to be the type of the foo() argument
* interface methods must be those needed (called) by foo()

* EXAMPLE (Pool in a Yard)
   * I need to put a pool in my yard
   * Pool needs to fit in my yard
   * total area must be limited
   * Pool needs to be fenced
   * total perimeters must be limited
   * Need to determine if a pool shape satisfies criteria
      * FitInYard() bool
   * takes a shape as argument
   * returns true if the shape satisfies criteria
      * FitInYard()
   * many possible shape types
   * rectangle, triangle, circle, etc...
   * FitInYard() should take many shape types
   * Valid shape types must have:
      * Area()
      * Perimeter()
   * Any shape with these methods is OK

##### Interface for Shapes

```go
type Shape2D interface {
   Area() float64
   Perimeter float64
}

type Triangle {...}
func (t Triangle) Area() float64 {...}
func (t Triangle) Perimeter() float64 {...}

type Rectangle {...}
func (r Rectangle) Area() float64 {...}
func (r Rectangle) Perimeter() float64 {...}

Rectangle and Triangle satisfy interface Shape2D
func FitInYard(s Shape2D) bool {
   if (s.Area() < 100 && s.Perimeter() < 100) {
      return true}
   else {
      return false
   }
}

```

* parameter is any type that satisfies Shape2D interface

##### Empty Interface
* specifies no methods
* all types satisfy the empty interface
* use it to have a function accept any type as a parameter
* use interface{} to specify it
* val can be any type
```go
func PrintMe(val interface{}) {
   fmt.Println(val)
}
```

#### Type Assertions
##### Concealing Type Differences
* a lot of point of interfaces is to hide differences (or highlight similarities) between types
* example:
   * Triangles and Rectangles are treated in the same way in FitInYard() - as long as they satisfy Shape2D interface
   * different types which has some similarities are treated in the same way
   * sometimes you need to treat different types in different ways
   * sometimes we need to differentiate based on the type, to figure out what is the concrete type
   * in FitInYard() it does not matter what is the concrete type; it can be Rectangle or Triangle
* Exposing Type Differences
* example:
   * Graphics program
   * DrawShape() will draw any shape, it can take any Shape2D as an argument
   * func DrawShape(s Shape2D) { ...
   * Underlying API has different drawing functions for each shape so they have to take particular, specific, concrete types as arguments:
   * func DrawRect (r Rectangle) { ...
   * func DrawTriangle (t Triangle) { ...
   * Inside DrawShape() we need to find out what is the concrete type of s so we know which underlying function to call; concrete type must be determined
   * type assertions are used for that
   * Type Assertions for Disambiguation
   * type assertions can be used to determine and extract the underlying concrete type 

```go
func DrawShape(s Shape2D) bool {
   rect, ok := s.(Rectangle)
   if ok {
      DrawRect(rect)
   }
   tri, ok := s.(Triangle)
   if ok {
      DrawTriangle(tri)
   }
}
```

* type assertions extract Rectangle from Shape2D
* concrete type has to be specified in parentheses
* If interface contains specified concrete type 
* rect == specified concrete type
* ok == true
* If interface does not contain concrete type 
* rect == zero value for that type
* ok == false

#### Type Switch 

Interface can be satisfied by many concrete types; we might be interested only in some of them
Another way to do this disambiguation is to use switch
switch statement used with a type assertion
use keyword type in parentheses: .(type)
if s is Triangle then sh will be Triangle
```go
func DrawShape(s Shape2D) bool {
   switch := sh := s.(type) {
   case Rectangle:
      DrawRectangle(sh)
   case Triangle:
      DrawTriangle(sh)
   }
}
```

#### Error Handling
* common use of Error interface in Go
##### Error Interface
* there are lot of Go functions built in packages where return two values: 
   * value
   * error (error interface objects which indicates an error)

```go
type error interface {
   Error() string // prints the error message
}
```

* correct / successful operation: error == nil
* incorrect / failed operation: error != nil
* if Go function returns an error (usually as a second value) you should check that error and handle it! (BK: compiler will complain if returned error is not checked explicitly in the code)

```go
f, err := os.Open("harris/test.txt")
if err != nil {
   fmt.Println(err)
   return
}
```
* check whether error is nil
* if it's not nil, handle it
