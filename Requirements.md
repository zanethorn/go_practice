🧱 1. Warm‑up: Types, Functions, and Control Flow
These build muscle memory for Go’s syntax and standard library.

Exercise A — FizzBuzz with a Twist
Write a function FizzBuzz(n int) []string that:

Returns a slice of strings from 1 to n

Replaces multiples of 3 with "Fizz", 5 with "Buzz", and 15 with "FizzBuzz"

Uses a switch statement

Includes a table‑driven test in _test.go

Exercise B — Temperature Converter
Create a small CLI that:

Accepts a temperature and unit (C or F)

Converts it

Uses fmt.Scanln or flag for input

Returns errors for invalid units

📦 2. Collections: Slices, Maps, Structs
These teach you how Go handles memory, references, and iteration.

Exercise C — Unique Words Counter
Given a string, return a map of word → count.
Requirements:

Normalize to lowercase

Strip punctuation

Sort the output by frequency before printing

Write a benchmark test

Exercise D — Simple Inventory System
Define a struct:

go
type Item struct {
Name  string
Count int
}
Then:

Store items in a slice

Add functions to add, remove, and list items

Return errors when removing nonexistent items

Add a method func (i *Item) Restock(n int)

⚙️ 3. Interfaces and Composition
Go’s interfaces are implicit—these exercises force you to design around behavior.

Exercise E — Shape Interface
Define:

go
type Shape interface {
Area() float64
}
Implement Circle, Rectangle, and Triangle.
Then write a function TotalArea(shapes ...Shape) float64.

Exercise F — Logger Interface
Create a Logger interface with Info, Warn, and Error.
Implement:

A console logger

A file logger

A “null” logger for tests

🔄 4. Concurrency: Goroutines and Channels
These are the heart of Go. Keep them small but real.

Exercise G — Worker Pool
Build a worker pool that:

Accepts integers on a jobs channel

Squares them

Sends results on a results channel

Uses sync.WaitGroup

Allows configurable worker count

Exercise H — Rate‑Limited Downloader
Write a function that:

Accepts a list of URLs

Downloads them concurrently

Limits concurrency to N workers

Times out after 3 seconds per request using context.Context

🌐 5. Mini‑Project: Log Analyzer CLI
This ties everything together.

Build a CLI tool that:

Reads a log file

Counts occurrences of log levels (INFO, WARN, ERROR)

Streams lines through a channel

Uses goroutines to process chunks

Outputs a JSON summary

Includes unit tests and benchmarks

This is small enough to finish in an afternoon but deep enough to teach real Go patterns.

🧭 6. Optional Stretch Tasks (If You Want More Challenge)
Implement your own io.Reader that transforms text (ROT13, uppercase, etc.)

Build a tiny HTTP server with one GET endpoint and one POST endpoint

Write a JSON schema validator using struct tags

Create a simple in‑memory key/value store with TTL expiration using goroutines

If you want, I can turn these into a progressive 30‑day Go learning plan or generate starter code templates for each exercise.

Today
A clean, simple starter set of Go exercises that actually build fluency comes from touching the core language features one by one. Each task is small enough to finish quickly but deep enough to teach something real.

🧱 Core Syntax and Functions
These warm you up on Go’s structure, types, and flow control.

1. Basic Calculator Function
   Write a function Calc(a, b int, op string) (int, error) that supports +, -, *, /.
   Return an error for unknown operators or division by zero.

2. Even/Odd Filter
   Given a slice of integers, return two slices: one with evens, one with odds.
   Practice:

Ranging over slices

Appending

Returning multiple values

📦 Slices, Maps, and Structs
These build comfort with Go’s data structures and memory behavior.

3. Word Frequency Counter
   Given a string, return a map[string]int of word → count.
   Normalize to lowercase and strip punctuation.

4. Todo List Struct
   Define:

go
type Todo struct {
Title string
Done  bool
}
Then:

Store todos in a slice

Add functions to add, mark done, and list incomplete items

⚙️ Interfaces and Methods
These teach Go’s implicit interface model.

5. Shape Interface
   Define:

go
type Shape interface {
Area() float64
}
Implement Circle and Rectangle.
Write a function that takes a slice of Shape and returns the total area.

6. Stringer Practice
   Create a struct Person{Name string, Age int} and implement the fmt.Stringer interface so printing a Person shows a custom message.

🔄 Concurrency Basics
These introduce goroutines and channels without overwhelming complexity.

7. Concurrent Counter
   Launch 5 goroutines. Each increments a shared counter 1000 times.
   Use a mutex to prevent race conditions.
   Print the final count.

8. Channel Pipeline
   Create a pipeline:

Goroutine A sends numbers 1–10 into a channel

Goroutine B reads them, squares them, and sends results to another channel

Main goroutine prints the squared values

🌐 Small Real-World Mini‑Project
This ties everything together in a practical way.

9. File Line Counter CLI
   Write a CLI tool that:

Accepts a filename

Reads the file line by line

Counts lines containing a given substring

Prints the result
Use buffered I/O (bufio.Scanner) and return errors cleanly.