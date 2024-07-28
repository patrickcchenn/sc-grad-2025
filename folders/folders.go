package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders retrieves all folders for a given organization ID
// improvements that are implemented:
//  1. use more descriptive variable names.
//  2. remove unused variables.
//  3. apply error handling for `FetchAllFoldersByOrgID` function call.
//  4. Improve readability by removing redundant code and make code more straightforward. e.g:
//     4.1. remove `f` and `fp` variables to avoid redundancy.
//     4.2. merge `ffr` variable declaration with assignment.
// 	   4.3. remove unecessary loops and conversions.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	ffr := &FetchFolderResponse{Folders: folders}
	return ffr, nil
}

// FetchAllFoldersByOrgID retrieves folders filtered by organization ID
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
