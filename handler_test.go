package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	h := NewHandler()
	h.AddAction("Trim", "Remove leading and trailing whitespace", Trim)
	h.AddAction("Capitalize", "Capitalize words if text is fully uppercase", CapitalizeUppercase)

	if h.Len() != 2 {
		t.Fatalf("Expected 2 actions, got %d", h.Len())
	}

	// Test processing when both are active
	input := "  TEST INPUT  "
	expected := "Test Input"
	got := h.Process(input)
	if got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}

	// Toggle first action (Trim) to inactive
	h.ToggleAction(0)

	// Trim is inactive, so it shouldn't trim whitespace before capitalize.
	// But wait, CapitalizeUppercase checks if the text is fully uppercase (with spaces).
	// "  TEST INPUT  " in uppercase is "  TEST INPUT  ", so it gets converted to "  Test Input  ".
	expectedAfterToggle := "  Test Input  "
	gotAfterToggle := h.Process(input)
	if gotAfterToggle != expectedAfterToggle {
		t.Errorf("Expected %q after toggle, got %q", expectedAfterToggle, gotAfterToggle)
	}

	// Get actions info and verify status
	info := h.GetActionsInfo()
	if len(info) != 2 {
		t.Fatalf("Expected 2 actions info, got %d", len(info))
	}
	if info[0].IsActive {
		t.Error("Expected action index 0 to be inactive")
	}
	if !info[1].IsActive {
		t.Error("Expected action index 1 to be active")
	}
}
