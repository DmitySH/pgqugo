package entity

import "time"

type GetWaitingTasksParams struct {
	KindID           int16
	BatchSize        int
	AttemptsInterval time.Duration
}

type DeleteTerminalTasksParams struct {
	KindID int16
	Limit  int
	After  time.Duration
}
