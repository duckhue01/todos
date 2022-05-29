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
	red   = lipgloss.NewStyle().Foreground(lipgloss.Color("#d30d0d"))
	green = lipgloss.NewStyle().Foreground(lipgloss.Color("#27a300"))
	gray  = lipgloss.NewStyle().Foreground(lipgloss.Color("#f2e7c9"))

	// list styles
	titleStyle    = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("200"))
	itemStyle     = lipgloss.NewStyle().PaddingLeft(2)
	doneStyle     = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#27a300"))
	unDoneStyle   = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#d30d0d"))
	selectedStyle = lipgloss.NewStyle().PaddingLeft(4).Foreground(lipgloss.Color("#ffdd00"))

	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(2)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(2).PaddingBottom(1)
	quitTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4)

	timerContainerStyle = lipgloss.NewStyle().Width(30).Height(10)

	timerStyle = lipgloss.NewStyle().Width(20).Align(lipgloss.Center).Border(lipgloss.NormalBorder())
	listStyle  = lipgloss.NewStyle().Height(10).Align(lipgloss.Center)
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

	str := i.Title

	fn := itemStyle.Render

	if index == m.Index() {

		if !i.IsDone {
			fn = func(s string) string {
				return selectedStyle.Render("[ ]" + s)
			}
		} else {
			fn = func(s string) string {
				return selectedStyle.Render("[x]" + s)
			}
		}

	} else {
		if !i.IsDone {
			fn = func(s string) string {
				return unDoneStyle.Render("[ ]" + s)
			}
		} else {
			fn = func(s string) string {
				return doneStyle.Render("[x]" + s)
			}
		}

	}

	fmt.Fprintf(w, fn(str))
}

type model struct {
	list        list.Model
	timer       timer.Model
	choice      Item
	quitting    bool
	help        help.Model
	state       string
	pomoConfig  *models.PomoConfig
	round       int
	start       string
	timerKeymap timerKeymap
}

type timerKeymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
}

func NewModel(l list.Model, t timer.Model, pomoConfig *models.PomoConfig, state string, startTime string) model {
	return model{
		list:       l,
		timer:      t,
		choice:     Item{},
		quitting:   false,
		help:       help.Model{},
		state:      state,
		pomoConfig: pomoConfig,
		round:      1,
		start:      startTime,
		timerKeymap: timerKeymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
		},
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
		m.timerKeymap.stop.SetEnabled(m.timer.Running())
		m.timerKeymap.start.SetEnabled(!m.timer.Running())
		return m, cmd

	case timer.TimeoutMsg:
		if m.state == "focus" && m.round < m.pomoConfig.Interval {
			m.state = "break"
			m.timer = timer.New(m.pomoConfig.Break * time.Second)
			cmd := m.timer.Start()
			return m, cmd
		}

		if m.state == "break" && m.round < m.pomoConfig.Interval {
			m.state = "focus"
			m.round++
			m.timer = timer.New(m.pomoConfig.Pomo * time.Second)
			cmd := m.timer.Start()
			return m, cmd
		}

		if m.round == m.pomoConfig.Interval {
			m.quitting = true
			return m, tea.Quit
		}

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

		case "s":
			return m, m.timer.Toggle()
		}
	}

	var cmd tea.Cmd
	m.list, _ = m.list.Update(msg)
	return m, cmd

}

func (m model) View() string {

	if m.quitting {
		return quitTextStyle.Render("Well done!!!")
	}

	var timerContainerView string
	var listView string

	startTime := gray.Render(fmt.Sprintf("Start time: %v", m.start))

	listView = listStyle.Render(m.list.View())

	if m.state == "focus" {
		timerView := timerStyle.Render(red.Render(m.timer.View()))
		timerContainerView = timerContainerStyle.Render(fmt.Sprintf("%v\n%v\nstate: %v - round: %v ", startTime, timerView, red.Render(m.state), red.Render(fmt.Sprint(m.round))))

	} else {
		timerView := timerStyle.Render(green.Render(m.timer.View()))
		timerContainerView = timerContainerStyle.Render(fmt.Sprintf("%v\n%v\nstate: %v - round: %v ", startTime, timerView, green.Render(m.state), green.Render(fmt.Sprint(m.round))))

	}

	return lipgloss.JoinHorizontal(0, timerContainerView, listView)
}

func (m model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.timerKeymap.start,
		m.timerKeymap.stop,
		m.timerKeymap.reset,
	})
}

func StartPomo(pomoConfig *models.PomoConfig, state string, startTime string) {
	items := []list.Item{
		Item{
			Title:  "this is some thing that different",
			Tag:    "asdasd",
			IsDone: true,
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
			IsDone: false,
			Id:     1,
		},
		Item{
			Title:  "this is some thing that different",
			Tag:    "asdasd",
			IsDone: false,
			Id:     1,
		},
	}

	const defaultWidth = 20

	// new list model
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = gray.Render("My today tasks:")
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	// new timer model
	s := timer.New(pomoConfig.Pomo * time.Minute)

	if err := tea.NewProgram(NewModel(l, s, pomoConfig, state, startTime)).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
