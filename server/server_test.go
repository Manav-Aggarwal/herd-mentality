package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadQuestions(t *testing.T) {
	s := &server{}

	questions := strings.NewReader("question 1\nquestion 2\n question3\n")

	s.loadQuestions(questions)
	assert.Len(t, s.questions, 3)
}
