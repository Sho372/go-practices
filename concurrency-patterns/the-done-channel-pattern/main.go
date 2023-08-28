package main

func searchData(s string, searchaers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchaers {
		// worker goroutine
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}
