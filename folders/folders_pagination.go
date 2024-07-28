package folders

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// New struct for paginated response
type FetchFolderPaginatedResponse struct {
	Folders []*Folder
	Token   string
}

// This function takes a request, the desired length of page, and a token
// returns a list of folders with the desired length and token for next page.
// Short explanation:
// 1. we first fetch all of the folders with `FetchAllFoldersByOrgID`
// 2. the token is base64 encoded. If a token is provided, we decode to get the starting index for current page
// 3. slice the array to get current page of folders from the starting index with the desired length
// 4. return the current page and the next token(if applicable) by encoding the end index.
// Implemented this function in `main.go`. run `main.go` to see the output.
func GetAllFoldersPaginated(req *FetchFolderRequest, pageLen int, token string) (*FetchFolderPaginatedResponse, error) {
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	startIndex := 0
	if token != "" {
		decodedToken, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			return nil, fmt.Errorf("invalid pagination token")
		}
		startIndex, err = strconv.Atoi(string(decodedToken))
		if err != nil {
			return nil, fmt.Errorf("failed to convert ASCII to int")
		}
	}
	endIndex := startIndex + pageLen

	if endIndex > len(folders) {
		endIndex = len(folders)
	}

	newToken := ""
	if endIndex < len(folders) {
		newToken = base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(endIndex)))
	}

	folders = folders[startIndex:endIndex]

	ffr := &FetchFolderPaginatedResponse{Folders: folders, Token: newToken}

	return ffr, nil
}
