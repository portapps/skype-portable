//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
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
	Papp.Args = []string{ "--datapath=" + Papp.DataPath }
	Papp.WorkingDir = Papp.AppPath

	Launch()
}
