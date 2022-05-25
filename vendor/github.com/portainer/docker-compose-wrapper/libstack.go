package libstack

import (
	"context"
)

type Deployer interface {
	Deploy(ctx context.Context, workingDir, host, projectName string, filePaths []string, envFilePath string, forceRereate bool) error
	Remove(ctx context.Context, workingDir, host, projectName string, filePaths []string) error
	Pull(ctx context.Context, workingDir, host, projectName string, filePaths []string) error
}