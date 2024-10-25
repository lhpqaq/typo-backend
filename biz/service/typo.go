package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/go-git/go-git/v5"
)

// const proxyURL = "socks5://127.0.0.1:7890"

func cloneRepo(gitUrl string) (string, error) {
	repoURL := gitUrl
	dir := path.Join(os.TempDir(), "repo")
	err := os.RemoveAll(dir)
	if err != nil {
		return "", err
	}
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:             repoURL,
		InsecureSkipTLS: true,
	})
	if err != nil {
		return "", err
	}
	return dir, nil
}

// CheckTypo 检查 Git 仓库中的拼写错误
func CheckTypo(gitUrl string) (string, error) {
	dir, err := cloneRepo(gitUrl)
	if err != nil {
		return "", err
	}

	// 切换到克隆的目录
	err = os.Chdir(dir)
	if err != nil {
		return "", err
	}

	// 执行 typos 命令
	cmd := exec.Command("typos") // 这里假设 typos 是可执行命令
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // 捕获错误输出

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command execution error: %v, output: %s", err, out.String())
	}

	// 获取输出，处理颜色转义字符
	output := out.String()

	// 这里可以添加颜色转义字符的处理逻辑
	return output, nil
}
