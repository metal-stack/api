package client

import "os"

type filesystem struct {
	tokenPath string
}

func NewFilesystemTokenPersiter(tokenPath string) *filesystem {
	return &filesystem{
		tokenPath: tokenPath,
	}
}

func (f *filesystem) Persist(token string) error {
	return os.WriteFile(f.tokenPath, []byte(token), 0600)
}
