package note

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/haochend413/mantis/defs"
)

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
	// m.ti.Width = width - 4
	m.ti.SetHeight(height - 1)
	m.ti.SetWidth(width - 1)

}

// sendnote
// type NoteSendMsg = *defs.Note

func (m *Model) SendNoteCmd() tea.Cmd {
	content := m.ti.Value()
	if content == "" {
		return nil
	}
	m.ti.Reset()
	return func() tea.Msg {
		return &defs.Note{
			Content: content,
		}
	}
}

func (m *Model) SendTopicCmd() tea.Cmd {
	content := m.ti.Value()
	if content == "" {
		return nil
	}
	m.ti.Reset()
	return func() tea.Msg {
		return &defs.Topic{
			Topic: content,
		}
	}
}

// func (m *Model) FocusView() {

// }
