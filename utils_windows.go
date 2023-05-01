//go:build windows
// +build windows

package ishell

import (
	readline "github.com/steeringwaves/go-readline"
)

func clearScreen(s *Shell) error {
	return readline.ClearScreen(s.writer)
}
