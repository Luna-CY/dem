package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/Luna-CY/dem/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
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
				var filename = extension + ".tar.gz"
				var url = path.Join(repo, filename)

				var file *os.File
				var size int64
				var err error

				if !local {
					_ = echo.Info("下载[%s]索引库: %s", extension, url)
					file, size, err = utils.DownloadRemoteWithTmpFileAndProgress(cmd.Context(), filename, url)
					if nil != err {
						return echo.Error("下载[%s]索引库失败: %s", extension, err)
					}
				} else {
					_ = echo.Info("下载[%s]索引库: %s", extension, url)
					file, size, err = utils.DownloadLocalWithProgress(cmd.Context(), filename, url)
					if nil != err {
						return echo.Error("下载[%s]索引库失败: %s", extension, err)
					}
				}

				if err := os.RemoveAll(filepath.Join(system.GetIndexPath(), extension)); nil != err {
					return echo.Error("清理就的[%s]索引库失败: %s", extension, err)
				}

				if err := utils.GzipDecompressWithProgress(cmd.Context(), system.GetIndexPath(), filename, file, size); nil != err {
					_ = file.Close()
					_ = os.Remove(file.Name())

					return echo.Error("解压[%s]索引库失败: %s", extension, err)
				}

				_ = file.Close()
				_ = os.Remove(file.Name())

				_ = echo.Info("索引库[%s]更新完成", extension)
			}

			return nil
		},
	}

	command.Flags().StringVarP(&repo, "repo", "r", repo, "DEM索引数据存储路径")
	command.Flags().StringVarP(&extensions, "extensions", "e", extensions, "扩展的工具库名称，用,分割")
	command.Flags().BoolVarP(&local, "local", "l", local, "使用本地索引库")

	return command
}
