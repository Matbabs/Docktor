package services

import (
	"encoding/json"
	"os/exec"
)

const CONFIG_REPORT_FILE string = ".config.report.json"

func ScanLocalConfig(path string) any {
	var scan any
	exec.Command("trivy", "conf", "-f", "json", "-o", CONFIG_REPORT_FILE, path).Run()
	cmd, _ := exec.Command("cat", CONFIG_REPORT_FILE).Output()
	json.Unmarshal(cmd, &scan)
	return scan
}
