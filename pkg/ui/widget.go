package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Widget struct {
	Name           string
	Body           string
	X0, Y0, X1, Y1 int
}

func (w *Widget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.Name, w.X0, w.Y0, w.X1, w.Y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := fmt.Fprint(v, w.Body); err != nil {
			return err
		}
	}
	return nil
}
