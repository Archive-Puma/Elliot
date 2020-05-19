package cli

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/cosasdepuma/elliot/app/env"
)

// Modal TODO: Doc
var Modal = struct {
	Active  bool
	Current string
}{}

func displayModal(gui *gocui.Gui) error {
	if Modal.Active {
		w, h := gui.Size()

		env.Params.Arguments = nil
		modal := ModalViews[Modal.Current]
		if popup, err := gui.SetView("Modal", w/3, (h/2)-1, 2*w/3, h/2+1, 0); err != nil {
			if !gocui.IsUnknownView(err) {
				return err
			}
			popup.Clear()
			popup.Wrap = true
			popup.Title = modal.name
			popup.Frame = true
			popup.Editable = modal.input

			if Current >= 0 {
				Current -= len(MainViews)
			}
			if _, err := setCurrentViewOnTop(gui, "Modal"); err != nil {
				return err
			}

			if err := gui.SetKeybinding("Modal", gocui.KeyEnter, gocui.ModNone, modalHandler); err != nil {
				return err
			}
		}
	}

	return nil
}

func disableModal(gui *gocui.Gui) error {
	gui.DeleteView("Modal")
	Modal.Active = false
	if Current < 0 {
		Current += len(MainViews)
	}
	return nil
}

func modalHandler(gui *gocui.Gui, view *gocui.View) error {
	modalView, err := gui.View("Modal")
	if err != nil {
		return err
	}

	if ModalViews[Modal.Current].input {
		content := modalView.Buffer()
		if len(content) == 0 {
			disableModal(gui)
			return nil
		}
		env.Params.Arguments = []interface{}{content}
	}
	fmt.Println("Params", env.Params.Arguments)
	disableModal(gui)
	runModule(gui, view)
	return nil
}
