package core

const (
	STATE_WAITTING = iota
	STATE_RUNNING
	STATE_FAILED
	STATE_FINISHED
	STATE_KILLING
)

const (
	TASK_LONG_TERM  = "LONG_TERM"
	TASK_SHORT_TERM = "SHORT_TERM"
)
