package domain

import (
	"github.com/stackus/edat/es"
)

type AccountRepository interface {
	es.AggregateRepository
}
