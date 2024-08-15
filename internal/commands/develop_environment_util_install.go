package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
	"os"
	"slices"
	"strings"
)

func NewDevelopEnvironmentUtilInstallCommand() *cobra.Command {
	var overwrite bool

	var command = &cobra.Command{
		Use:   "install [options] package [package [...]]",
		Short: "install packages",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, name := range args {
				installed, err := pkg.Installed(name)
				if nil != err {
					echo.Errorln("check package[%s] installed status failed: %s", true, name, err)

					os.Exit(1)
				}

				if installed && !overwrite {
					echo.Infoln("package[%s] has been installed, skip", name)

					continue
				}

				ind, err := index.Lookup(name)
				if nil != err {
					echo.Errorln("lookup package[%s] index failed: %s", true, name, err)

					os.Exit(1)
				}

				deps, err := DiscoverDepends(ind)
				if nil != err {
					echo.Errorln(err.Error(), true)

					os.Exit(1)
				}

				slices.Reverse(deps)

				// 去重
				var mapping = make(map[string]struct{})
				// 必须保留依赖顺序
				var names []string

				for _, dep := range deps {
					if _, ok := mapping[dep]; ok {
						continue
					}

					installed, err := pkg.Installed(dep)
					if nil != err {
						echo.Errorln("find package[%s] index failed: %s", true, name, err)

						os.Exit(1)
					}

					if installed {
						continue
					}

					names = append(names, dep)
				}

				if 0 == len(names) {
					echo.Infoln("install package[%s]", name)
				} else {
					echo.Infoln("install package[%s] and its depends[%s]", name, strings.Join(names, ","))
				}

				names = append(names, name)

				for _, name := range names {
					if err := pkg.Install(cmd.Context(), name); nil != err {
						echo.Errorln(err.Error(), true)

						os.Exit(1)
					}

					echo.Infoln("package[%s] installed successfully", name)
				}
			}
		},
	}

	command.Flags().BoolVar(&overwrite, "overwrite", false, "overwrite install")

	return command
}

func DiscoverDepends(ind *index.Index) ([]string, error) {
	var depends []string

	platform, ok := ind.Platforms[system.GetSystemArch()]
	if !ok {
		return nil, fmt.Errorf("package[%s] does not support current platform: %s", ind.PackageName, system.GetSystemArch())
	}

	// 包比较少，暂时不考虑循环依赖
	for _, pn := range platform.Depends {
		depends = append(depends, pn)

		subInd, err := index.Lookup(pn)
		if nil != err {
			return nil, fmt.Errorf("find package[%s] index failed: %s", pn, err)
		}

		subs, err := DiscoverDepends(subInd)
		if nil != err {
			return nil, err
		}

		depends = append(depends, subs...)
	}

	return depends, nil
}
