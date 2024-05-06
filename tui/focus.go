package tui

import "fmt"

type FocusState uint

const (
	FcOnTable   FocusState = iota
	FcOnCmdLine FocusState = iota
	FcOnList    FocusState = iota
)

func nextFocusState(fs FocusState) (FocusState, error) {
	switch fs {
	case FcOnTable:
		return FcOnCmdLine, nil
	case FcOnCmdLine:
		return FcOnList, nil
	case FcOnList:
		return FcOnTable, nil
	default:
		return fs, fmt.Errorf("FocusState has unknown value %d", fs)
	}
}
func prevFocusState(fs FocusState) (FocusState, error) {
	switch fs {
	case FcOnTable:
		return FcOnList, nil
	case FcOnCmdLine:
		return FcOnTable, nil
	case FcOnList:
		return FcOnCmdLine, nil
	default:
		return fs, fmt.Errorf("FocusState has unknown value %d", fs)
	}
}
