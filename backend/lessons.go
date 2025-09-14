package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Lesson represents a tutorial lesson
type Lesson struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Explanation string   `json:"explanation"`
	Variants    []string `json:"variants"`
	Exercise    string   `json:"exercise"`
	Solution    string   `json:"solution"`
	Difficulty  string   `json:"difficulty"`
	Order       int      `json:"order"`
	Category    string   `json:"category"`
}

// Get comprehensive Go tutorial lessons
func getTutorialLessons() []Lesson {
	return []Lesson{
		{
			ID:          1,
			Title:       "Hello, Go!",
			Description: "Write your first Go program and understand the basics",
			Content: `Welcome to Go! In this lesson, you'll learn:

• How to write a basic Go program
• Understanding the main function
• Using the fmt package for output
• Go's package system

Go programs start with a package declaration and have a main function as the entry point.`,
			Explanation: `Let's break down your first Go program:

1. **Package Declaration**: Every Go file starts with a package declaration. The 'main' package is special - it tells Go this is an executable program, not a library.

2. **Import Statement**: The 'import "fmt"' line brings in Go's formatting package, which provides functions for printing text and formatting output.

3. **Main Function**: The main() function is the entry point of every Go program. When you run a Go program, execution starts here.

4. **Function Call**: fmt.Println() is a function that prints text to the console and adds a newline at the end.

**Key Go Concepts:**
- Go is compiled, not interpreted
- Every Go program must have a main package and main function
- Go uses explicit imports - you must import what you use
- Semicolons are optional in Go (the compiler adds them automatically)`,
			Variants: []string{
				`package main

import "fmt"

func main() {
    fmt.Print("Hello, ")  // Print without newline
    fmt.Print("World!")
    fmt.Println()         // Add newline
}`,
				`package main

import "fmt"

func main() {
    name := "Go Developer"
    fmt.Printf("Hello, %s!\n", name)  // Formatted printing
}`,
				`package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Welcome to Go programming!")
    fmt.Println("Let's learn together!")
}`,
			},
			Exercise: `Write a program that prints "Hello, World!" to the console.

Hint: Use fmt.Println() to print text.`,
			Solution: `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`,
			Difficulty: "beginner",
			Order:      1,
			Category:   "basics",
		},
		{
			ID:          2,
			Title:       "Variables and Types",
			Description: "Learn about Go's type system and variable declarations",
			Content: `Go is statically typed, meaning variables have a specific type that cannot change.

Key concepts:
• Variable declaration with var keyword
• Type inference with := operator
• Basic types: int, float64, string, bool
• Constants with const keyword`,
			Explanation: `Go's type system is designed for safety and clarity. Here's what you need to know:

**Variable Declaration Methods:**
1. **Explicit Declaration**: var name string = "value" - declares variable with explicit type
2. **Type Inference**: name := "value" - Go infers the type from the value
3. **Zero Values**: var name string - declares variable with zero value (empty string for strings)

**Basic Types in Go:**
- int - integers (32 or 64 bits depending on platform)
- float64 - floating-point numbers (64-bit)
- string - text strings
- bool - true/false values
- byte - alias for uint8 (0-255)

**Key Differences from Other Languages:**
- Go doesn't allow implicit type conversions
- Variables must be used (unused variables cause compilation errors)
- Zero values are meaningful (0 for numbers, "" for strings, false for bools)
- Constants are immutable and must be known at compile time`,
			Variants: []string{
				`package main

import "fmt"

func main() {
    // Different ways to declare variables
    var name string = "Alice"
    var age int = 25
    var isStudent bool = true
    
    fmt.Printf("Name: %s, Age: %d, Student: %t\n", name, age, isStudent)
}`,
				`package main

import "fmt"

func main() {
    // Using type inference
    name := "Bob"
    age := 30
    salary := 50000.50
    isEmployed := true
    
    fmt.Printf("%s is %d years old, earns $%.2f, employed: %t\n", 
               name, age, salary, isEmployed)
}`,
				`package main

import "fmt"

func main() {
    // Constants and multiple declarations
    const pi = 3.14159
    const company = "TechCorp"
    
    var (
        firstName = "Charlie"
        lastName  = "Brown"
        id        = 12345
    )
    
    fmt.Printf("Employee: %s %s (ID: %d) at %s\n", 
               firstName, lastName, id, company)
}`,
			},
			Exercise: `Create variables of different types and print them:
- A string variable with your name
- An integer variable with your age
- A boolean variable set to true
- A float variable with a decimal number`,
			Solution: `package main

import "fmt"

func main() {
    var name string = "Go Developer"
    age := 25
    isLearning := true
    var score float64 = 95.5
    
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Learning:", isLearning)
    fmt.Println("Score:", score)
}`,
			Difficulty: "beginner",
			Order:      2,
			Category:   "basics",
		},
		{
			ID:          3,
			Title:       "Functions",
			Description: "Create and use functions in Go",
			Content: `Functions are the building blocks of Go programs.

Key concepts:
• Function declaration syntax
• Parameters and return types
• Multiple return values
• Named return values
• Function calls`,
			Explanation: `Functions in Go are powerful and flexible. Here's what makes them special:

**Function Declaration Syntax:**
func functionName(parameters) returnType { ... }

**Key Features:**
1. **Multiple Return Values**: Go functions can return multiple values, commonly used for returning a result and an error
2. **Named Return Values**: You can name return values for clarity and documentation
3. **Variadic Functions**: Functions can accept a variable number of arguments using ...
4. **Function Values**: Functions are first-class citizens - you can assign them to variables

**Common Patterns:**
- Error handling: func doSomething() (result, error)
- Multiple results: func divide(a, b int) (int, int) (quotient, remainder)
- Named returns: func calculate() (sum, product int)

**Best Practices:**
- Keep functions small and focused
- Use descriptive names
- Return errors as the last return value
- Prefer multiple return values over complex structs`,
			Variants: []string{
				`package main

import "fmt"

// Function with multiple return values
func divide(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := divide(17, 5)
    fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)
}`,
				`package main

import "fmt"

// Function with named return values
func calculate(x, y int) (sum, product int) {
    sum = x + y
    product = x * y
    return // naked return
}

func main() {
    s, p := calculate(4, 5)
    fmt.Printf("Sum: %d, Product: %d\n", s, p)
}`,
				`package main

import "fmt"

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println("Sum of 1,2,3:", sum(1, 2, 3))
    fmt.Println("Sum of 1,2,3,4,5:", sum(1, 2, 3, 4, 5))
}`,
			},
			Exercise: `Create a function called 'add' that takes two integers and returns their sum.
Then call this function with the numbers 15 and 27 and print the result.`,
			Solution: `package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    result := add(15, 27)
    fmt.Println("Sum:", result)
}`,
			Difficulty: "intermediate",
			Order:      3,
			Category:   "functions",
		},
		{
			ID:          4,
			Title:       "Control Flow - If/Else",
			Description: "Make decisions in your Go programs",
			Content: `Control flow allows your program to make decisions and execute different code paths.

Key concepts:
• if statements
• else and else if
• Comparison operators (==, !=, <, >, <=, >=)
• Logical operators (&&, ||, !)
• Short if statement syntax`,
			Explanation: `Go's control flow is clean and expressive. Here's what you need to know:

**If Statement Syntax:**
if condition { ... }

**Key Features:**
1. **No Parentheses**: Unlike C/Java, Go doesn't require parentheses around conditions
2. **Short If**: You can declare variables in the if condition: if x := getValue(); x > 0 { ... }
3. **Comparison Operators**: ==, !=, <, >, <=, >= work as expected
4. **Logical Operators**: && (AND), || (OR), ! (NOT)

**Best Practices:**
- Use early returns to reduce nesting
- Prefer explicit conditions over complex boolean expressions
- Use short if statements for initialization
- Keep conditions simple and readable

**Common Patterns:**
- Error checking: if err != nil { ... }
- Range checking: if index >= 0 && index < len(slice) { ... }
- Type assertions: if str, ok := value.(string); ok { ... }`,
			Variants: []string{
				`package main

import "fmt"

func main() {
    score := 85
    
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: F")
    }
}`,
				`package main

import "fmt"

func main() {
    // Short if statement
    if age := 18; age >= 18 {
        fmt.Println("You are an adult")
    } else {
        fmt.Println("You are a minor")
    }
    
    // Multiple conditions
    temperature := 25
    isSunny := true
    
    if temperature > 20 && isSunny {
        fmt.Println("Perfect weather for a walk!")
    }
}`,
				`package main

import "fmt"

func main() {
    // Complex conditions
    username := "admin"
    password := "secret123"
    
    if username == "admin" && password == "secret123" {
        fmt.Println("Access granted")
    } else if username == "admin" {
        fmt.Println("Wrong password")
    } else {
        fmt.Println("Invalid username")
    }
}`,
			},
			Exercise: `Write a program that checks if a number is positive, negative, or zero.
Use variables for the number and print the appropriate message.`,
			Solution: `package main

import "fmt"

func main() {
    number := -5
    
    if number > 0 {
        fmt.Println("The number is positive")
    } else if number < 0 {
        fmt.Println("The number is negative")
    } else {
        fmt.Println("The number is zero")
    }
}`,
			Difficulty: "beginner",
			Order:      4,
			Category:   "control-flow",
		},
		{
			ID:          5,
			Title:       "Loops",
			Description: "Repeat code execution with loops",
			Content: `Go has only one loop construct: the for loop, but it's very flexible.

Key concepts:
• Basic for loop syntax
• Range-based loops
• Infinite loops
• Break and continue statements
• Loop variations`,
			Explanation: `Go's for loop is incredibly versatile. Here's what makes it special:

**For Loop Variations:**
1. **Traditional**: for i := 0; i < 10; i++ { ... }
2. **While-style**: for condition { ... }
3. **Infinite**: for { ... }
4. **Range**: for index, value := range slice { ... }

**Key Features:**
- Only one loop construct (for) but very flexible
- Range loops work with slices, maps, strings, and channels
- Break exits the loop immediately
- Continue skips to the next iteration
- Labels allow breaking/continuing outer loops

**Range Loop Details:**
- For slices: for i, v := range slice { ... }
- For maps: for key, value := range map { ... }
- For strings: for i, char := range "hello" { ... }
- Ignore index: for _, value := range slice { ... }

**Best Practices:**
- Use range loops when possible (more readable)
- Use break/continue sparingly
- Consider using labels for nested loops
- Prefer explicit conditions over infinite loops`,
			Variants: []string{
				`package main

import "fmt"

func main() {
    // Traditional for loop
    fmt.Println("Counting up:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
    
    // Counting down
    fmt.Println("Counting down:")
    for i := 5; i >= 1; i-- {
        fmt.Println(i)
    }
}`,
				`package main

import "fmt"

func main() {
    // Range loop with slice
    fruits := []string{"apple", "banana", "orange"}
    
    fmt.Println("Fruits:")
    for i, fruit := range fruits {
        fmt.Printf("%d: %s\n", i, fruit)
    }
    
    // Range loop with map
    scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}
    
    fmt.Println("Scores:")
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }
}`,
				`package main

import "fmt"

func main() {
    // While-style loop
    count := 0
    for count < 3 {
        fmt.Printf("Count: %d\n", count)
        count++
    }
    
    // Loop with break
    fmt.Println("Finding first even number:")
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Printf("First even number: %d\n", i)
            break
        }
    }
}`,
			},
			Exercise: `Write a program that prints numbers from 1 to 10, but skip the number 5.
Use a for loop and continue statement.`,
			Solution: `package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            continue
        }
        fmt.Println(i)
    }
}`,
			Difficulty: "beginner",
			Order:      5,
			Category:   "control-flow",
		},
		{
			ID:          6,
			Title:       "Arrays and Slices",
			Description: "Work with collections of data",
			Content: `Arrays and slices are fundamental data structures in Go.

Key concepts:
• Arrays: fixed-size collections
• Slices: dynamic arrays
• Slice operations (append, copy, len, cap)
• Range over slices
• Slice literals`,
			Explanation: `Understanding arrays and slices is crucial for Go programming:

**Arrays:**
- Fixed-size collections: var arr [5]int
- Size is part of the type: [5]int and [10]int are different types
- Zero-initialized by default
- Passed by value (copied)

**Slices:**
- Dynamic arrays built on top of arrays
- Reference type (passed by reference)
- Have length (len) and capacity (cap)
- Can grow using append()

**Key Operations:**
- len(slice) - get length
- cap(slice) - get capacity
- append(slice, elements...) - add elements
- copy(dst, src) - copy elements
- slice[start:end] - create sub-slice

**Slice Internals:**
- A slice is a struct with pointer, length, and capacity
- Multiple slices can share the same underlying array
- Modifying a slice affects all slices sharing the same array

**Best Practices:**
- Prefer slices over arrays
- Use make() to create slices with specific capacity
- Be careful with slice sharing
- Use copy() when you need independent slices`,
			Variants: []string{
				`package main

import "fmt"

func main() {
    // Array declaration and initialization
    var numbers [5]int
    numbers[0] = 10
    numbers[1] = 20
    
    fmt.Println("Array:", numbers)
    
    // Array literal
    colors := [3]string{"red", "green", "blue"}
    fmt.Println("Colors:", colors)
    
    // Array with inferred size
    fruits := [...]string{"apple", "banana"}
    fmt.Println("Fruits:", fruits)
}`,
				`package main

import "fmt"

func main() {
    // Slice creation
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", 
               numbers, len(numbers), cap(numbers))
    
    // Adding elements
    numbers = append(numbers, 6, 7, 8)
    fmt.Printf("After append: %v, Length: %d, Capacity: %d\n", 
               numbers, len(numbers), cap(numbers))
    
    // Sub-slice
    subSlice := numbers[2:5]
    fmt.Println("Sub-slice:", subSlice)
}`,
				`package main

import "fmt"

func main() {
    // Creating slices with make
    slice1 := make([]int, 5)        // length 5, capacity 5
    slice2 := make([]int, 3, 10)     // length 3, capacity 10
    
    fmt.Printf("Slice1: %v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
    fmt.Printf("Slice2: %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))
    
    // Copying slices
    slice3 := make([]int, len(slice1))
    copy(slice3, slice1)
    fmt.Println("Copied slice:", slice3)
}`,
			},
			Exercise: `Create a slice of strings with your favorite programming languages.
Then add "Go" to the slice and print all languages.`,
			Solution: `package main

import "fmt"

func main() {
    languages := []string{"Python", "JavaScript", "Java"}
    languages = append(languages, "Go")
    
    fmt.Println("My favorite languages:")
    for i, lang := range languages {
        fmt.Printf("%d. %s\n", i+1, lang)
    }
}`,
			Difficulty: "intermediate",
			Order:      6,
			Category:   "data-structures",
		},
		{
			ID:          7,
			Title:       "Maps",
			Description: "Store key-value pairs with maps",
			Content: `Maps are Go's built-in associative data type (like dictionaries in Python).

Key concepts:
• Map declaration and initialization
• Adding and accessing elements
• Checking if a key exists
• Deleting elements
• Iterating over maps`,
			Explanation: `Maps are Go's implementation of hash tables. Here's what you need to know:

**Map Creation:**
- make(map[keyType]valueType) - creates empty map
- map[keyType]valueType{key: value} - map literal
- var m map[string]int - zero value is nil

**Key Features:**
- Keys must be comparable (==, !=)
- Values can be any type
- Maps are reference types
- Zero value is nil (not empty map)

**Operations:**
- m[key] = value - add/update
- value := m[key] - access
- value, ok := m[key] - check existence
- delete(m, key) - remove
- len(m) - get size

**Important Notes:**
- Accessing non-existent key returns zero value
- Use comma ok idiom to check if key exists
- Maps are not safe for concurrent access
- Iteration order is not guaranteed

**Best Practices:**
- Always check if key exists before using
- Use make() for empty maps
- Consider sync.Map for concurrent access
- Use descriptive key types`,
			Variants: []string{
				`package main

import "fmt"

func main() {
    // Creating maps
    ages := make(map[string]int)
    ages["Alice"] = 30
    ages["Bob"] = 25
    
    fmt.Println("Ages:", ages)
    
    // Map literal
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    fmt.Println("Colors:", colors)
}`,
				`package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice":   95,
        "Bob":     87,
        "Charlie": 92,
    }
    
    // Check if key exists
    if score, exists := scores["Alice"]; exists {
        fmt.Printf("Alice's score: %d\n", score)
    }
    
    // Check non-existent key
    if score, exists := scores["David"]; exists {
        fmt.Printf("David's score: %d\n", score)
    } else {
        fmt.Println("David not found")
    }
}`,
				`package main

import "fmt"

func main() {
    inventory := map[string]int{
        "apples":  10,
        "bananas": 5,
        "oranges": 8,
    }
    
    // Iterate over map
    fmt.Println("Inventory:")
    for item, quantity := range inventory {
        fmt.Printf("%s: %d\n", item, quantity)
    }
    
    // Delete an item
    delete(inventory, "bananas")
    fmt.Println("After deleting bananas:", inventory)
}`,
			},
			Exercise: `Create a map that stores student names as keys and their grades as values.
Add at least 3 students, then print all students and their grades.`,
			Solution: `package main

import "fmt"

func main() {
    grades := make(map[string]int)
    grades["Alice"] = 95
    grades["Bob"] = 87
    grades["Charlie"] = 92
    
    fmt.Println("Student Grades:")
    for name, grade := range grades {
        fmt.Printf("%s: %d\n", name, grade)
    }
}`,
			Difficulty: "intermediate",
			Order:      7,
			Category:   "data-structures",
		},
		{
			ID:          8,
			Title:       "Structs",
			Description: "Create custom data types with structs",
			Content: `Structs allow you to group related data together.

Key concepts:
• Struct definition
• Creating struct instances
• Accessing struct fields
• Struct literals
• Anonymous structs`,
			Explanation: `Structs are Go's way of creating custom data types. Here's what makes them powerful:

**Struct Definition:**
type StructName struct { field1 type1; field2 type2 }

**Key Features:**
- Group related data together
- Fields can be any type (including other structs)
- Fields can be exported (capitalized) or unexported
- Zero value has all fields set to their zero values

**Creating Instances:**
- var p Person - zero value
- p := Person{Name: "Alice", Age: 30} - struct literal
- p := Person{"Alice", 30} - positional literal
- p := &Person{Name: "Bob"} - pointer to struct

**Field Access:**
- p.Name - direct access
- (*p).Name or p.Name - pointer access (automatic dereferencing)

**Anonymous Structs:**
- struct{name string; age int}{"Alice", 30}
- Useful for one-off data structures

**Best Practices:**
- Use descriptive field names
- Group related fields together
- Consider field ordering for memory layout
- Use exported fields for public APIs`,
			Variants: []string{
				`package main

import "fmt"

type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Different ways to create structs
    person1 := Person{Name: "Alice", Age: 30, City: "New York"}
    person2 := Person{"Bob", 25, "London"}  // positional
    person3 := Person{Name: "Charlie"}      // partial
    
    fmt.Printf("Person 1: %+v\n", person1)
    fmt.Printf("Person 2: %+v\n", person2)
    fmt.Printf("Person 3: %+v\n", person3)
}`,
				`package main

import "fmt"

type Address struct {
    Street string
    City   string
    Zip    string
}

type Employee struct {
    Name    string
    Age     int
    Address Address  // embedded struct
}

func main() {
    emp := Employee{
        Name: "John Doe",
        Age:  35,
        Address: Address{
            Street: "123 Main St",
            City:   "Boston",
            Zip:    "02101",
        },
    }
    
    fmt.Printf("Employee: %s, %d years old\n", emp.Name, emp.Age)
    fmt.Printf("Address: %s, %s %s\n", emp.Address.Street, emp.Address.City, emp.Address.Zip)
}`,
				`package main

import "fmt"

func main() {
    // Anonymous struct
    person := struct {
        name string
        age  int
    }{"Alice", 30}
    
    fmt.Printf("Anonymous struct: %s, %d\n", person.name, person.age)
    
    // Anonymous struct with pointer
    ptr := &struct {
        x, y int
    }{10, 20}
    
    fmt.Printf("Pointer to anonymous struct: x=%d, y=%d\n", ptr.x, ptr.y)
}`,
			},
			Exercise: `Create a struct called 'Person' with fields for name (string) and age (int).
Create two Person instances and print their information.`,
			Solution: `package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person1 := Person{Name: "Alice", Age: 30}
    person2 := Person{Name: "Bob", Age: 25}
    
    fmt.Printf("Person 1: %s, %d years old\n", person1.Name, person1.Age)
    fmt.Printf("Person 2: %s, %d years old\n", person2.Name, person2.Age)
}`,
			Difficulty: "intermediate",
			Order:      8,
			Category:   "data-structures",
		},
		{
			ID:          9,
			Title:       "Methods",
			Description: "Add behavior to your structs with methods",
			Content: `Methods are functions that belong to a specific type.

Key concepts:
• Method syntax
• Value receivers vs pointer receivers
• Method calls
• Method chaining
• Interface methods`,
			Explanation: `Methods add behavior to types in Go. Here's what you need to know:

**Method Syntax:**
func (receiver Type) methodName(parameters) returnType { ... }

**Receiver Types:**
1. **Value Receiver**: func (p Person) methodName()
   - Receives a copy of the value
   - Cannot modify the original
   - Use for small types or when you don't need to modify

2. **Pointer Receiver**: func (p *Person) methodName()
   - Receives a pointer to the value
   - Can modify the original
   - Use for large types or when you need to modify

**Key Features:**
- Methods can be defined on any type (not just structs)
- Go automatically handles pointer/value conversion
- Methods can be chained
- Methods can satisfy interfaces

**Best Practices:**
- Use pointer receivers for structs (consistency)
- Use value receivers for small, immutable types
- Keep methods focused and small
- Use descriptive method names
- Consider method chaining for fluent APIs

**Common Patterns:**
- Getters/Setters: func (p *Person) GetName() string
- Validators: func (p *Person) IsValid() bool
- Builders: func (p *Person) WithName(name string) *Person`,
			Variants: []string{
				`package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Value receiver method
func (p Person) Greet() {
    fmt.Printf("Hello, I'm %s and I'm %d years old!\n", p.Name, p.Age)
}

// Pointer receiver method
func (p *Person) HaveBirthday() {
    p.Age++
    fmt.Printf("%s is now %d years old!\n", p.Name, p.Age)
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    person.Greet()
    person.HaveBirthday()
    person.Greet()
}`,
				`package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// Method chaining
func (r *Rectangle) SetWidth(w float64) *Rectangle {
    r.Width = w
    return r
}

func (r *Rectangle) SetHeight(h float64) *Rectangle {
    r.Height = h
    return r
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{}
    area := rect.SetWidth(5.0).SetHeight(3.0).Area()
    fmt.Printf("Rectangle area: %.2f\n", area)
}`,
				`package main

import "fmt"

// Method on non-struct type
type MyInt int

func (m MyInt) IsEven() bool {
    return int(m)%2 == 0
}

func (m MyInt) Double() MyInt {
    return m * 2
}

func main() {
    num := MyInt(4)
    fmt.Printf("%d is even: %t\n", num, num.IsEven())
    fmt.Printf("Double of %d is %d\n", num, num.Double())
}`,
			},
			Exercise: `Add a method called 'Greet' to the Person struct that prints a greeting.
Create a Person instance and call the Greet method.`,
			Solution: `package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() {
    fmt.Printf("Hello, I'm %s and I'm %d years old!\n", p.Name, p.Age)
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    person.Greet()
}`,
			Difficulty: "intermediate",
			Order:      9,
			Category:   "methods",
		},
		{
			ID:          10,
			Title:       "Interfaces",
			Description: "Define behavior contracts with interfaces",
			Content: `Interfaces define a set of methods that a type must implement.

Key concepts:
• Interface declaration
• Implicit interface implementation
• Interface values
• Empty interface
• Type assertions`,
			Explanation: `Interfaces are Go's way of achieving polymorphism. Here's what makes them powerful:

**Interface Declaration:**
type InterfaceName interface { method1() returnType; method2() returnType }

**Key Features:**
- **Implicit Implementation**: Types implement interfaces automatically if they have the required methods
- **Interface Values**: Can hold any value that implements the interface
- **Empty Interface**: interface{} can hold any type
- **Type Assertions**: Extract concrete types from interface values

**Interface Values:**
- Have a type and a value
- Can be nil
- Support type assertions and type switches

**Common Patterns:**
- **Reader/Writer**: io.Reader, io.Writer
- **Error Handling**: error interface
- **String Representation**: fmt.Stringer
- **Sorting**: sort.Interface

**Type Assertions:**
- value, ok := interfaceValue.(ConcreteType)
- value := interfaceValue.(ConcreteType) (panics if wrong type)

**Best Practices:**
- Keep interfaces small (1-3 methods)
- Use descriptive names
- Prefer composition over inheritance
- Use interfaces for abstraction
- Consider interface{} sparingly`,
			Variants: []string{
				`package main

import "fmt"

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

func main() {
    shapes := []Shape{
        Rectangle{Width: 5, Height: 3},
        Circle{Radius: 4},
    }
    
    for _, shape := range shapes {
        fmt.Printf("Area: %.2f, Perimeter: %.2f\n", 
                   shape.Area(), shape.Perimeter())
    }
}`,
				`package main

import "fmt"

// Empty interface
func printAny(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

func main() {
    printAny(42)
    printAny("Hello")
    printAny(true)
    printAny(3.14)
}`,
				`package main

import "fmt"

type Writer interface {
    Write(data string)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) {
    fmt.Println("Console:", data)
}

type FileWriter struct{
    filename string
}

func (fw FileWriter) Write(data string) {
    fmt.Printf("File [%s]: %s\n", fw.filename, data)
}

func logMessage(w Writer, message string) {
    w.Write(message)
}

func main() {
    console := ConsoleWriter{}
    file := FileWriter{filename: "app.log"}
    
    logMessage(console, "Application started")
    logMessage(file, "Application started")
}`,
			},
			Exercise: `Create an interface called 'Shape' with a method 'Area()' that returns a float64.
Create a struct 'Rectangle' that implements this interface.`,
			Solution: `package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 5.0, Height: 3.0}
    fmt.Printf("Rectangle area: %.2f\n", rect.Area())
}`,
			Difficulty: "advanced",
			Order:      10,
			Category:   "interfaces",
		},
	}
}

// getLessons returns all available lessons
func getLessons(c *gin.Context) {
	lessons := getTutorialLessons()
	c.JSON(http.StatusOK, lessons)
}

// getLesson returns a specific lesson by ID
func getLesson(c *gin.Context) {
	lessonID := c.Param("id")
	lessons := getTutorialLessons()

	for _, lesson := range lessons {
		if fmt.Sprintf("%d", lesson.ID) == lessonID {
			c.JSON(http.StatusOK, lesson)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
}
