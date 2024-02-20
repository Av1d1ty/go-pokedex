package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {"HELLO world", []string{"hello", "world"}},
    }

    for _, cs := range cases {
        actual := cleanInput(cs.input)
        if len(actual) != len(cs.expected) {
            t.Errorf("Expected %v, but got %v", cs.expected, actual)
            continue
        }
        for i := range actual {
            if actual[i] != cs.expected[i] {
                t.Errorf("Expected %v, but got %v", cs.expected, actual)
                break
            }
        }
    }
}
