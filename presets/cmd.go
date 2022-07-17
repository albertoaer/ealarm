package presets

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

type PresetCommand struct {
	cmd *exec.Cmd
}

func (m *PresetCommand) Launch(next chan bool) {
	e := m.cmd.Run()
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
	m.cmd = exec.Command(line[0], line[1:]...)
	return nil
}
