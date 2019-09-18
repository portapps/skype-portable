//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Skype.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"io/ioutil"
	"os"
	"path"

	_ "github.com/kevinburke/go-bindata"
	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/proc"
	"github.com/portapps/portapps/pkg/shortcut"
	"github.com/portapps/portapps/pkg/utl"
	"github.com/portapps/skype-portable/assets"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("skype-portable", "Skype"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "Skype.exe")
	app.Args = []string{
		"--datapath=" + app.DataPath,
	}

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Skype Portable.lnk")
	defaultShortcut, err := assets.Asset("Skype.lnk")
	if err != nil {
		Log.Error().Err(err).Msg("Cannot load asset Skype.lnk")
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error().Err(err).Msg("Cannot write default shortcut")
	}

	// Remove old appdata when closed
	defer func() {
		if err := os.RemoveAll(path.Join(os.Getenv("APPDATA"), "Skype")); err != nil {
			Log.Error().Err(err).Msg("Cannot remove old appdata folder")
		}
		if err := os.RemoveAll(path.Join(os.Getenv("APPDATA"), "Microsoft", "Skype for Desktop")); err != nil {
			Log.Error().Err(err).Msg("Cannot remove old appdata folder")
		}
	}()

	// Update default shortcut
	err = shortcut.Create(shortcut.Shortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       app.Process,
		Arguments:        shortcut.Property{Clear: true},
		Description:      shortcut.Property{Value: "Skype Portable by Portapps"},
		IconLocation:     shortcut.Property{Value: app.Process},
		WorkingDirectory: shortcut.Property{Value: app.AppPath},
	})
	if err != nil {
		Log.Error().Err(err).Msg("Cannot create shortcut")
	}
	defer func() {
		if err := os.Remove(shortcutPath); err != nil {
			Log.Error().Err(err).Msg("Cannot remove shortcut")
		}
	}()

	// Stop data collector on RtcPalVideoEtwSession
	defer func() {
		if err := proc.QuickCmd("logman", []string{
			"stop",
			"RtcPalVideoEtwSession",
			"-ets",
		}); err != nil {
			Log.Error().Err(err).Msg("cannot stop data collector on RtcPalVideoEtwSession")
		}
	}()

	app.Launch(os.Args[1:])
}
