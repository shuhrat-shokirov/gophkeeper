//nolint:dupl,gocritic
package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func Test_gateway_SaveLoginAndPass(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveLogin", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveLoginAndPass(t.Context(), 1, &LoginAndPass{})
		assert.Error(t, err)
	})

	t.Run("unsuccessful save", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveLogin", mock.Anything, mock.Anything).
			Return(&pb.Response{Status: pb.ResponseStatus_ERROR.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveLoginAndPass(t.Context(), 1, &LoginAndPass{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveLogin", mock.Anything, mock.Anything).
			Return(&pb.Response{Status: pb.ResponseStatus_SUCCESS.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveLoginAndPass(t.Context(), 1, &LoginAndPass{})
		assert.NoError(t, err)
	})
}

func Test_gateway_GetLoginList(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetLoginList", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		logins, err := g.GetLoginList(t.Context(), 1, 10, 0)
		assert.Error(t, err)
		assert.Nil(t, logins)
	})

	t.Run("not found", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetLoginList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{Status: pb.ResponseListStatus_LIST_NOT_FOUND.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		logins, err := g.GetLoginList(t.Context(), 1, 10, 0)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
		assert.Nil(t, logins)
	})

	t.Run("unsuccessful get", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetLoginList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{Status: pb.ResponseListStatus_LIST_ERROR.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		logins, err := g.GetLoginList(t.Context(), 1, 10, 0)
		assert.Error(t, err)
		assert.Nil(t, logins)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetLoginList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{
				Status: pb.ResponseListStatus_LIST_SUCCESS.Enum(),
				Items: []*pb.ListResp{
					{Id: proto.Int64(1), Title: proto.String("Login1")},
					{Id: proto.Int64(2), Title: proto.String("Login2")},
				},
			}, nil)

		g := &gateway{dataServiceClient: mockClient}

		logins, err := g.GetLoginList(t.Context(), 1, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, logins, 2)
		assert.Equal(t, "Login1", logins[0].Title)
		assert.Equal(t, "Login2", logins[1].Title)
	})
}

func Test_gateway_GetLoginByID(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetLoginByID", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		login, err := g.GetLoginByID(t.Context(), 1, 1)
		assert.Error(t, err)
		assert.Nil(t, login)
	})

	t.Run("not found", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetLoginByID", mock.Anything, mock.Anything).
			Return(&pb.LoginDataResponse{
				Status: pb.ResponseListStatus_LIST_NOT_FOUND.Enum(),
			}, nil)

		g := &gateway{dataServiceClient: mockClient}

		login, err := g.GetLoginByID(t.Context(), 1, 1)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
		assert.Nil(t, login)
	})

	t.Run("unsuccessful get", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetLoginByID", mock.Anything, mock.Anything).
			Return(&pb.LoginDataResponse{
				Status: pb.ResponseListStatus_LIST_ERROR.Enum(),
			}, nil)

		g := &gateway{dataServiceClient: mockClient}

		login, err := g.GetLoginByID(t.Context(), 1, 1)
		assert.Error(t, err)
		assert.Nil(t, login)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		createdAt := time.Now().UnixNano()

		mockClient.On("GetLoginByID", mock.Anything, mock.Anything).
			Return(&pb.LoginDataResponse{
				Id:     proto.Int64(1),
				Status: pb.ResponseListStatus_LIST_SUCCESS.Enum(),
				Data: &pb.LoginData{
					Meta: &pb.BaseData{
						Title:     proto.String("Login title"),
						Note:      proto.String("Some note"),
						CreatedAt: proto.Int64(createdAt),
					},
					Login:    proto.String("user"),
					Password: proto.String("pass"),
				},
			}, nil)

		g := &gateway{dataServiceClient: mockClient}

		login, err := g.GetLoginByID(t.Context(), 1, 1)
		assert.NoError(t, err)
		assert.NotNil(t, login)
		assert.Equal(t, int64(1), login.ID)
		assert.Equal(t, "user", login.Login)
		assert.Equal(t, "pass", login.Pass)
		assert.Equal(t, "Login title", login.Title)
		assert.Equal(t, "Some note", login.Note)
	})
}
