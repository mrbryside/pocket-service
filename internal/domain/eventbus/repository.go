package eventbus

//go:generate mockgen -source=./repository.go -destination=../../core/generated/mockgen/eventbus_domain/repository.go -package=mockEventBusDomain
type Repository interface {
	InsertEvents(e *EventBus) error
}
