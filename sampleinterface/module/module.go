package module

type Module struct {
}

func (m *Module) Hello() string {
	return "Hello World"
}

func (m *Module) Bye() string {
	return "Bye World"
}

func (m *Module) Name() string {
	return "Module"
}
