package multichoose

type Option func(*Model)

func WithHelp(show bool) Option {
	return func(m *Model) {
		m.showHelp = show
	}
}

func WithTheme(theme Theme) Option {
	return func(m *Model) {
		m.theme = theme
	}
}

func WithKeyMap(keyMap KeyMap) Option {
	return func(m *Model) {
		m.keys = keyMap
	}
}