package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/Luna-CY/dem/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

const DemRepoPath = "https://github.com/Luna-CY/dem/raw/master/packages/"

func NewDevelopEnvironmentUtilUpdateCommand() *cobra.Command {
	var repo = ""
	var extensions = ""
	var local = false

	var command = &cobra.Command{
		Use:   "update",
		Short: "更新DEM索引数据",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo = utils.GetStringFromEnv(repo, system.DemEnvPrefix+"REPO_PATH", DemRepoPath)

			var exts = []string{"base"}
			for _, extension := range strings.Split(extensions, ",") {
				if "" == extension {
					continue
				}

				exts = append(exts, "e-"+extension)
			}

			for _, extension := range exts {
				var url = path.Join(repo, extension+".tar.gz")

				var file *os.File
				var size int64
				var err error

				if !local {
					echo.Info("下载[%s]索引库: %s", extension, url)
					file, size, err = utils.DownloadRemoteWithProgress(cmd.Context(), extension+".tar.gz", url)
					if nil != err {
						return echo.Error("下载[%s]索引库失败: %s", extension, err)
					}
				} else {
					file, size, err = utils.DownloadLocalWithProgress(cmd.Context(), url)
					if nil != err {
						return echo.Error("下载[%s]索引库失败: %s", extension, err)
					}
				}

				_ = size
				_ = file.Close()
				_ = os.Remove(file.Name())
			}

			return nil
		},
	}

	command.Flags().StringVarP(&repo, "repo", "r", repo, "DEM索引数据存储路径")
	command.Flags().StringVarP(&extensions, "extensions", "e", extensions, "扩展的工具库名称，用,分割")
	command.Flags().BoolVarP(&local, "local", "l", local, "使用本地索引库")

	return command
}
