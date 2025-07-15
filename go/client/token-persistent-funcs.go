package client

import (
	"fmt"
	"os"
)

type (
	filesystem struct {
		tokenPath string
	}
)

func NewFilesystemTokenPersiter(tokenPath string) (PersistTokenFn, error) {
	fileInfo, err := os.Stat(tokenPath)
	if err != nil {
		return nil, fmt.Errorf("unable to stat tokenpath:%w", err)
	}
	mode := fileInfo.Mode()
	if mode&os.ModePerm == os.ModePerm {
		return nil, fmt.Errorf("tokenpath %s is not writable", tokenPath)
	}
	f := &filesystem{
		tokenPath: tokenPath,
	}
	return f.persist, nil
}

func (f *filesystem) persist(token string) error {
	return os.WriteFile(f.tokenPath, []byte(token), 0600)
}
