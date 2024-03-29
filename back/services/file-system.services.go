package services

import (
	"encoding/json"
	"strings"
)

const PATH_REPORT_FILE string = ".path.report.json"

func CheckLocalPath(path string) []string {
	var stdout []string
	cmd, _ := executor.Command("ls", "-a", path).Output()
	stdout = append(stdout, strings.Split(string(cmd), "\n")...)
	return stdout
}

func ScanLocalPath(path string) any {
	var scan any
	executor.Command("trivy", "fs", "-f", "json", "-o", PATH_REPORT_FILE, path).Run()
	cmd, _ := executor.Command("cat", PATH_REPORT_FILE).Output()
	json.Unmarshal(cmd, &scan)
	return scan
}
