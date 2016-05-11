package cms

type RunStatus int

const (
	NotStarted = RunStatus(0) //未启动
	Starting   = RunStatus(1) //正在操作中
	Started    = RunStatus(2) //启动
)
