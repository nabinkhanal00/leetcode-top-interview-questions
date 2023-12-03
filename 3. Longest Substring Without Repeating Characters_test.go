package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3(t *testing.T) {
	t.Run("abba", func(t *testing.T) {
		got := lengthOfLongestSubstring("abba")
		assert.Equal(t, 2, got)
	})
}
