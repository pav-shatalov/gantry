package docker

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"strings"
	"time"

	dockerSdkContainer "github.com/docker/docker/api/types/container"
	dockerSdkClient "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Container struct {
	Id    string
	Name  string
	Image string
}

func (c Container) String() string {
	return c.Name
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()
	dockerContainers, err := c.dockerClient.ContainerList(ctx, dockerSdkContainer.ListOptions{})
	if err != nil {
		return containers, err
	}

	for _, ctr := range dockerContainers {
		containers = append(containers, Container{
			Id:    ctr.ID,
			Name:  strings.Trim(strings.Join(ctr.Names, "|"), "/"),
			Image: ctr.Image,
		})
	}

	return containers, nil
}

func (c Client) Version() string {
	return c.dockerClient.ClientVersion()
}

func (c Client) ServerVersion() (string, error) {
	if c.dockerClient == nil {
		return "", errors.New("Docker client is missing")
	}

	v, err := c.dockerClient.ServerVersion(context.TODO())
	if err != nil {
		return "", err
	}

	return v.Version, nil
}

func (c Client) ContainerLogs(ctrId string) ([]string, error) {
	var logs []string
	if c.dockerClient == nil {
		return logs, errors.New("Docker client is missing")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	reader, err := c.dockerClient.ContainerLogs(ctx, ctrId, dockerSdkContainer.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: false,
		Follow:     false,
		Tail:       "50",
	})
	if err != nil {
		return logs, err
	}
	defer reader.Close()

	buf := new(bytes.Buffer)
	_, stdCopyErr := stdcopy.StdCopy(buf, buf, reader)
	if stdCopyErr != nil {
		return logs, err
	}

	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return logs, err
	}

	return logs, nil
}
