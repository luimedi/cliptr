package main

import (
	"context"

	"golang.design/x/clipboard"
)

type Handler struct {
	actions []*HandlerAction
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AddAction(title string, tooltip string, callback func(string) string) {
	action := HandlerAction{true, callback}

	h.actions = append(h.actions, &action)

}

func (h *Handler) Process(text string) string {
	for _, action := range h.actions {
		if action.IsActive {
			text = action.Callback(text)
		}
	}
	return text
}

func (h *Handler) Listen() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		output := h.Process(string(data))

		if output != string(data) {
			clipboard.Write(clipboard.FmtText, []byte(output))
		}
	}
}
