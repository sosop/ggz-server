package gitlabClient

func (gitlab *GitLabClient) ListBranch(projectID string) ([]byte, error) {
	data, err := gitlab.get("/api/v3/projects/" + projectID + "/repository/branches", nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (gitlab *GitLabClient) GetBranch(projectID, branchName string) ([]byte, error) {
	data, err := gitlab.get("/api/v3//projects/" + projectID +  "/repository/branches/" + branchName, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}