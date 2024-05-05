package event

import "github.com/gookit/event"

type BackupEvent struct {
	event.BasicEvent
	BackupData []map[string]any
}

func (e *BackupEvent) SetBackupData(name string, model any) {
	e.BackupData = append(e.BackupData, map[string]any{
		"name":  name,
		"model": model,
	})
}

func (e *BackupEvent) GetBackupData() []map[string]any {
	return e.BackupData
}
