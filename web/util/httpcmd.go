package util

type HttpCmd struct {
	Command string
	Action interface{}
}

func NewHttpCmd(command string, action interface{}) HttpCmd{
	return HttpCmd{Command: command, Action : action}
}