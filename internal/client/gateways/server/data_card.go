package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	pb "gophkeeper/proto"
)

func (g *gateway) SaveCard(ctx context.Context, userID int64, data *Card) error {
	card, err := g.dataServiceClient.SaveCard(ctx, &pb.CardData{
		Meta: &pb.BaseData{
			UserId:    proto.Int64(userID),
			Title:     proto.String(data.Title),
			Note:      proto.String(data.Note),
			CreatedAt: proto.Int64(data.ModifiedAt.UnixNano()),
		},
		Pan:    proto.String(data.Pan),
		Expiry: proto.String(data.Expiry),
		Cvv:    proto.String(data.Cvv),
	})
	if err != nil {
		return fmt.Errorf("save card error: %w", err)
	}

	if card.GetStatus() != pb.ResponseStatus_SUCCESS {
		return fmt.Errorf("save card failed: %s", card.GetMessage())
	}

	return nil
}
