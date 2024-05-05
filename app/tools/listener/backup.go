package listener

import (
	toolsEvent "dux-project/app/tools/event"
	"dux-project/app/tools/models"
	"github.com/gookit/event"
)

// BackupListener @Listener(name = "tools.backup")
func BackupListener(e event.Event) error {
	e.(*toolsEvent.BackupEvent).SetBackupData("file_area", models.ToolsArea{})
	e.(*toolsEvent.BackupEvent).SetBackupData("file_dir", models.ToolsFileDir{})
	return nil
}
