package logic

import (
	"ascenda/common"
	"ascenda/logic/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildDataset(t *testing.T) {

	t.Run("error while fetching hotelIDs", func(t *testing.T) {

		var supplier common.Supplier
		supplier = mocks.MockSupplier{}

		keyVsDataSource = map[string]common.Supplier{
			"dummySupplier": supplier,
		}

		err := buildDataset()
		assert.Error(t, err)
	})

	t.Run("happy path", func(t *testing.T) {
		keyVsDataSource = map[string]common.Supplier{}
		buildDataSource()
		err := buildDataset()
		assert.Nil(t, err)
	})

}
