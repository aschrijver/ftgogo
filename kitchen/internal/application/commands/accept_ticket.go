package commands

import (
	"context"
	"time"

	"github.com/stackus/ftgogo/kitchen/internal/adapters"
	"github.com/stackus/ftgogo/kitchen/internal/domain"
)

type AcceptTicket struct {
	TicketID string
	ReadyBy  time.Time
}

type AcceptTicketHandler struct {
	repo adapters.TicketRepository
}

func NewAcceptTicketHandler(repo adapters.TicketRepository) AcceptTicketHandler {
	return AcceptTicketHandler{
		repo: repo,
	}
}

func (h AcceptTicketHandler) Handle(ctx context.Context, cmd AcceptTicket) error {
	_, err := h.repo.Update(ctx, cmd.TicketID, &domain.AcceptTicket{ReadyBy: cmd.ReadyBy})
	if err != nil {
		return err
	}

	return nil
}
