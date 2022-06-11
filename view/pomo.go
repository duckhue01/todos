package view

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/duckhue01/todos/model"
)

var (
	red   = lipgloss.NewStyle().Foreground(lipgloss.Color("#d30d0d"))
	green = lipgloss.NewStyle().Foreground(lipgloss.Color("#27a300"))
	gray  = lipgloss.NewStyle().Foreground(lipgloss.Color("#f2e7c9"))

	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)

	timerStyle = lipgloss.NewStyle().Width(20).Align(lipgloss.Center).Border(lipgloss.NormalBorder())

	timerContainerStyle = lipgloss.NewStyle().Width(30).Height(5).Border(lipgloss.NormalBorder())
)

type timerModel struct {
	timer       timer.Model
	quitting    bool
	help        help.Model
	state       string
	pomoConfig  *model.PomoConfig
	round       int
	start       string
	timerKeymap timerKeymap
}

type timerKeymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	// quit  key.Binding
}

func NewModel(t timer.Model, pomoConfig *model.PomoConfig, state string, startTime string) timerModel {
	return timerModel{
		timer:      t,
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

func (m timerModel) Init() tea.Cmd {

	return m.timer.Init()
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// window size event
	case tea.WindowSizeMsg:
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
		case "s":
			return m, m.timer.Toggle()
		}
	}

	var cmd tea.Cmd
	return m, cmd

}

func (m timerModel) View() string {

	if m.quitting {
		return quitTextStyle.Render("Well done!!!")
	}

	var timerContainerView string

	startTime := gray.Render(fmt.Sprintf("Start time: %v", m.start))

	if m.state == "focus" {
		timerView := timerStyle.Render(red.Render(m.timer.View()))
		timerContainerView = timerContainerStyle.Render(fmt.Sprintf("%v\n%v\nState: %v - Round: %v ",
			startTime, timerView,
			red.Render(m.state),
			gray.Render(fmt.Sprint(m.round))))

	} else {
		timerView := timerStyle.Render(green.Render(m.timer.View()))
		timerContainerView = timerContainerStyle.Render(fmt.Sprintf("%v\n%v\nState: %v - Round: %v ",
			startTime,
			timerView,
			green.Render(m.state),
			gray.Render(fmt.Sprint(m.round))))

	}

	return timerContainerView
}

func StartPomo(pomoConfig *model.PomoConfig, state string, startTime string) {

	if err := tea.NewProgram(NewModel(timer.New(pomoConfig.Pomo*time.Minute),
		pomoConfig,
		state,
		startTime)).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
