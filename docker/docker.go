package docker

import (
	"context"
	"errors"
	"strings"

	dockerSdkContainer "github.com/docker/docker/api/types/container"
	dockerSdkClient "github.com/docker/docker/client"
)

type Container struct {
	Id    string
	Name  string
	Image string
}

type Client struct {
	dockerClient *dockerSdkClient.Client
}

func NewClient() (Client, error) {
	cli, err := dockerSdkClient.NewClientWithOpts(dockerSdkClient.FromEnv, dockerSdkClient.WithAPIVersionNegotiation())

	if err != nil {
		return Client{}, err
	}

	client := Client{
		dockerClient: cli,
	}

	return client, nil

}

func (c Client) LoadContainerList() ([]Container, error) {
	var containers []Container

	if c.dockerClient == nil {
		return containers, errors.New("Docker client is missing")
	}
	dockerContainers, err := c.dockerClient.ContainerList(context.TODO(), dockerSdkContainer.ListOptions{})
	if err != nil {
		return containers, err
	}

	for _, ctr := range dockerContainers {
		containers = append(containers, Container{
			Id:    ctr.ID,
			Name:  strings.Join(ctr.Names, "|"),
			Image: ctr.Image,
		})
	}

	return containers, nil
}
