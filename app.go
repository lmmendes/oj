package main

type App struct {
	Parser Parser
}

// Returns a new app wrapper
func NewApp(dr Parser) *App {
	return &App{
		Parser: dr,
	}
}

// parse
func (t *App) parse(c string) string {
	return t.Parser.Parse(c)
}

// set parser
func (t *App) setParser(dr Parser) {
	t.Parser = dr
}
