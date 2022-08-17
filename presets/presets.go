package presets

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
)

type Profile map[string]string
type Presets struct {
	Actions  map[string]*PresetCommand `json:"actions,omitempty"`
	Profiles map[string]Profile        `json:"profiles,omitempty"`
}

const presets_file string = "./presets.json"

var presets *Presets

func LoadPresets() *Presets {
	if presets == nil {
		loaded := Presets{}
		if _, e := os.Stat(presets_file); errors.Is(e, fs.ErrNotExist) {
			if e := ioutil.WriteFile(presets_file, []byte("{}"), os.ModePerm); e != nil {
				panic(e)
			}
		}
		data, e := ioutil.ReadFile(presets_file)
		if e != nil {
			panic(e)
		}
		if e := json.Unmarshal(data, &loaded); e != nil {
			panic(e)
		}
		presets = &loaded
	}
	return presets
}
