package generator

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ConvertFile(srcPath string, op Options) error {
	ext := filepath.Ext(srcPath)
	if ext != ".go" {
		return nil
	}

	srcCode, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	srcFilename := filepath.Base(srcPath)
	base := strings.TrimSuffix(srcFilename, ext)
	destDir := filepath.Dir(srcPath)
	destFilename := fmt.Sprintf("%s%s%s", base, op.Suffix, ext)
	destCode, err := Convert(srcCode, srcFilename, destFilename)
	if err != nil {
		return err
	}
	if destCode == nil {
		return nil
	}

	destPath := filepath.Join(destDir, destFilename)
	return ioutil.WriteFile(destPath, destCode, 0664)
}
