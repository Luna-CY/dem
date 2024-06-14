package commands

import (
	"fmt"
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/index"
	"github.com/Luna-CY/dem/internal/pkg"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/spf13/cobra"
	"slices"
	"strings"
)

func NewDevelopEnvironmentUtilInstallCommand() *cobra.Command {
	var overwrite bool

	var command = &cobra.Command{
		Use:   "install [options] package [package [...]]",
		Short: "安装工具包",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, name := range args {
				installed, err := pkg.Installed(name)
				if nil != err {
					cmd.PrintErrf("查询工具包[%s]索引失败: %s\n", name, err)

					return nil
				}

				if installed && !overwrite {
					_ = echo.Info("工具包[%s]已安装，跳过", name)

					continue
				}

				ind, err := index.Lookup(name)
				if nil != err {
					cmd.PrintErrf("查询工具包[%s]索引失败: %s\n", name, err)

					return nil
				}

				deps, err := DiscoverDepends(ind)
				if nil != err {
					cmd.PrintErrln(err)

					return nil
				}

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
						cmd.PrintErrf("查询工具包[%s]索引失败: %s\n", name, err)

						return nil
					}

					if installed && !overwrite {
						continue
					}

					names = append(names, dep)
				}

				fmt.Printf("安装工具包[%s]及其依赖:[%s]\n", name, strings.Join(deps, ","))
				names = append(names, name)

				for _, name := range names {
					if err := pkg.Install(cmd.Context(), name); nil != err {
						cmd.PrintErrln(err)

						return nil
					}

					_ = echo.Info("工具包[%s]安装成功", name)
				}
			}

			return nil
		},
	}

	command.Flags().BoolVar(&overwrite, "overwrite", false, "覆盖安装")

	return command
}

func DiscoverDepends(ind *index.Index) ([]string, error) {
	var depends []string

	platform, ok := ind.Platforms[system.GetSystemArch()]
	if !ok {
		return nil, fmt.Errorf("工具包[%s]不支持当前平台: %s", ind.PackageName, system.GetSystemArch())
	}

	// 包比较少，暂时不考虑循环依赖
	for _, pn := range platform.Depends {
		depends = append(depends, pn)

		subInd, err := index.Lookup(pn)
		if nil != err {
			return nil, fmt.Errorf("查询工具包[%s]索引失败: %s", pn, err)
		}

		subs, err := DiscoverDepends(subInd)
		if nil != err {
			return nil, err
		}

		depends = append(depends, subs...)
	}

	slices.Reverse(depends)

	return depends, nil
}
