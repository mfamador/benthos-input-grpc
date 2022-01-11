// Package repository contains the repositories interfaces
package repository

// Poster exposes the interface to post messages
type Poster interface {
	Post(string) error
}
