package ecs

type CancelCommand struct {
	cancel bool
}

func NewCancelCommand(cancel bool) *CancelCommand {
	return &CancelCommand{
		cancel: cancel,
	}
}

func (e *CancelCommand) Write(w *World) {
	if e.cancel {
		w.Cancel()
	}
}
