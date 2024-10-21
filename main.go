package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"schedule/repository/db/dao"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	dao.InitSqlite()
	noteDao := dao.NewNoteDao(nil)
	// noteDao := db.NewNoteApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "schedule",
		DisableResize: true, // 禁用调整窗口尺寸
		Width:         350,
		Height:        600,
		// MinWidth:      350,
		// MinHeight:     600,
		// WindowStartState: options.Minimised,//窗口启动状态
		// MaxWidth:  1280,
		// MaxHeight: 1024,
		AlwaysOnTop: true,  // 窗口固定在最顶层
		StartHidden: false, // 启动时隐藏窗口
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0}, // 背景颜色
		// OnStartup:        app.startup,
		Bind: []interface{}{
			noteDao,
		},
		// 修改 windows配置
		Windows: &windows.Options{
			WebviewIsTransparent:              true, // 网页透明
			WindowIsTranslucent:               true, // 窗口透明
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			Theme:                             windows.SystemDefault,
			// BackdropType:                      windows.Acrylic,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(20, 20, 20),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				DarkModeBorder:     windows.RGB(20, 0, 20),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
			},
		},
		// Frameless: true, // 添加无边框
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
