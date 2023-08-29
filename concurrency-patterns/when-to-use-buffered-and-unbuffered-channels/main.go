package main

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc) // buffered channle
	for i := 0; i < conc; i++ {
		go func() {
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func main() {

}
