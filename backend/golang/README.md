# Interview Questions for Go Developer

## Question Structure by Levels

### 👶 Junior (Entry Level)
*Language basics, data structures, simple functions.*

#### Data Types and Syntax
**Packages and Scope**
- What are packages in Go?
- How does scope work in Golang?
- What is a global variable and how to declare it?

**Operators and Basic Syntax**
- What are the main operators in Go?
- How does switch case work? Can multiple conditions be executed in one case?
- What is variable capture in a loop?

**Strings**
- What are strings in Go?
- How can you manipulate strings?
- How to determine the number of characters in a string?
- What is `len()` for strings and what does it return?

**Numeric Types (Int)**
- What numeric types are available in Go?
- What's the difference between int and uint?
- What is a regular int and what does its size depend on?
- How to convert string to int and vice versa? Why can't you use `int(string)` and `string(int)`?

**Constants (Const)**
- What are constants and can they be changed?
- What is `iota` and how does it work?

**Arrays and Slices**
- What is a slice and how does it differ from an array?
- How does the `append()` function work?
- What is zero value for a slice and what operations are possible on it?

**Maps**
- What is a map in Go?
- How to declare and initialize a map?
- What happens in a map if you don't use make or short assign?
- What methods are available for maps (len, cap)?

**Interfaces**
- What are interfaces in Go?
- What is an empty interface `interface{}` / `any`?
- How is polymorphism implemented in Go?

#### Functions and Constructs
**Functions and Flow Control**
- How does `defer` work? What is the order of calling multiple `defer` statements?
- How many values can be returned from a function?
- What is an unnamed block (curly braces with no declared operator in a function)?

#### OOP in Go
**OOP Basics**
- How is encapsulation implemented in Go?
- How is polymorphism implemented in Go?
- How is inheritance (or its alternative) implemented in Go?

#### Concurrency (Basic Level)
**Concurrency Basics**
- What is a goroutine? How to start it?
- What is a channel? What's the difference between buffered and unbuffered channels?
- What is `WaitGroup` and what is it used for?

### 🧑 Middle (Intermediate Level)
*In-depth understanding of language mechanisms, memory management, concurrency.*

#### Data Types (Advanced)
**Strings (Advanced Level)**
- What are the nuances when iterating over a string (runes vs bytes)?
- How are strings structured at the runtime level?

**Slices (Advanced Level)**
- What array size is allocated for a slice when it expands beyond its capacity (growth strategy)?
- How does memory allocation work for slices?

**Maps (Advanced Level)**
- How is map implemented internally in Go?
- How does map grow? What is "evacuation" and when does it occur?
- Why (previously) couldn't you take the address of a map element? (Modern versions: you can, but with caveats)
- What types can be map keys? Can a struct be a key?
- Is map thread-safe? What is race condition?

**Interfaces (Advanced Level)**
- What is a nil-interface and nil under the hood?
- How to convert interface to another type (type assertion, type switch)?
- How to determine the type of an interface?
- On which side should interfaces be described - on the sending or receiving side? (Principle "Accept interfaces, return structs")

#### Concurrency and Synchronization
**Goroutines and Scheduler**
- How does a goroutine differ from an OS thread?
- What are the minimum and maximum stack sizes of a goroutine?
- How does the Go scheduler work? What are M, P, G?
- Do goroutines share CPU time equally?

**Channels (Advanced Level)**
- What happens if you write to/read from a nil channel?
- What happens if you write to/read from a closed channel?
- What happens if you write to/read from a buffered/unbuffered channel?
- How to close a channel and what happens after closing?
- How does select work? How to make it non-blocking?
- What is the execution order of case operations in select?

**Synchronization Primitives**
- How is mutex implemented? What's the difference between `sync.Mutex` and `sync.RWMutex`?
- What is atomic used for?
- What is `sync.Map` and when to use it?
- What other synchronization primitives do you know (`sync.Once`, `sync.Pool`, `sync.Cond`)?

**Contexts**
- What is context and what is it used for?
- What's the difference between `context.Background()` and `context.TODO()`?
- What are the differences between `context.WithCancel`, `context.WithDeadline`, `context.WithTimeout`?
- How to pass values and read them from context?
- How to handle context cancellation?

#### Memory Management
**Memory Management**
- How are parameters passed to functions (by value or by reference)?
- Are there any special behaviors when passing maps and slices to functions?
- What are heap and stack? How does escape analysis work?
- What do `*` and `&` denote? How do pointers work?

**Garbage Collector (GC)**
- What is the garbage collector in Go?
- How does the mark and sweep algorithm work (tri-color marking)?
- When does the garbage collector start?
- What resources does it consume?

#### Testing
**Testing (Basic Level)**
- What are Table-Driven Tests (TDT)?
- How to name packages with tests?
- What are static analyzers (linters)?

### 🧙‍♂️ Senior (Senior Level)
*Internal language structure, fine-tuning, system design.*

