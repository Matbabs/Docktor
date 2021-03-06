package services

import (
	"encoding/json"
	"main/models"
	"os/exec"
	"strings"
)

const DOCKER_REPORT_FILE string = "./reports/docker.report.json"

func GetLocalDockerImages() []models.DockerImage {
	var dockerImages []models.DockerImage
	cmd, _ := exec.Command("docker", "images", "--format", "'{{.ID}};{{.Repository}};{{.Tag}};{{.CreatedSince}};{{.Size}}'").Output()
	for _, line := range strings.Split(string(cmd), "\n") {
		s := strings.Split(line, ";")
		if len(s) == 5 && s[2] != "<none>" {
			dockerImage := models.DockerImage{Id: strings.Trim(s[0], "'"), Name: s[1], Tag: s[2], Date: s[3], Size: strings.Trim(s[4], "'")}
			dockerImages = append(dockerImages, dockerImage)
		}
	}
	return dockerImages
}

func ScanLocalDockerImage(image string) any {
	var scan any
	exec.Command("trivy", "i", "-f", "json", "-o", DOCKER_REPORT_FILE, image).Run()
	cmd, _ := exec.Command("cat", DOCKER_REPORT_FILE).Output()
	json.Unmarshal(cmd, &scan)
	return scan
}

func ScanRemoteDockerImage(image string) any {
	exec.Command("docker", "pull", image).Run()
	return ScanLocalDockerImage(image)
}
