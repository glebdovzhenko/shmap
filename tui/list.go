package tui

import (
	"fmt"
    tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	shdb "github.com/glebdovzhenko/shmap/database"
)

func (m TuiModel) updateList(msg tea.Msg) (TuiModel, tea.Cmd) {
    var cmd tea.Cmd
    m.List, cmd = m.List.Update(msg)	
	return m, cmd
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func InitTuiModelList(md *TuiModel, tb *[]shdb.DBTableData) *TuiModel {

	items := make([]list.Item, len(*tb))
	for ii, table := range *tb {
		items[ii] = list.Item(item{title: table.Name, desc: fmt.Sprintf("%d columns %d rows", len(*table.Rows), len(*table.ColumnNames))})
	}

	md.List = list.New(items, list.NewDefaultDelegate(), 20, 20)
	md.List.Title = "Tables"
	md.List.SetShowStatusBar(false)
	md.List.SetFilteringEnabled(false)
	return md
}

