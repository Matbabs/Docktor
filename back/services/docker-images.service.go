package services

import (
	"encoding/json"
	"fmt"
	"main/models"
	"strings"
)

const DOCKER_REPORT_FILE string = ".docker.report.json"

func GetLocalDockerImages() []models.DockerImage {
	var dockerImages []models.DockerImage
	cmd, _ := executor.Command("docker", "images", "--format", "'{{.ID}};{{.Repository}};{{.Tag}};{{.CreatedSince}};{{.Size}}'").Output()
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
	fmt.Println(DOCKER_REPORT_FILE)
	fmt.Println(image)
	fmt.Println(executor.Command("trivy", "i", "-f", "json", "-o", DOCKER_REPORT_FILE, image))
	out, err := executor.Command("trivy", "i", "-f", "json", "-o", DOCKER_REPORT_FILE, image).Output()
	fmt.Println(out)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err.Error())
	}
	cmd, _ := executor.Command("cat", DOCKER_REPORT_FILE).Output()
	json.Unmarshal(cmd, &scan)
	return scan
}

func ScanRemoteDockerImage(image string) any {
	executor.Command("docker", "pull", image).Run()
	return ScanLocalDockerImage(image)
}
