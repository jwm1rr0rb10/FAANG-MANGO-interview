package main


// так я говорил что я брал за основу, только в той репе, я пошел дальше разделил на слои. 
.. я взял за основу как раз та
func questionGolang(question []string) {

	// ------------------------------------------------
	// BEGINNER QUESTIONS
	// ------------------------------------------------
	questions = append(questions, `What is the blank identifier?`)
	questions = append(questions, "What's the difference between declare, assign, and initialize?")
	questions = append(questions, "Explain allocation and initialization.")
	questions = append(questions, "What is the difference between make([]int, 10), make([]int, 0, 10), make([]int, 10, 10)")
	questions = append(questions, "What types can a map use as a key in the Go programming language?")
	questions = append(questions, "What does it mean to write idiomatic Go code?")
	questions = append(questions, "Which do you choose: performance or readability?")
	questions = append(questions, "Why was Go created, and who created it?")
	questions = append(questions, "Can you explain what is meant by 'Go is strongly typed?'")
	questions = append(questions, "What is the var keyword used for and when do you use it?")
	questions = append(questions, `When dealing with computer architecture, what does the 'word size' mean?`)
	questions = append(questions, "How does a computer work?")
	questions = append(questions, "What is a compiler?")
	questions = append(questions, "What is a garbage collector?")
	questions = append(questions, "Can you give an example of a problem you've solved using Go? What challenges did you encounter, and how did you overcome them?")
	questions = append(questions, "Explain how package management works in Go.")
	questions = append(questions, `What is the difference between make and new?`)
	questions = append(questions, `Tell us about bytes, code points, and characters in relation to strings and UTF-8.`)
	questions = append(questions, `What is embedding a struct and inner-type promotion?`)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)

	questions = append(questions, `Fix this code so that type HUMAN is embedded in type SECRETAGENT

	type human struct {
		name  string
		email string
	}
	
	type secretAgent struct {
		person human
		id  string
	}

	`)
	// fixed: https://go.dev/play/p/zO0ZBicZM7u

	questions = append(questions, `Fix this code working with a map:
	package main

	import "fmt"
	
	func main() {
		var m map[string]int
		m["b"] = 42
		fmt.Println(m)
	}
	
	`)
	// not fixed: https://go.dev/play/p/d001JBoJcAV
	// fixed: https://go.dev/play/p/RCVcb_4z2oK

	questions = append(questions, `What's problematic with this code working with append:

	func slice() {
		xi := make([]int, 10)
		for i := 0; i < 10; i++ {
			xi1 = append(xi1, i)
		}
		fmt.Println(xi1)
	}		
	`)
	// not fixed: https://go.dev/play/p/04NLYE1I7W4
	// fixed: https://go.dev/play/p/guYNTv-lfFi

	questions = append(questions, `What will this code print:
	
	func main() {
		sports := make([]string, 5)
		sports[0] = "ski"
		sports[1] = "surf"
		sports[2] = "swim"
		sports[3] = "sail"
		sports[4] = "sumo wrestling"
	
		xs := sports[1:3]
		xs[0] = "CHANGED"

		inspectSlice(sports)
		inspectSlice(xs)
	}
	
	func inspectSlice(xs []string) {
		fmt.Printf("len: %v \ncap: %v \n", len(xs), cap(xs))
		for i := range xs {
			fmt.Printf("%p \t %v \n", &xs[i], xs[i])
		}
	}
	`)
	// https://go.dev/play/p/Rw5cmIrXlwT

	questions = append(questions, `Fix this code so that when xs[0] is changed this doesn't change 'sports':
	
	func main() {
		sports := make([]string, 5)
		sports[0] = "ski"
		sports[1] = "surf"
		sports[2] = "swim"
		sports[3] = "sail"
		sports[4] = "sumo wrestling"
	
		xs := sports[1:3]
		xs[0] = "CHANGED"
		inspectSlice(sports)
		inspectSlice(xs)
	}
	
	func inspectSlice(xs []string) {
		fmt.Printf("len: %v \ncap: %v \n", len(xs), cap(xs))
		for i := range xs {
			fmt.Printf("%p \t %v \n", &xs[i], xs[i])
		}
	}
	`)
	// not fixed: https://go.dev/play/p/Rw5cmIrXlwT
	// fixed: https://go.dev/play/p/N-32AbjJJZw

}

func liveConding(question []string) {

}
