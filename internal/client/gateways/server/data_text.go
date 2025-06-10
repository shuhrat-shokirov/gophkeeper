package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	pb "gophkeeper/proto"
)

func (g *gateway) SaveText(ctx context.Context, userID int64, data *Text) error {
	response, err := g.dataServiceClient.SaveText(ctx, &pb.TextData{
		Meta: &pb.BaseData{
			UserId:    proto.Int64(userID),
			Title:     proto.String(data.Title),
			Note:      proto.String(data.Note),
			CreatedAt: proto.Int64(data.ModifiedAt.UnixNano()),
		},
		Content: proto.String(data.Content),
	})
	if err != nil {
		return fmt.Errorf("save text: %w", err)
	}

	if response.GetStatus() != pb.ResponseStatus_SUCCESS {
		return fmt.Errorf("save text: status %s", response.GetStatus().String())
	}

	return nil
}
