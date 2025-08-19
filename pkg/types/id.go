package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// ID represents a unique identifier
type ID string

// NewID generates a new unique ID
func NewID() ID {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		// Fallback to a simple counter-based approach if crypto/rand fails
		return ID(fmt.Sprintf("id_%d", generateCounter()))
	}
	return ID(hex.EncodeToString(bytes))
}

// String returns the string representation of the ID
func (id ID) String() string {
	return string(id)
}

// IsEmpty checks if the ID is empty
func (id ID) IsEmpty() bool {
	return string(id) == ""
}

// Validate validates the ID format
func (id ID) Validate() error {
	if id.IsEmpty() {
		return fmt.Errorf("ID cannot be empty")
	}
	if len(string(id)) < 8 {
		return fmt.Errorf("ID must be at least 8 characters long")
	}
	return nil
}

// Simple counter for fallback ID generation
var counter int64

func generateCounter() int64 {
	counter++
	return counter
}
