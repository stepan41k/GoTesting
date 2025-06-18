package age

import (
	"testing"

	"github.com/stepan41k/GoTesting/task-6/age/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAge(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mocks.NewMockAgeGetter(ctrl)
		mockService.EXPECT().GetAge("Arnold").Return(22, nil)

		result, err := Age(mockService, "Arnold")
		assert.Equal(t, 22, result)
		assert.NoError(t, err)
	})
}