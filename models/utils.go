package models

import (
	"fmt"
	"os"
	"os/exec"
)

type State int

const (
	Tree State = iota
	Burning
	Burnt
	Empty
)

func (s *State) ToString() string {
	switch *s {
	case Tree:
		return "t "
	case Burning:
		return "* "
	case Burnt:
		return ", "
	case Empty:
		return "  "
	default:
		return "? "
	}
}

func (s *State) ToInt() int {
	return int(*s)
}

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func GameOver() {
	gameOver := `
  ____                         ___                
 / ___| __ _ _ __ ___   ___   / _ \__   _____ _ __ 
| |  _ / _` + "`" + ` | '_ ` + "`" + ` _ \ / _ \ |
| |_| | (_| | | | | | |  __/ | |_| |\ V /  __/ |
 \____|\__,_|_| |_| |_|\___|  \___/  \_/ \___|_|
                                                  
`
	fmt.Println(gameOver)
}
