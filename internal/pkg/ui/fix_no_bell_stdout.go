package ui

import "github.com/chzyer/readline"

/// Taken from: https://github.com/manifoldco/promptui/issues/49#issuecomment-1012640880

type noBellStdout struct{}

func (n *noBellStdout) Write(p []byte) (int, error) {
	if len(p) == 1 && p[0] == readline.CharBell {
		return 0, nil
	}
	return readline.Stdout.Write(p)
}

func (n *noBellStdout) Close() error {
	return readline.Stdout.Close()
}
