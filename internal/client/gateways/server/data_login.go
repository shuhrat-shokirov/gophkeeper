package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	pb "gophkeeper/proto"
)

func (g *gateway) SaveLoginAndPass(ctx context.Context, userID int64, pass *LoginAndPass) error {
	login, err := g.dataServiceClient.SaveLogin(ctx, &pb.LoginData{
		Meta: &pb.BaseData{
			UserId:    proto.Int64(userID),
			Title:     proto.String(pass.Title),
			Note:      proto.String(pass.Note),
			CreatedAt: proto.Int64(pass.ModifiedAt.UnixNano()),
		},
		Login:    proto.String(pass.Login),
		Password: proto.String(pass.Pass),
	})
	if err != nil {
		return fmt.Errorf("save login: %w", err)
	}

	if login.GetStatus() != pb.ResponseStatus_SUCCESS {
		return fmt.Errorf("failed to save login: %s", login.GetMessage())
	}

	return nil
}
