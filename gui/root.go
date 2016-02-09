// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"fmt"
	"log"

	"github.com/gophersiesta/gophersiesta/Godeps/_workspace/src/github.com/jroimartin/gocui"
	"github.com/gophersiesta/gophersiesta/common"
)

type ScreenConf struct {
	Width       int
	Height      int
	TitleHeight int
	LinesOfHelp int
	AppsWidth   int
	UseFrame    int
}

var globalC int
var api *common.API

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v.Name() == "apps" {
		return g.SetCurrentView("placeholders")
	}
	if v.Name() == "placeholders" {
		return g.SetCurrentView("apps")
	}
	return g.SetCurrentView("apps")
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {

	globalC += 10
	_ = layout(g)

	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("apps", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("apps", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("apps", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("placeholders", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("placeholders", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("placeholders", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("apps", gocui.KeyEnter, gocui.ModNone, loadPlaceholders); err != nil {
		return err
	}
	if err := g.SetKeybinding("placeholders", gocui.KeyEnter, gocui.ModNone, editValue); err != nil {
		return err
	}
	if err := g.SetKeybinding("dialog", gocui.KeyEnter, gocui.ModNone, saveValue); err != nil {
		return err
	}
	return nil
}

func getScreenConf(g *gocui.Gui) ScreenConf {
	var sc ScreenConf
	w, h := g.Size()
	sc.Width = w
	sc.Height = h

	sc.TitleHeight = 1
	sc.LinesOfHelp = 2
	sc.AppsWidth = 30

	sc.UseFrame = 1
	if w < 100 {
		sc.UseFrame = 0
		sc.AppsWidth = 24
	}
	return sc
}

func layout(g *gocui.Gui) error {
	sc := getScreenConf(g)

	layoutTitle(g, sc)
	layoutHelp(g, sc)
	layoutApps(g, sc)
	layoutPlaceHolders(g, sc)
	layoutValues(g, sc)
	return nil
}

func layoutTitle(g *gocui.Gui, sc ScreenConf) error {

	if v, err := g.SetView("title", -1, -1, sc.Width, sc.TitleHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = false
		v.Frame = false
		v.BgColor = gocui.ColorCyan
		v.FgColor = gocui.ColorBlack
		fmt.Fprint(v, "GopherSiesta - Drunken Kittens Config Manager")
	}

	return nil
}

func layoutHelp(g *gocui.Gui, sc ScreenConf) error {
	if v, err := g.SetView("help", -1, sc.Height-(sc.LinesOfHelp+1), sc.Width, sc.Height); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = false
		v.Frame = false
		v.BgColor = gocui.ColorRed
		v.FgColor = gocui.ColorWhite

		// FIRST LINE OF HELP
		fmt.Fprint(v, "^N New App   ") // Some padding / centering functions here
		fmt.Fprint(v, "^R Refresh list   ")
		fmt.Fprint(v, "^W Search   ")
		fmt.Fprintln(v, "")

		// SECOND LINE OF HELP
		fmt.Fprint(v, "^S SOME COMMAND   ") // Some padding / centering functions here
		fmt.Fprint(v, "^A ANOTHER COMMAND   ")
		fmt.Fprint(v, "^C Close   ")
		fmt.Fprintln(v, "")

	}
	return nil
}

func layoutApps(g *gocui.Gui, sc ScreenConf) error {
	if v, err := g.SetView("apps", 0, sc.TitleHeight, sc.AppsWidth, sc.Height-(sc.LinesOfHelp+1)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Apps"
		v.Highlight = true

		apps, _ := api.GetApps()

		for _, app := range apps.Apps {
			fmt.Fprintln(v, app)
		}

		if err := g.SetCurrentView("apps"); err != nil {
			return err
		}
	}
	return nil
}

func layoutPlaceHolders(g *gocui.Gui, sc ScreenConf) (*gocui.View, error) {
	if v, err := g.SetView("placeholders", sc.AppsWidth+sc.UseFrame, sc.TitleHeight, 2*sc.AppsWidth+2*sc.UseFrame, sc.Height-(sc.LinesOfHelp+1)); err != nil {
		if err != gocui.ErrUnknownView {
			return v, err
		}
		v.Highlight = true
		v.Title = "Placeholders"
		fmt.Fprintln(v, "Loading ...")
		return v, nil
	}

	return nil, nil
}

func layoutValues(g *gocui.Gui, sc ScreenConf) (*gocui.View, error) {
	if v, err := g.SetView("values", 2*sc.AppsWidth+3*sc.UseFrame, sc.TitleHeight, sc.Width-1, sc.Height-(sc.LinesOfHelp+1)); err != nil {
		if err != gocui.ErrUnknownView {
			return v, err
		}
		v.Highlight = true
		v.Title = "Values"
		fmt.Fprintln(v, "Loading ...")
		return v, nil
	}
	return nil, nil
}

func loadPlaceholders(g *gocui.Gui, v *gocui.View) error {
	sc := getScreenConf(g)

	_, y := v.Cursor()

	if err := g.DeleteView("placeholders"); err != nil {
		return err
	}

	if err := g.DeleteView("values"); err != nil {
		return err
	}

	vp, _ := layoutPlaceHolders(g, sc)
	vv, _ := layoutValues(g, sc)
	vp.SetCursor(0, 0)
	vp.Clear()
	vv.Clear()
	appName, _ := v.Line(y)
	pls, _ := api.GetPlaceholders(appName)

	for _, pl := range pls.Placeholders {
		fmt.Fprintln(vp, pl.PropertyName)
		fmt.Fprintln(vv, pl.PropertyValue+" ")
	}

	if err := g.SetCurrentView("placeholders"); err != nil {
		return err
	}
	return nil
}

func editValue(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("dialog", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
		if err := g.SetCurrentView("dialog"); err != nil {
			return err
		}
	}
	return nil
}

func saveValue(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("dialog"); err != nil {
		return err
	}

	if err := g.SetCurrentView("placeholders"); err != nil {
		return err
	}
	return nil
}

func Execute() {

	api = common.NewAPI("http://localhost:4747")
	//api.Debug(true)

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(layout)
	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}
	g.SelBgColor = gocui.ColorGreen
	g.SelFgColor = gocui.ColorBlack
	g.Cursor = true

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
