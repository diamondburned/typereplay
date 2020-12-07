package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/pkg/errors"
)

type Action interface {
	Do(tape *ActionTape)
}

type ActionTape struct {
	Ticker  time.Ticker
	Actions []Action
}

func (t *ActionTape) Do() {
	for _, action := range t.Actions {
		action.Do(t)
	}
}

type WaitAction struct {
	Duration time.Duration
}

func (a WaitAction) Do(tape *ActionTape) {
	time.Sleep(a.Duration)
}

type PutsAction struct {
	Text []rune
}

func typeText(text []rune, tape *ActionTape) {
	for _, r := range text {
		robotgo.UnicodeType(uint32(r))
		<-tape.Ticker.C
	}
}

func (a PutsAction) Do(tape *ActionTape) {
	typeText(a.Text, tape)
}

type TypeAction struct {
	Text []rune
}

func (a TypeAction) Do(tape *ActionTape) {
	typeText(a.Text, tape)

	robotgo.KeyTap("enter")
	<-tape.Ticker.C
}

type TapAction struct{ Key string }

func (a TapAction) Do(tape *ActionTape) {
	robotgo.KeyTap(a.Key)
	<-tape.Ticker.C
}

var (
	sepSpace  = []byte(" ")
	poundByte = []byte("#")
)

func ParseInput(input io.Reader) (*ActionTape, error) {
	var actions []Action

	var duration = 100 * time.Millisecond

	var scanner = bufio.NewScanner(input)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		words := bytes.SplitN(line, sepSpace, 2)

		// ignore comments and blank lines
		if len(words[0]) == 0 || bytes.HasPrefix(words[0], poundByte) {
			continue
		}

		var argument []byte
		if len(words) == 2 {
			argument = words[1]
		}

		switch command := string(words[0]); command {
		case "setduration":
			d, err := time.ParseDuration(string(argument))
			if err != nil {
				return nil, errors.Wrap(err, "invalid value for setduration")
			}
			duration = d

		case "enter":
			actions = append(actions, TapAction{Key: "enter"})

		case "space":
			actions = append(actions, TapAction{Key: "space"})

		case "tap":
			actions = append(actions, TapAction{Key: string(argument)})

		case "type":
			text := []rune(string(argument))
			actions = append(actions, TypeAction{Text: text})

		default:
			return nil, fmt.Errorf("unknown command %q", command)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to scan input")
	}

	return &ActionTape{
		Ticker:  *time.NewTicker(duration),
		Actions: actions,
	}, nil
}
