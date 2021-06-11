package executor

type TaskTp int

const (
	TaskTpFoo TaskTp = iota + 1
	TaskTpBar
)

type TaskStatus int

const (
	TaskStatusWait TaskStatus = iota + 1
	TaskStatusExc
	TaskStatusFinish
)

type TaskData struct {
	Data string
	Tp   TaskTp
}

type Executor interface {
	Exc(data TaskData) error
	GetStatus(taskId string) (TaskStatus, error)
}
