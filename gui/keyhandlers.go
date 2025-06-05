package gui

import (
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

// View switch
// Should not go to read-only views
func (gui *Gui) HandleViewLoop(g *gocui.Gui, v *gocui.View) error {
	switch v.Name() {
	case "note":
		g.SetCurrentView("note-history")
		return nil
	case "note-history":
		g.SetCurrentView("note")
		return nil
	default:
		return nil
	}
}

// Move Cursor;
func (gui *Gui) HandleNoteCursorMove(direction string) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		switch direction {
		case "up":
			err := controllers.CursorUp(gui.windows[1].View)
			_, Current_Note_Index = gui.windows[1].View.Cursor()
			UpdateSelectedNote(Current_Note_Index, gui.g)
			// fmt.Fprintln(os.Stdout, Current_Note_Index)
			return err
		case "down":
			err := controllers.CursorDown(gui.windows[1].View)
			_, Current_Note_Index = gui.windows[1].View.Cursor()
			UpdateSelectedNote(Current_Note_Index, gui.g)
			// fmt.Fprintln(os.Stdout, Current_Note_Index)
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

// Send Note.
// Update history & detail demo.
func (gui *Gui) HandleSendNote(g *gocui.Gui, v *gocui.View) error {
	// fmt.Fprint(os.Stdout, gui.windows[0].Name)
	//update the note-history view and the note-detail view
	//current position
	px, _ := gui.windows[1].View.Cursor()
	//here, py should not be lower than the last line; -2 : trimmed empty line

	if err := gui.windows[1].View.SetCursor(px, len(gui.windows[1].View.BufferLines())-1); err != nil {
		return err
	}
	_, Current_Note_Index = gui.windows[1].View.Cursor()
	UpdateSelectedNote(Current_Note_Index, gui.g)
	return gui.SendNote()
}

func (gui *Gui) HandleSwitchLine(g *gocui.Gui, v *gocui.View) error {
	// fmt.Fprint(os.Stdout, gui.windows[0].Name)
	v.EditNewLine()
	return nil
}
