package model

type Request struct {
	No    int      `validate:"required,min=10,max=100"`
	Types []string `validate:"required,validType"`
	Name  string   `validate:"required"`
}
