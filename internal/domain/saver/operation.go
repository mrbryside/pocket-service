package saver

//go:generate mockgen -source=./operation.go -destination=../../core/generated/mockgen/saver_domain/operation.go -package=mockSaverDomain
type Operation interface {
	InsertPocket(s *Saver) error
}
