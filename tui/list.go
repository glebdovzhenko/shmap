package tui

import (
	"fmt"
    tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
)

func (m TuiModel) updateList(msg tea.Msg) (TuiModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
        case "enter":
            cmds = append(cmds, tea.Batch(emitSwitchTableMsg(m.List.Index())))
		}
	}

    m.List, cmd = m.List.Update(msg)	
    cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func InitTuiModelList(md *TuiModel) *TuiModel {

	items := make([]list.Item, len(*md.DBData))
    var(
        title string
        n_rows, n_cols int
    )
	for ii, _ := range items {
        title, _ = md.DBData.Name(ii)
        n_cols, _ = md.DBData.ColumnsLen(ii)
        n_rows, _ = md.DBData.RowsLen(ii)
		items[ii] = list.Item(item{
            title: title, 
            desc: fmt.Sprintf("%d columns %d rows", n_cols, n_rows),
        })
	}

	md.List = list.New(items, list.NewDefaultDelegate(), 20, 20)
	md.List.Title = "Tables"
	md.List.SetShowStatusBar(false)
	md.List.SetFilteringEnabled(false)
	return md
}

