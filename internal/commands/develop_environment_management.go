package commands

import (
	"errors"
	"fmt"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func NewDevelopEnvironmentManagementCommand() *cobra.Command {
	return &cobra.Command{
		Use:                   "dem command [options] [args]",
		Short:                 "通过DEM运行命令，[options]和[args]将被传递给command",
		DisableFlagParsing:    true,
		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
		DisableSuggestions:    true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if 0 == len(args) {
				return nil
			}

			command, err := findCommand(args[0])
			if nil != err {
				fmt.Println(err)

				return nil
			}

			pwd, err := os.Getwd()
			if nil != err {
				fmt.Printf("获取当前工作目录失败: %s\n", err)

				return nil
			}

			var systemCommand = exec.Command(command, args[1:]...)
			systemCommand.Dir = pwd

			// 绑定管道
			systemCommand.Stdin = os.Stdin
			systemCommand.Stdout = os.Stdout
			systemCommand.Stderr = os.Stderr

			environments, err := getEnvironments()
			if nil != err {
				fmt.Printf("获取环境变量失败: %s\n", err)

				return nil
			}

			for k, v := range environments {
				systemCommand.Env = append(systemCommand.Env, fmt.Sprintf("%s=%s", k, v))
			}

			if err := systemCommand.Run(); nil != err {
				if errors.Is(err, exec.ErrNotFound) {
					cmd.PrintErrf("全部路径中未找到命令[%s]\n", command)

					return nil
				}

				cmd.PrintErrln(err)
			}

			return nil
		},
	}
}

func findCommand(name string) (string, error) {
	me, err := environment.GetMixedEnvironment()
	if nil != err {
		return "", fmt.Errorf("查询环境配置失败: %s\n", err)
	}

	var command string
	for pkg, version := range me.Packages {
		ind, err := index.Lookup(pkg + "@" + version)
		if nil != err {
			return "", fmt.Errorf("查询工具包[%s@%s]索引失败: %s\n", pkg, version, err)
		}

		for _, fp := range ind.Platforms[system.GetSystemArch()].Paths {
			var file = filepath.Join(system.ReplaceVariables(fp, "{ROOT}", system.GetPackageRootPath(ind.PackageName)), name)
			if system.Executable(file) {
				command = file

				break
			}
		}
	}

	if "" == command {
		command, err = exec.LookPath(name)
		if nil != err && !errors.Is(err, exec.ErrNotFound) {
			return "", fmt.Errorf("查找命令失败: %s\n", err)
		}
	}

	if "" == command {
		command = name
	}

	return command, nil
}

func getEnvironments() (map[string]string, error) {
	var environments = make(map[string]string)

	me, err := environment.GetMixedEnvironment()
	if nil != err {
		return nil, fmt.Errorf("查询环境配置失败: %s\n", err)
	}

	var paths []string
	for pkg, version := range me.Packages {
		ind, err := index.Lookup(pkg + "@" + version)
		if nil != err {
			return nil, fmt.Errorf("查询工具包[%s@%s]索引失败: %s\n", pkg, version, err)
		}

		// 索引中的环境变量
		for k, v := range ind.Platforms[system.GetSystemArch()].Environments {
			environments[k] = system.ReplaceVariables(v, "{ROOT}", system.GetPackageRootPath(ind.PackageName))
		}

		for _, path := range ind.Platforms[system.GetSystemArch()].Paths {
			paths = append(paths, system.ReplaceVariables(path, "{ROOT}", system.GetPackageRootPath(ind.PackageName)))
		}
	}

	for k, v := range me.Environments {
		environments[k] = v
	}

	for _, env := range os.Environ() {
		var kv = strings.SplitN(env, "=", 2)
		environments[kv[0]] = kv[1]
	}

	environments["PATH"] = strings.Join(paths, string(os.PathListSeparator)) + string(os.PathListSeparator) + environments["PATH"]

	return environments, nil
}
