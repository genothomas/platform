package main

import (
	"github.com/syncloud/platform/event"
	"log"
	"os"

	"github.com/syncloud/platform/backup"
	"github.com/syncloud/platform/installer"
	"github.com/syncloud/platform/job"
	"github.com/syncloud/platform/rest"
	"github.com/syncloud/platform/storage"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("usage: ", os.Args[0], "/path.sock")
		return
	}

	os.Remove(os.Args[1])
	Backend().Start(os.Args[1])

}

func Backend() *rest.Backend {
	master := job.NewMaster()
	backup := backup.NewDefault()
	eventTrigger := event.New()
	installer := installer.New()
	storage := storage.New()
	worker := job.NewWorker(master, backup, installer, storage)

	return rest.NewBackend(master, backup, eventTrigger, worker)

}
