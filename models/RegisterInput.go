package models

// RegisterInput defines the structure for registration input.

type RegisterInput struct {
	StudentID uint `json:"studentId"` // Use the appropriate type for StudentID; uint is a common choice for IDs.
}
