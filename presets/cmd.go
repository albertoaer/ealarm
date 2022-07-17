package presets

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type PresetCommand struct {
	cmdLine []string
}

func (m *PresetCommand) Launch(next chan bool) {
	cmd := exec.Command(m.cmdLine[0], m.cmdLine[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	if e != nil {
		fmt.Printf("Error running preset command, %s", e.Error())
	}
	next <- (e == nil)
}

func (m *PresetCommand) UnmarshalJSON(data []byte) error {
	var line []string
	if e := json.Unmarshal(data, &line); e != nil {
		return e
	}
	if len(line) == 0 {
		return errors.New("Expecting at least command name")
	}
	m.cmdLine = line
	return nil
}
