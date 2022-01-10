// Package repository contains the repositories interfaces
package repository

// Server exposes the interface to post messages
type Server interface {
	Post(string) error
}
