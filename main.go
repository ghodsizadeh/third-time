package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	stopwatch stopwatch.Model
	keymap    keymap
	timer     timer.Model
	resting   bool
	help      help.Model
	quitting  bool
	fraction  int
}

const break_length = 5

type keymap struct {
	start      key.Binding
	stop       key.Binding
	start_rest key.Binding
	reset      key.Binding
	quit       key.Binding
}

func (m model) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m model) View() string {
	// Note: you could further customize the time output by getting the
	// duration from m.stopwatch.Elapsed(), which returns a time.Duration, and
	// skip m.stopwatch.View() altogether.
	if m.resting {
		s := m.timer.View() + "\n"
		if !m.quitting {
			s = "Resting: " + s
			s += m.helpView()
		} else {
			s = "Resting is over\n"
			// add what is time now
			s += "Now " + time.Now().Format("15:04:05") + "\n"
			s += m.helpView()
		}
		return s
	}

	// return "Resting\n"
	s := m.stopwatch.View() + "\n"
	if !m.quitting {
		s = "Elapsed: " + s
		s += m.helpView()
	}
	return s
}

func (m model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.start,
		m.keymap.stop,
		m.keymap.start_rest,
		m.keymap.reset,
		m.keymap.quit,
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keymap.quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keymap.reset):
			return m, m.stopwatch.Reset()
		case key.Matches(msg, m.keymap.start_rest):
			elapsedTime := m.stopwatch.Elapsed()
			restTime := elapsedTime / time.Duration(m.fraction)
			m.timer = timer.NewWithInterval(restTime, time.Millisecond)
			m.resting = true
			m.timer.Init()
			return m, m.timer.Start()
		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			m.keymap.stop.SetEnabled(!m.stopwatch.Running())
			m.keymap.start.SetEnabled(m.stopwatch.Running())
			return m, m.stopwatch.Toggle()
		}
	}
	var cmd tea.Cmd
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, cmd
}

func main() {
	fraction := flag.Int("f", 3, "How you much you want to rest after each time")
	flag.Parse()

	m := model{
		stopwatch: stopwatch.NewWithInterval(time.Millisecond),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			start_rest: key.NewBinding(
				key.WithKeys("t"),
				key.WithHelp("t", "start rest"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			quit: key.NewBinding(
				key.WithKeys("ctrl+c", "q"),
				key.WithHelp("q", "quit"),
			),
		},
		help:     help.New(),
		fraction: *fraction,
	}

	m.keymap.start.SetEnabled(false)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Oh no, it didn't work:", err)
		os.Exit(1)
	}
}
