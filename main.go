//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"
	"path"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "skype-portable"
	Papp.Name = "Skype"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")
	Papp.Process = PathJoin(Papp.AppPath, "Skype.exe")
	Papp.Args = []string{"--datapath=" + Papp.DataPath}
	Papp.WorkingDir = Papp.AppPath

	Launch(os.Args[1:])

	oldAppData := path.Join(os.Getenv("APPDATA"), "Skype")
	if _, err := os.Stat(oldAppData); err == nil {
		os.RemoveAll(oldAppData)
	}
}
