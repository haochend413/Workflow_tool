package gui

import (
	"fmt"

	"github.com/haochend413/mantis/controllers"
	"github.com/jroimartin/gocui"
)

// Define layout for all views;
func (gui *Gui) layout(g *gocui.Gui) error {
	//init template
	if gui.first_init_check {
		gui.windows = gui.CreateWindowTemplates()
	}

	//here, only check logic
	// init views
	for _, w := range gui.windows {
		// fmt.Fprint(os.Stdout, w.Name)
		if !w.OnDisplay {
			// Don't show views that are off
			continue
		}
		//here it set up view prepare;
		// Only initialize if the view was just created

		v, err := gui.prepareView(w)
		// //Dont know why here, but might be useful
		// if !w.OnDisplay {
		// 	g.DeleteView(w.Name)
		// }
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}

		//check: init if only first created
		if err == gocui.ErrUnknownView {
			//view config
			v.Title = w.Title
			w.View = v
			if w.Editable {
				v.Editable = true
			}
			if w.Scroll {
				v.Autoscroll = true
			}
			if w.Cursor {
				controllers.CursorOn(g, v)
			}
		}

		//view-specific logic here
		if w.Name == "note-history" {
			nh := w.View
			nh.Clear()
			//display history
			for _, note := range DB_Data.NoteDBData {
				fmt.Fprintln(nh, note.Content)
			}
			// return nil
		}

		// if w.Name == "note-detail" {
		// 	nh, e := g.View("note-detail")
		// 	nh.Clear()
		// 	// fmt.Fprint(os.Stdout, "hihi, \n hihi, \n hihi")
		// }

	}

	//setstartview
	if gui.first_init_check {
		g.SetCurrentView("note")
		gui.first_init_check = false
	}

	return nil
}
