package services

import (
	"main/models"
	"reflect"
	"testing"
)

type testGetLocalDockerImages struct {
	name string
	res  []models.DockerImage
}

func TestGetLocalDockerImages(t *testing.T) {

	executor = execCommandMock{}

	tests := []testGetLocalDockerImages{
		{
			"images",
			[]models.DockerImage{
				{Id: "b97ca34de947", Name: "test", Tag: "latest", Date: "6 hours ago", Size: "1.44GB"},
				{Id: "ca26c522ce21", Name: "matbabs/docktor", Tag: "latest", Date: "6 months ago", Size: "1.39GB"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := GetLocalDockerImages(); !reflect.DeepEqual(res, test.res) {
				t.Errorf("Expected %v, received %v", test.res, res)
			}

		})
	}
}
