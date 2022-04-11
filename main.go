package main

import (
	"embed"
	"log"
	"os"

	"github.com/martinlindhe/notify"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"easyclash/app"
	"easyclash/utils"
)

//go:embed build/appicon.png
var icon []byte

//go:embed frontend/dist
var ui embed.FS

//go:embed .easy_clash
var ruleSet embed.FS

var mode = ""

func init() {
	path, _ := os.UserHomeDir()
	err := utils.SaveDir(ruleSet, path, mode == "dev")
	if err != nil {
		notify.Alert("EasyClash", "", "init conf dir error", "")
		log.Fatal(err)
	}
	_ = os.Chmod(path+"/.easy_clash/easy_clash_core", os.ModePerm)
}

func main() {
	_app := app.NewApp()
	err := wails.Run(&options.App{
		Title:             "EasyClash",
		Width:             1024,
		Height:            768,
		MinWidth:          985,
		MinHeight:         768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: true,
		RGBA:              &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            ui,
		LogLevel:          logger.DEBUG,
		OnStartup:         _app.Startup,
		OnDomReady:        _app.DomReady,
		OnShutdown:        _app.Shutdown,
		Bind: []interface{}{
			_app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameVibrantLight,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "EasyClash",
				Message: "clash config manager",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		notify.Alert("", "", err.Error(), "")
		log.Fatal(err)
	}
}
