package application

import (
	"encoding/json"
	"fmt"
	crosscutting "go-starter-eventhandler/src/crosscutting/error"
	"go-starter-eventhandler/src/crosscutting/messaging"
	"go-starter-eventhandler/src/domain/event"
	"go-starter-eventhandler/src/domain/model"

	"time"
)

//BusinessService : Business service
func BusinessService(message string) {

	fmt.Println(model.DomainModel{Propertie: message})
	raisedDomainEvent := event.DomainEvent{DomainModel: model.DomainModel{Propertie: message}, PublicationDateTimeUnix: time.UnixDate}
	body, err := json.Marshal(raisedDomainEvent)
	rabbitMqProvider := messaging.RabbitMqProvider{}
	rabbitMqProvider.PublishOneMessage("raisedDomainEvent", []byte(body))
	crosscutting.RaiseError("Failed to publish a message", err)
}
