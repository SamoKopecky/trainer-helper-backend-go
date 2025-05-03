// package service
//
// import (
//
//	"testing"
//	"trainer-helper/model"
//	store "trainer-helper/store/mock"
//	"trainer-helper/testutil"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//
// )
//
//	func TestWeekGetBlocks(t *testing.T) {
//		m := store.NewMockWeek(t)
//		service := Week{Store: m}
//
//		mockModels := []model.Week{
//			*testutil.WeekFactory(testutil.WeekUserId("1")),
//			*testutil.WeekFactory(testutil.WeekUserId("1")),
//			*testutil.WeekFactory(testutil.WeekUserId("1")),
//		}
//		m.EXPECT().GetByUserId("1").Return(mockModels, nil).Once()
//
//		// Act
//		models, err := service.GetBlocks("1")
//
//		// Assert
//		require.Nil(t, err)
//		assert.Equal(t, models[10], []model.Week{mockModels[0], mockModels[1]})
//		assert.Equal(t, models[20], []model.Week{mockModels[2]})
//
// }
package service
