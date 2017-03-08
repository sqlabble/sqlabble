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

	base := strings.TrimSuffix(filepath.Base(srcPath), ext)
	distDir := filepath.Dir(srcPath)
	distFileName := fmt.Sprintf("%s%s%s", base, op.Suffix, ext)
	distCode, err := Convert(srcCode, distFileName)
	if err != nil {
		return err
	}
	if distCode == nil {
		return nil
	}

	distPath := filepath.Join(distDir, distFileName)
	return ioutil.WriteFile(distPath, distCode, 0664)
}
