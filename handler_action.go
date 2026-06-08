package main

type HandlerAction struct {
	Title    string
	Tooltip  string
	IsActive bool
	Callback func(string) string
}
