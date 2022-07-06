package services

import (
	"encoding/json"
	"os/exec"
)

const REPO_REPORT_FILE string = "./reports/repo.report.json"

func ScanRemoteRepoGit(uri string) any {
	var scan any
	exec.Command("trivy", "repo", "-f", "json", "-o", REPO_REPORT_FILE, uri).Run()
	cmd, _ := exec.Command("cat", REPO_REPORT_FILE).Output()
	json.Unmarshal(cmd, &scan)
	return scan
}
