package gui

import (
	"fmt"
	"strings"

	"github.com/haochend413/mantis/controllers"
	"github.com/haochend413/mantis/models"
	"github.com/jroimartin/gocui"
)

// Control window display
// make sure that view change only happens here for now
// it starts with nothing
var VIEW_SWITCH_HISTORY = []string{""}

// Usually works. Might be buggy
func ToggleWindowDisplay(w *models.Window, g *gocui.Gui) error {
	w.OnDisplay = !w.OnDisplay

	// Safe GUI update
	g.Update(func(g *gocui.Gui) error {
		VIEW_SWITCH_HISTORY = append(VIEW_SWITCH_HISTORY, g.CurrentView().Name())
		if w.OnDisplay {
			// Create/show the view
			v, err := g.SetView(w.Name, w.X0, w.Y0, w.X1, w.Y1)
			if err != nil && err != gocui.ErrUnknownView {
				return err
			}
			v.Title = w.Title
			w.View = v
			g.SetCurrentView(w.Name)
		} else {
			//manage focus

			if len(VIEW_SWITCH_HISTORY) > 1 {
				VIEW_SWITCH_HISTORY = VIEW_SWITCH_HISTORY[:len(VIEW_SWITCH_HISTORY)-1]
			}
			// Delete/hide the view
			err := g.DeleteView(w.Name)
			g.SetCurrentView(VIEW_SWITCH_HISTORY[len(VIEW_SWITCH_HISTORY)-1])
			if err != nil && err != gocui.ErrUnknownView {
				return err
			}
			w.View = nil
		}
		return nil
	})
	return nil
}

// fetch the current content input of that view;
func FetchContent(w *models.Window, g *gocui.Gui) string {
	return strings.TrimSpace(w.View.Buffer())
}

// store note to DB_Data
func (gui *Gui) SendNote() error {
	content := FetchContent(gui.windows[0], gui.G())
	if content == "" {
		return nil
	}
	note := &models.Note{Content: content}

	DB_Data.NoteDBData = append(DB_Data.NoteDBData, note)

	gui.g.CurrentView().Clear()
	controllers.CursorOn(gui.g, gui.g.CurrentView())
	return nil
}

// Note Detail Index
var Current_Note_Index = 0

// Display Note Content Detail
func UpdateSelectedNote(index int, g *gocui.Gui) error {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("note-detail")
		if err != nil {
			return nil
		}
		v.Clear()
		v.Wrap = true
		fmt.Fprint(v, DB_Data.NoteDBData[index].Content)
		return nil
	})
	return nil
}
