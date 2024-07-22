package commands

import (
	"github.com/Luna-CY/dem/internal/echo"
	"github.com/Luna-CY/dem/internal/system"
	"github.com/Luna-CY/dem/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

const DemRepoPath = "https://github.com/Luna-CY/dem/raw/main/packages"

func NewDevelopEnvironmentUtilUpdateCommand() *cobra.Command {
	var repo = ""
	var extensions = ""
	var local = false

	var command = &cobra.Command{
		Use:   "update",
		Short: "update dem index",
		Run: func(cmd *cobra.Command, args []string) {
			repo = utils.GetStringFromEnv(repo, system.DemEnvPrefix+"REPO_PATH", DemRepoPath)

			var exts = []string{"base"}
			for _, extension := range strings.Split(extensions, ",") {
				if "" == extension {
					continue
				}

				exts = append(exts, extension)
			}

			for _, extension := range exts {
				var filename = extension + ".tar.gz"
				var target = filepath.Join(system.GetIndexPath(), filename)
				var url = strings.TrimSuffix(repo, "/") + "/" + filename

				if !local {
					echo.Infoln("download [%s] index: %s", extension, url)

					if err := utils.DownloadRemoteWithProgress(cmd.Context(), filename, target, url, ""); nil != err {
						echo.Errorln("download [%s] index failed: %s", true, extension, err)

						os.Exit(1)
					}
				} else {
					echo.Infoln("download [%s] local index: %s", extension, url)

					if err := utils.DownloadLocalWithProgress(cmd.Context(), filename, target, url); nil != err {
						echo.Errorln("download [%s] local index failed: %s", true, extension, err)

						os.Exit(1)
					}
				}

				if err := os.RemoveAll(filepath.Join(system.GetIndexPath(), extension)); nil != err {
					echo.Errorln("clean old [%s] index failed: %s", true, extension, err)

					os.Exit(1)
				}

				if err := utils.GzipDecompressWithProgress(cmd.Context(), system.GetIndexPath(), filename, target); nil != err {
					echo.Errorln("decompress [%s] index failed: %s", true, extension, err)

					os.Exit(1)
				}

				echo.Infoln("index [%s] update completed", extension)
			}
		},
	}

	command.Flags().StringVarP(&repo, "repo", "r", repo, "set index repository path")
	command.Flags().StringVarP(&extensions, "extensions", "e", extensions, "used extension indexes (comma separated)")
	command.Flags().BoolVarP(&local, "local", "l", local, "use local index repository")

	return command
}
