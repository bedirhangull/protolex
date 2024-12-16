package config

import (
	"os"

	"github.com/bedirhangull/protolex/internal/adapter/protolexError"
)

type Config struct {
	ProtoPath string
	ProtoFile os.FileInfo
}

func NewConfig(protoPath string) (*Config, error) {
	fileInfo, err := os.Stat(protoPath)
	if os.IsNotExist(err) {
		notFound := protolexError.NewError("NotFound", "File not found in the given path")
		filePathInfo := protolexError.NewError("Info", protoPath)
		notFound.LogError()
		filePathInfo.LogError()
		return nil, notFound
	}

	return &Config{
		ProtoPath: protoPath,
		ProtoFile: fileInfo,
	}, nil
}

func ReadProtoFile(protoPath string) (string, error) {
	file, err := os.ReadFile(protoPath)
	if err != nil {
		return "", err
	}

	protoContent := string(file)
	return protoContent, nil
}
