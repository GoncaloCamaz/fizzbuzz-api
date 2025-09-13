/*
Package repository holds the functionality needed to interact with storage
*/
package repository

// PGRepo represents a postgres repository
type PGRepo struct{}

// NewPGRepo returns a new PGRepo
func NewPGRepo() *PGRepo {
	return &PGRepo{}
}
