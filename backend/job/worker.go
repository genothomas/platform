package job

import (
	"github.com/syncloud/platform/backup"
	"log"
)

type Worker struct {
	backup *backup.Backup
	master *Master
}

func NewWorker(master *Master, backup *backup.Backup) *Worker {
	return &Worker{
		backup: backup,
		master: master,
	}
}

func (worker *Worker) Start() {
	go func() {
		for {
			job := <-worker.master.JobQueue()
			switch jobtype := job.(type) {
			case JobBackupCreate:
				v := job.(JobBackupCreate)
				worker.backup.Create(v.app, v.file)
			default:
				log.Println("not supported job type", jobtype)
			}
		}
	}()
}