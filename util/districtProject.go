package util

import "github.com/sosop/gitlabClient"

func DistictProject(projs []gitlabClient.ProjectInfo) []gitlabClient.ProjectInfo {
	cap := len(projs)
	mainIDs := make(map[int]struct{}, cap)
	distProjs := make([]gitlabClient.ProjectInfo, 0, cap)
	for _, p := range projs {
		// 没有重复
		if _, ok := mainIDs[p.ProjectID]; !ok {
			distProjs = append(distProjs, p)
			mainIDs[p.ProjectID] = struct{}{}
		}
	}

	return distProjs
}
