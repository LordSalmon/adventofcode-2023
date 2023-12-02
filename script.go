package main

type ScriptHolder interface {
	GetInput() string
	GetLines() []string
	ReadInput(path string)
}

type Script struct {
	InputPath string
}

func (s *Script) GetInput() string {
	return s.InputPath
}
