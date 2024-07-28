package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	// "github.com/georgechieng-sc/interns-2022/folders"
	// "github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	// Test case 1: method call with existing folder ID.
	t.Run("retrieval with existing folderID", func(t *testing.T) {
		orgID := uuid.FromStringOrNil(folders.DefaultOrgID)
		req := &folders.FetchFolderRequest{OrgID: orgID}

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.Folders)
		for _, folder := range res.Folders {
			assert.Equal(t, orgID, folder.OrgId)
		}
	})

	// Test case 2: method call with non-existing folder ID. A new uuid is created and used when calling the method.
	t.Run("retrieval with non-existent folderID", func(t *testing.T) {
		orgID := uuid.Must(uuid.NewV4())
		req := &folders.FetchFolderRequest{OrgID: orgID}

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Empty(t, res.Folders)
	})
}
