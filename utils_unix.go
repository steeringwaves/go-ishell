//go:build darwin || dragonfly || freebsd || (linux && !appengine) || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux,!appengine netbsd openbsd solaris

package ishell

import (
	readline "github.com/steeringwaves/go-readline"
)

func clearScreen(s *Shell) error {
	_, err := readline.ClearScreen(s.writer)
	return err
}
