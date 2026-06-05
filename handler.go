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
	action := HandlerAction{
		Title:    title,
		Tooltip:  tooltip,
		IsActive: true,
		Callback: callback,
	}

	h.actions = append(h.actions, &action)
}

func (h *Handler) Len() int {
	return len(h.actions)
}

func (h *Handler) ToggleAction(index int) {
	if index >= 0 && index < len(h.actions) {
		h.actions[index].IsActive = !h.actions[index].IsActive
	}
}

type ActionInfo struct {
	Title    string
	Tooltip  string
	IsActive bool
}

func (h *Handler) GetActionsInfo() []ActionInfo {
	info := make([]ActionInfo, len(h.actions))
	for i, action := range h.actions {
		info[i] = ActionInfo{
			Title:    action.Title,
			Tooltip:  action.Tooltip,
			IsActive: action.IsActive,
		}
	}
	return info
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
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		inputStr := string(data)
		output := h.Process(inputStr)

		if output != inputStr {
			clipboard.Write(clipboard.FmtText, []byte(output))
		}
	}
}
