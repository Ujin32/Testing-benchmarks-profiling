package counter

import (
	"testing"

	mock_post "module06/test/gomock/mocks"

	"module06/internal/app/services/post"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetList(t *testing.T) {
	req := require.New(t)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClient := mock_post.NewMockClient(mockController)

	t.Run("get post return empty list", func(t *testing.T) {
		empty_list := []post.Post{}
		mockClient.EXPECT().GetList().Return(empty_list, nil).Times(1)
		posts, err := mockClient.GetList()
		req.NoError(err)
		req.Empty(posts)

	})
	t.Run("count post with empty list", func(t *testing.T) {
		expectedPosts := []post.Post{
			{ID: 1, Title: "Пост 1"},
			{ID: 2, Title: "Пост 2"},
			{ID: 3, Title: "Пост 3"},
		}
		mockClient.EXPECT().GetList().Return(expectedPosts, nil).AnyTimes()
		posts, err := mockClient.GetList()
		req.NoError(err)
		req.NotEmpty(posts)

		lenPosts, err := PostCount(mockClient)
		req.NoError(err)
		expPosts := len(expectedPosts)
		req.Equal(expPosts, lenPosts, "Количество постов должно быть равным %d, а результат функции %d", expPosts, lenPosts)

	})

}
