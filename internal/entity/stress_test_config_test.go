package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcurrencyEqualsOne(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 1, 1)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(1), requests)
	assert.Equal(t, uint64(0), remaining)
}

func TestConcurrencyWithEvenValue2(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 1, 1)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(1), requests)
	assert.Equal(t, uint64(0), remaining)
}
