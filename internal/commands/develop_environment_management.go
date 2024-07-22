package commands

import (
	"errors"
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/environment"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/Luna-CY/dem/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func NewDevelopEnvironmentManagementCommand() *cobra.Command {
	return &cobra.Command{
		Use:                   "dem command [command options] [command args]",
		Short:                 "Run command with dem，[options] and [args] provider for command",
		DisableFlagParsing:    true,
		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
		DisableSuggestions:    true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if 0 == len(args) {
				fmt.Printf("develop Environment Management %s\n\n", system.Version)
				fmt.Printf("usage: dem command [command options] [command args]\n\n")

				echo.Infoln("command find path of dem")

				me, err := environment.GetMixedEnvironment()
				if nil != err {
					echo.Errorln("find environment configuration failed: %s", true, err)

					os.Exit(1)
				}

				for pkg, version := range me.Packages {
					ind, err := index.Lookup(pkg + "@" + version)
					if nil != err {
						echo.Errorln("find index of package[%s@%s] failed: %s", true, pkg, version, err)

						os.Exit(1)
					}

					for _, fp := range ind.Platforms[system.GetSystemArch()].Paths {
						fmt.Println(system.ReplaceVariables(fp, "{ROOT}", system.GetPackageRootPath(ind.PackageName)))
					}
				}

				fmt.Println("")
				echo.Infoln("command find path of system")

				var paths = strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
				for _, fp := range paths {
					fmt.Println(fp)
				}

				fmt.Println("build in command or alias of shell")

				return nil
			}

			command, err := findCommand(args[0])
			if nil != err {
				echo.Errorln("find command failed: %s", true, err)

				os.Exit(1)
			}

			pwd, err := os.Getwd()
			if nil != err {
				echo.Errorln("get current work directory failed: %s", true, err)

				os.Exit(1)
			}

			var systemCommand = exec.Command(command, args[1:]...)
			systemCommand.Dir = pwd

			// 绑定管道
			systemCommand.Stdin = os.Stdin
			systemCommand.Stdout = os.Stdout
			systemCommand.Stderr = os.Stderr

			environments, err := getEnvironments()
			if nil != err {
				echo.Errorln("get environment variables failed: %s", true, err)

				os.Exit(1)
			}

			for k, v := range environments {
				systemCommand.Env = append(systemCommand.Env, fmt.Sprintf("%s=%s", k, v))
			}

			if err := systemCommand.Run(); nil != err {
				if errors.Is(err, exec.ErrNotFound) {
					echo.Errorln("command not found in all paths: %s", false, command)

					os.Exit(1)
				}

				echo.Errorln("run command failed: %s", true, err)
			}

			return nil
		},
	}
}

func findCommand(name string) (string, error) {
	me, err := environment.GetMixedEnvironment()
	if nil != err {
		return "", fmt.Errorf("find environment configuration failed: %s", err)
	}

	var command string
	for pkg, version := range me.Packages {
		ind, err := index.Lookup(pkg + "@" + version)
		if nil != err {
			return "", fmt.Errorf("find index of package[%s@%s] failed: %s", pkg, version, err)
		}

		for _, fp := range ind.Platforms[system.GetSystemArch()].Paths {
			var file = filepath.Join(system.ReplaceVariables(fp, "{ROOT}", system.GetPackageRootPath(ind.PackageName)), name)
			if utils.Executable(file) {
				command = file

				break
			}
		}
	}

	if "" == command {
		command, err = exec.LookPath(name)
		if nil != err && !errors.Is(err, exec.ErrNotFound) {
			return "", fmt.Errorf("find command failed: %s", err)
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
		return nil, fmt.Errorf("find environment configuration failed: %s", err)
	}

	var paths []string
	for pkg, version := range me.Packages {
		ind, err := index.Lookup(pkg + "@" + version)
		if nil != err {
			return nil, fmt.Errorf("find index of package[%s@%s] failed: %s", pkg, version, err)
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
