package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Lesson represents a tutorial lesson
type Lesson struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Exercise    string `json:"exercise"`
	Solution    string `json:"solution"`
	Difficulty  string `json:"difficulty"`
	Order       int    `json:"order"`
	Category    string `json:"category"`
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
			Exercise: `Write a program that prints "Hello, World!" to the console.

Hint: Use fmt.Println() to print text.`,
			Solution: `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`,
			Difficulty:  "beginner",
			Order:       1,
			Category:    "basics",
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
			Difficulty:  "beginner",
			Order:       2,
			Category:    "basics",
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
			Difficulty:  "intermediate",
			Order:       3,
			Category:    "functions",
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
			Difficulty:  "beginner",
			Order:       4,
			Category:    "control-flow",
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
			Difficulty:  "beginner",
			Order:       5,
			Category:    "control-flow",
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
			Difficulty:  "intermediate",
			Order:       6,
			Category:    "data-structures",
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
			Difficulty:  "intermediate",
			Order:       7,
			Category:    "data-structures",
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
			Difficulty:  "intermediate",
			Order:       8,
			Category:    "data-structures",
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
			Difficulty:  "intermediate",
			Order:       9,
			Category:    "methods",
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
			Difficulty:  "advanced",
			Order:       10,
			Category:    "interfaces",
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
