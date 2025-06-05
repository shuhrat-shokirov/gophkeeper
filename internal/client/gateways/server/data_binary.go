package server

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"

	pb "gophkeeper/proto"
)

func (g *gateway) SaveBinary(ctx context.Context, userID int64, data *Binary) error {
	binary, err := g.dataServiceClient.SaveBinary(ctx, &pb.BinaryData{
		Meta: &pb.BaseData{
			UserId:    proto.Int64(userID),
			Title:     proto.String(data.Title),
			Note:      proto.String(data.Note),
			CreatedAt: proto.Int64(data.ModifiedAt.UnixNano()),
		},
		Data: data.Content,
	})
	if err != nil {
		return fmt.Errorf("save binary error: %w", err)
	}

	if binary.GetStatus() != pb.ResponseStatus_SUCCESS {
		return fmt.Errorf("save binary failed: %s", binary.GetMessage())
	}

	return nil

}
