package main

func searchData(s string, searchaers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchaers {
		// worker goroutine
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s): // send
			case <-done: // receive
			}
		}(searcher)
	}
	r := <-result // receive
	close(done) // send
	return r
}
