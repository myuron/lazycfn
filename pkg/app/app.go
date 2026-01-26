package app

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/myuron/lazycfn/pkg/ui"
)

type App struct {
	gui *gocui.Gui
}

func New() (*App, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	maxX, maxY := g.Size()
	widgets := ui.NewDefaultLayout(maxX, maxY)

	managers := make([]gocui.Manager, len(widgets))
	for i, w := range widgets {
		managers[i] = w
	}
	g.SetManager(managers...)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	return &App{gui: g}, nil
}

func (a *App) Run() error {
	err := a.gui.MainLoop()
	if err == gocui.ErrQuit {
		return nil
	}
	return err
}

func (a *App) Close() {
	a.gui.Close()
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
