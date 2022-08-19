package screens

import (
	"fmt"
	"github.com/novabankapp/ussdcommon/utils"
)

// Menu for USSD
type Menu struct {
	Header, Footer string
	Items          []*menuItem
	ZeroItem       *menuItem
}

type menuItem struct {
	Name  string
	Route utils.Route
}

// NewMenu creates a new Menu
func NewMenu() *Menu {
	return &Menu{
		Items: make([]*menuItem, 0),
	}
}

// Add to USSD menu.
func (m *Menu) Add(name, ctrl, action string) *Menu {
	item := &menuItem{Name: name,
		Route: utils.Route{Ctrl: ctrl, Action: action},
	}
	m.Items = append(m.Items, item)
	return m
}

// AddZero adds an item at the bottom of USSD menu.
// This item always routes to a choice of "0".
func (m *Menu) AddZero(name, ctrl, action string) *Menu {
	m.ZeroItem = &menuItem{
		Name:  name,
		Route: utils.Route{Ctrl: ctrl, Action: action},
	}
	return m
}

// render USSD menu.
func (m Menu) render() string {
	msg := utils.StrEmpty
	if m.Header != utils.StrEmpty {
		msg += m.Header + utils.StrNewLine
	}
	for i, item := range m.Items {
		msg += fmt.Sprintf("%d. %v"+utils.StrNewLine, i+1, item.Name)
	}
	if m.ZeroItem != nil {
		msg += "0. " + m.ZeroItem.Name + utils.StrNewLine
	}
	msg += m.Footer
	return msg
}
