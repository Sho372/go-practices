package main

// Type for implementation of Logic interface
type LogicProvider struct{}

// Implementating the Logic intreface
func (lp LogicProvider) Process(data string) string {
	return ""
}

// Logic intreface that I interested in at this time
type Logic interface {
	Process(data string) string
}

// Using the Logic interface
type Client struct {
	L Logic
}

// Expecting the Logic interface is implemented
func (c Client) Program() {
	data := "get data from somewhere"
	c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
