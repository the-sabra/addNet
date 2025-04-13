package addition

import (
	"addition_service/internal/api/addition"
	"addition_service/internal/outbox"
	"context"
)

type AdditionService struct {
	addition.UnimplementedAdditionServiceServer
	outboxRepo *outbox.SQLRepository
}

func NewAdditionService(outboxRepo *outbox.SQLRepository) *AdditionService {
	return &AdditionService{
		outboxRepo: outboxRepo,
	}
}

func (s *AdditionService) Add(ctx context.Context, req *addition.AddRequest) (*addition.AddResponse, error) {
	result := req.GetA() + req.GetB()
	s.outboxRepo.Store(ctx, result)
	return &addition.AddResponse{Result: result}, nil
}
