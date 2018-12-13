//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Skype.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	_ "github.com/kevinburke/go-bindata"

	"io/ioutil"
	"os"
	"path"

	. "github.com/portapps/portapps"
	"github.com/portapps/skype-portable/assets"
)

func init() {
	Papp.ID = "skype-portable"
	Papp.Name = "Skype"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = CreateFolder(AppPathJoin("data"))
	Papp.Process = PathJoin(Papp.AppPath, "Skype.exe")
	Papp.Args = []string{"--datapath=" + Papp.DataPath}
	Papp.WorkingDir = Papp.AppPath

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Skype Portable.lnk")
	defaultShortcut, err := assets.Asset("Skype.lnk")
	if err != nil {
		Log.Error("Cannot load asset Skype.lnk:", err)
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error("Cannot write default shortcut:", err)
	}

	// Update default shortcut
	err = CreateShortcut(WindowsShortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       Papp.Process,
		Arguments:        WindowsShortcutProperty{Clear: true},
		Description:      WindowsShortcutProperty{Value: "Skype Portable by Portapps"},
		IconLocation:     WindowsShortcutProperty{Value: Papp.Process},
		WorkingDirectory: WindowsShortcutProperty{Value: Papp.AppPath},
	})
	if err != nil {
		Log.Error("Cannot create shortcut:", err)
	}

	Launch(os.Args[1:])

	oldAppData := path.Join(os.Getenv("APPDATA"), "Skype")
	_ = os.RemoveAll(oldAppData)
	_ = os.Remove(shortcutPath)
}
