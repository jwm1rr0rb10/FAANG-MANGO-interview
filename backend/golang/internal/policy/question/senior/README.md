# Live Coding Tasks.


## 1. Concurrency & Parallelism

### These are the most common "Go-specific" tasks.

- Worker Pool: Create a pool of N workers to process a queue of jobs.

    - Example: Use a job channel and a results channel to process 100 integers using 5 goroutines.

- Fan-in/Fan-out: Distribute tasks to multiple workers and aggregate the results into a single channel.

    - Example: Several functions generate random numbers; one function collects and prints them.

- Rate Limiter: Implement a function that allows only X requests per second.

    - Example: Use time.Tick or a token bucket approach with a channel.

- Graceful Shutdown: Write a program that waits for a SIGINT (Ctrl+C) and closes all open resources before exiting.

    - Example: Use os/signal and a context.WithCancel.

- Concurrent File Searcher: Search for a string in all files within a directory concurrently.

    - Example: Each file processed in a goroutine; use sync.WaitGroup to wait for completion.

- The "Select" Timeout: Implement a function that calls an API but returns an error if it takes longer than 2 seconds.

    - Example: Use select with chan and time.After.

- Ping-Pong: Create two goroutines that send "ping" and "pong" back and forth via a channel.

    - Example: Ensures you understand unbuffered channel blocking.

2. Slices, Maps & Logic

- Frequency Map: Count the occurrence of each word in a string.

    - Example: map[string]int to store counts.

- Remove Duplicates: Write a function that returns a slice with unique elements from an input slice.

    - Example: Use a map[T]struct{} to track seen items efficiently.

- Reverse a Slice: Reverse a slice in-place without using extra memory.

    - Example: Two-pointer approach swapping i and j.

- Merge Sorted Slices: Given two sorted slices, merge them into one sorted slice.

    - Example: Standard merge step from Merge Sort.

- Find the "Missing Number": Given a slice containing 0 to n, find the one that is missing.

    - Example: Use the sum formula: 2n(n+1)​−actual_sum.

## 3. Data Structures & Interfaces

- LRU Cache: Implement a Least Recently Used cache with Get and Put methods.

    - Example: Uses a map for O(1) access and a doubly linked list for ordering.

- Stack/Queue Implementation: Create a generic Stack or Queue using a slice.

    - Example: Use any or generics [T any] (Go 1.18+).

- Custom Interface Implementation: Create a "Shape" interface and implement it for "Circle" and "Square".

    - Example: Tests understanding of implicit interface implementation.

## 4. Systems & Web (The net/http package)

- Basic HTTP Middleware: Write a middleware that logs the execution time of every request.

    - Example: A function that takes http.Handler and returns http.Handler.

- JSON API Client: Fetch data from a public API (like JSONPlaceholder) and unmarshal it into a struct.

    - Example: Use http.Get and json.NewDecoder.

- In-Memory Store: Build a thread-safe Key-Value store with an HTTP interface.

    - Example: Use a map protected by a sync.RWMutex.

5. Algorithmic (Go Flavor)

- Binary Search: Implement binary search on a sorted slice.

    - Example: Iterative or recursive approach.

- FizzBuzz (The Classic): Print 1 to 100, but for multiples of 3 print "Fizz", 5 "Buzz", and both "FizzBuzz".

    - Example: Often used as a "warm-up" to check syntax.