#### Compiler and Low-Level Features
**Go Compilation**
- What stages does Go compilation consist of?
- What is static compilation/linking and what are its features in Go?
- What compiler directives do you know? (e.g., `//go:linkname`, `//go:noinline`, `//go:noescape`)

**Internal Type Implementation**
- How are strings implemented exactly (`stringHeader`)?
- How does key lookup work in map (hash table, buckets)?
- How are interfaces structured at runtime level (`iface`, `eface`)?

**Internal Concurrency Implementation**
- How are goroutines implemented at runtime level?
- How does context switching between goroutines occur?
- How to manually set the number of processors (`GOMAXPROCS`)?
- What is the maximum number of goroutines that can be launched?

#### Performance and Profiling
**Profiling and Optimization**
- How does pprof work? How to analyze CPU, memory, and blocking profiles?
- Example of using pprof for problem diagnosis.
- How does the profiler work in principle?
- What is graceful shutdown and how to implement it?

**Garbage Collector (Advanced Level)**
- Details of tri-color algorithm implementation (tri-color mark).
- How does GC affect pauses (STW - Stop The World)?
- How to optimize GC work?

**Problem Diagnosis**
- How to detect data race? (Flag `-race`)
- How to debug deadlock?
- How to profile memory consumption?

#### Architecture and Design
**Architectural Principles**
- How are SOLID principles implemented in Go?
- Especially: how is Dependency Inversion implemented?
- Package design in Go. How to manage dependencies?
- Design patterns in Go (worker pool, circuit breaker, dependency injection).

**Generics**
- What are generics in Go? When were they introduced?
- How to declare a parameterized function?
- How to declare a parameterized type?
- What are constraints?

**System Programming**
- How do system calls work in Go?
- How is low-level network I/O implemented?
- How does file I/O work?

#### Advanced Testing
**Testing (Advanced Level)**
- How to write integration and functional tests?
- Using mocks (gomock, testify).
- Testing concurrent code.
- Writing benchmarks. What is a benchmarking error?

## Additional Important Topics
*Not included in main list but important*

#### Functions and Methods
- What is the `init()` function? How and when is it called?
- Difference between function and method. Receivers by value and by pointer.
- When to use pointer receiver vs value receiver?

#### Structures (Structs)
- What are struct tags? For example, `json:"name"`.
- How do tags work with reflection?
- Struct embedding (composition).

#### Reflection
- What is the `reflect` package?
- When is reflection needed and why should it be avoided?
- Examples of reflection usage.

#### Error Handling
- Idiomatic error handling in Go.
- Wrapping errors (`errors.Is`, `errors.As`).
- Panics vs errors. When to use what?
- Custom error types.

#### Dependency Management
- Go modules (`go.mod`, `go.sum`).
- SemVer and versioning.
- Vendor directory.

#### IO Operations
- `io.Reader` and `io.Writer` interfaces.
- Byte buffers (`bytes.Buffer`).
- Working with files.

#### Network Programming
- Creating HTTP servers and clients.
- Middleware in Go.
- REST API development.

#### Cross-Compilation and Deployment
- How to build binaries for other OS/architecture?
- Reducing binary size.
- Docker images for Go applications.

#### Development Tools
- `go fmt`, `go vet`, `go mod`.
- Static analyzers: golangci-lint, staticcheck.
- Code generation (`go generate`).

#### Embedding Resources
- `embed` package for including resources in binaries.

## Key Go Concepts

#### OOP in Go
Go is not a pure OOP language but supports key principles:
- **Encapsulation** - through packages (exported/unexported identifiers)
- **Inheritance** - through composition and struct embedding
- **Polymorphism** - through interfaces (duck typing)
- **Abstraction** - through interfaces and unexported types

#### Concurrency vs Parallelism
- **Concurrency** - executing multiple tasks by switching between them
- **Parallelism** - executing multiple tasks simultaneously
- **Goroutines** - lightweight threads managed by Go runtime
- **Channels** - communication method between goroutines (share memory by communicating)

#### Go Features
- Simple and readable syntactic structure
- Static typing with type inference
- Garbage collection
- Built-in concurrency support
- Fast compilation
- Static linking (single binary file)

## Useful Preparation Resources

#### Official Resources
- [Official Documentation](https://go.dev)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Go Blog](https://blog.golang.org)

#### Books
- "The Go Programming Language" (Donovan & Kernighan)
- "Concurrency in Go" by Katherine Cox-Buday
- "100 Go Mistakes and How to Avoid Them" by Teiva Harsanyi

#### Community
- [Go Forum](https://forum.golangbridge.org/)
- [r/golang](https://www.reddit.com/r/golang/)
- [Go Time Podcast](https://changelog.com/gotime)

#### Practice Platforms
- [LeetCode Go Problems](https://leetcode.com/tag/golang/)
- [Exercism Go Track](https://exercism.org/tracks/go)






















































