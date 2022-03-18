package view

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/duckhue01/todos/models"
)

const listHeight = 10

var (
	// list styles
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("200"))
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	doneStyle         = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("100"))
	unDoneStyle       = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("30"))

	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(2)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(2).PaddingBottom(1)
	quitTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4)

	timerContainerStyle = lipgloss.NewStyle().Width(30).Height(10)
	timerStyle          = lipgloss.NewStyle().Width(20).Align(lipgloss.Center).Border(lipgloss.NormalBorder())
	listStyle           = lipgloss.NewStyle().Height(10).Align(lipgloss.Center)
)

type Item models.Todo

func (i Item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Title)

	fn := itemStyle.Render

	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	} else {
		if !i.IsDone {
			fn = func(s string) string {
				return unDoneStyle.Render(s)
			}
		} else {
			fn = func(s string) string {
				return doneStyle.Render(s)
			}
		}

	}

	fmt.Fprintf(w, fn(str))
}

type model struct {
	list     list.Model
	timer    timer.Model
	choice   Item
	quitting bool
	keymap   keymap
	help     help.Model
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

func NewModel(l list.Model, t timer.Model) model {
	return model{
		list:     l,
		choice:   Item{},
		quitting: false,
		timer:    t,
	}
}

func (m model) Init() tea.Cmd {

	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// window size event
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	// timer event
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.keymap.stop.SetEnabled(m.timer.Running())
		m.keymap.start.SetEnabled(!m.timer.Running())
		return m, cmd

	case timer.TimeoutMsg:
		m.quitting = true
		return m, tea.Quit

	// tea key msg
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = i
			}
			return m, nil
		}

	}

	var cmd tea.Cmd
	m.list, _ = m.list.Update(msg)
	return m, cmd

}

func (m model) View() string {

	if m.quitting {
		return quitTextStyle.Render("you have had a good day")
	}

	timerView := timerStyle.Render(m.timer.View())

	timerContainerView := timerContainerStyle.Render(fmt.Sprintf("start time: 1\n%v\nstate: focus - round: 1 ", timerView))
	listView := listStyle.Render(m.list.View())

	return lipgloss.JoinHorizontal(0, timerContainerView, listView)
}

func (m model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.start,
		m.keymap.stop,
		m.keymap.reset,
		m.keymap.quit,
	})
}

func Start() {
	items := []list.Item{
		Item{
			Title:  "this is some thing that different",
			Tag:    "asdasd",
			IsDone: false,
			Id:     1,
		},
		Item{
			Title:  "learn a bit english",
			Tag:    "asdasd",
			IsDone: true,
			Id:     1,
		},
		Item{
			Title:  "asdasd",
			Tag:    "asdasd",
			IsDone: true,
			Id:     1,
		},
		Item{
			Title:  "this is some thing that different",
			Tag:    "asdasd",
			IsDone: false,
			Id:     1,
		},
		Item{
			Title:  "learn a bit english",
			Tag:    "asdasd",
			IsDone: true,
			Id:     1,
		},
		Item{
			Title:  "asdasd",
			Tag:    "asdasd",
			IsDone: true,
			Id:     1,
		},

		Item{
			Title:  "this is some thing that different",
			Tag:    "asdasd",
			IsDone: false,
			Id:     1,
		},
		Item{
			Title:  "learn a bit english",
			Tag:    "asdasd",
			IsDone: true,
			Id:     1,
		},
		Item{
			Title:  "asdasd",
			Tag:    "asdasd",
			IsDone: true,
			Id:     1,
		}, Item{
			Title:  "this is some thing that different",
			Tag:    "asdasd",
			IsDone: false,
			Id:     1,
		},
		Item{
			Title:  "learn a bit english",
			Tag:    "asdasd",
			IsDone: true,
			Id:     1,
		},
		Item{
			Title:  "asdasd",
			Tag:    "asdasd",
			IsDone: true,
			Id:     1,
		},
	}

	const defaultWidth = 20

	// new list model
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "My today tasks:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	// new timer model
	s := timer.New(10 * time.Minute)

	if err := tea.NewProgram(NewModel(l, s)).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
