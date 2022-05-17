package core

import (
	"fmt"
	"internal/itoa"
	"os"
	"strings"
)

type UCI struct {
	board     Board
	useUCI    bool
	options   map[string]string
	debugMode bool
}

func longAlgebraicToMoves(algebraic string) (uint8, uint8) {
	from := AlgebraicToUint8(algebraic[0:2])
	to := AlgebraicToUint8(algebraic[2:4])
	return from, to
}

func removeExcessWhitespace(com string) string {
	var sb strings.Builder
	gotSpace := false
	gotFirstChar := false
	for _, ch := range com {
		if ch == ' ' || ch == '\t' {
			if !gotSpace && gotFirstChar {
				gotSpace = true
				sb.WriteRune(' ')
			} else {
				gotSpace = false
				gotFirstChar = true
				sb.WriteRune(ch)
			}
		}
	}
	return sb.String()
}

func printError(err string) {
	fmt.Println("[Error] " + err)
}

func printValidOptions() {
	// TODO: add some valid options
}

func returnToGUI(ret string) {
	fmt.Println(ret)
}

func (uci *UCI) ParseCommand(com string) {
	com = removeExcessWhitespace(com)
	split := strings.Split(com, " ")
	if len(split) < 1 {
		printError("Command too short")
	}
	switch split[0] {
	case "uci":
		uci.useUCI = true
		returnToGUI("id name gochess")
		returnToGUI("id author OFFTKP")
		printValidOptions()
		returnToGUI("uciok")
	case "setoption":
		if len(split) < 5 {
			printError("Expected 4 parameters, got " + itoa.Itoa(len(split)-1))
			return
		}
		if split[1] != "name" {
			printError("Bad 1st parameter, expected 'name'")
			return
		}
		if split[3] != "value" {
			printError("Bad 3rd parameter, expected 'value'")
			return
		}
		uci.options[split[2]] = strings.Join(split[4:], " ")
	case "debug":
		if split[1] == "on" {
			uci.debugMode = true
		} else if split[1] == "off" {
			uci.debugMode = false
		} else {
			printError("Bad 1st parameter, expected on/off")
		}
	case "isready":
		returnToGUI("readyok")
	case "ucinewgame":
		uci.board.Reset()
	case "quit":
		os.Exit(0)
	default:
		printError("Unknown command:" + split[0])
	}
}
