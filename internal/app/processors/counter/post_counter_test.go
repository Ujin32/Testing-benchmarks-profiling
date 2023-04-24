package counter

import (
	"testing"

	mock_post "module06/test/gomock/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetList(t *testing.T) {
	req := require.New(t)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClient := mock_post.NewMockClient(mockController)

	mockClient.EXPECT().GetList().Return().Times(1)

}
