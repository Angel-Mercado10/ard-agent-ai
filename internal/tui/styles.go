package tui

import "github.com/charmbracelet/lipgloss"

const (
	colorPrimary = "#7C3AED"
	colorSuccess = "#10B981"
	colorWarning = "#F59E0B"
	colorError   = "#EF4444"
	colorMuted   = "#6B7280"
	colorBorder  = "#374151"
)

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorPrimary)).
			PaddingLeft(1).
			PaddingRight(1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted))

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(colorBorder)).
			Padding(1, 2)

	BoxPrimaryStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(colorPrimary)).
			Padding(1, 3)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorSuccess)).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorError))

	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWarning))

	MutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted))

	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPrimary)).
			Bold(true)

	CheckboxStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPrimary))

	KeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			Italic(true)

	SectionLabelStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(colorMuted))

	DividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBorder))
)
