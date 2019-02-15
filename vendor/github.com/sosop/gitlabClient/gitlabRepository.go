package gitlabClient

import (
	"os/exec"
	"github.com/pkg/errors"
)

// 依赖客户端git
func (gitlab *GitLabClient) GetRepo(projectID int) ([]byte, error) {
	for _, proj := range gitlab.Projects {
		if projectID == proj.ProjectID {
			return exec.Command("git clone -b " + proj.Branch + " http://" + gitlab.Username + ":" + gitlab.Password + "@" + proj.WebURL + " /tmp/" + proj.ProjectName).Output()
		}
	}

	return nil, errors.New("没有找到可选项目")
}
