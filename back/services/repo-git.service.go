package services

import (
	"encoding/json"
)

const REPO_REPORT_FILE string = ".repo.report.json"

func ScanRemoteRepoGit(uri string) any {
	var scan any
	executor.Command("trivy", "repo", "-f", "json", "-o", REPO_REPORT_FILE, uri).Run()
	cmd, _ := executor.Command("cat", REPO_REPORT_FILE).Output()
	json.Unmarshal(cmd, &scan)
	return scan
}
