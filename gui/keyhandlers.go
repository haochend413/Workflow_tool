package gui

import (
	"fmt"
	"os"

	"github.com/haochend413/mantis/controllers"
	"github.com/jroimartin/gocui"
)

func (gui *Gui) HandleAppQuit(g *gocui.Gui, v *gocui.View) error {
	return gui.AppQuit()

}

func (gui *Gui) HandleDataUpdate(g *gocui.Gui, v *gocui.View) error {
	// fmt.Fprintln(v, "test")
	return gui.DBManager.RefreshAll(DB_Data)
	// return nil
}

// View setup;
func (gui *Gui) HandleNoteDisplay(g *gocui.Gui, v *gocui.View) error {
	// fmt.Fprint(os.Stdout, gui.windows[0].Name)
	return ToggleWindowDisplay(gui.windows[0], gui.g)
}

// View setup;
func (gui *Gui) HandleNoteHistoryDisplay(g *gocui.Gui, v *gocui.View) error {
	// fmt.Fprint(os.Stdout, gui.windows[0].Name)
	return ToggleWindowDisplay(gui.windows[1], gui.g)
}

// View setup;
func (gui *Gui) HandleCmdDisplay(g *gocui.Gui, v *gocui.View) error {
	// fmt.Fprint(os.Stdout, gui.windows[0].Name)
	return ToggleWindowDisplay(gui.windows[2], gui.g)
}

// Move Cursor;
func (gui *Gui) HandleNoteCursorMove(direction string) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		switch direction {
		case "up":
			err := controllers.CursorUp(gui.windows[1].View)
			_, Current_Note_Index = gui.windows[1].View.Cursor()
			UpdateSelectedNote(Current_Note_Index, gui.g)
			fmt.Fprintln(os.Stdout, Current_Note_Index)
			return err
		case "down":
			err := controllers.CursorDown(gui.windows[1].View)
			_, Current_Note_Index = gui.windows[1].View.Cursor()
			UpdateSelectedNote(Current_Note_Index, gui.g)
			fmt.Fprintln(os.Stdout, Current_Note_Index)
			return err
		case "left":
			return controllers.CursorLeft(gui.windows[1].View)
		case "right":
			return controllers.CursorRight(gui.windows[1].View)
		default:
			return nil
		}
	}
}

/*
Note view
*/

// View setup;
func (gui *Gui) HandleSendNote(g *gocui.Gui, v *gocui.View) error {
	// fmt.Fprint(os.Stdout, gui.windows[0].Name)
	return gui.SendNote()
}
