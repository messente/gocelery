package gocelery

import "time"

type CeleryTaskOption func(celeryTask *TaskMessage) error

func WithETA(eta string) CeleryTaskOption {
	return func(celeryTask *TaskMessage) error {
		celeryTask.ETA = &eta
		return nil
	}
}

func WithCountdown(countdown int) CeleryTaskOption {
	eta := time.Now().Add(time.Duration(countdown) * time.Second).Format(time.RFC3339)
	return func(celeryTask *TaskMessage) error {
		celeryTask.ETA = &eta
		return nil
	}
}
