package handlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/att-innovate/rancher-mesos-scheduler/tasks"
	"github.com/att-innovate/go-machine-service/events"
	"github.com/rancher/go-rancher/client"
)

func MesosScheduleCreate(event *events.Event, apiClient *client.RancherClient) (err error) {
	log.WithFields(log.Fields{
		"resourceId": event.ResourceId,
		"eventId":    event.Id,
	}).Info("Creating Machine")

	task := &tasks.Task{
		Id:   event.Id,
		Name: event.ResourceId,
		CPU:  1,
		Mem:  1024,
	}

	tasks.AddTask(task)

	reply := newReply(event)
	return publishReply(reply, apiClient)
}
