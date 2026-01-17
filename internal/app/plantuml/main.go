package plantuml

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/2lovecode/graffito/pkg/fs"
	"github.com/2lovecode/graffito/pkg/logging"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

type runtime struct {
	projectRootDir string
	appDir         string
	dataDir        string
	imageDir       string
	plantumlLib    string
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plantuml",
		Short: "PlantUML时序图和UML图生成工具",
		Long: `监听指定目录，自动将PlantUML文件转换为图片。

功能:
  - 监听 internal/app/plantuml/data/ 目录
  - 自动将 .txt 文件转换为 .png 图片
  - 生成的图片保存在 internal/app/plantuml/images/ 目录

要求:
  - 需要安装 Java 8+
  - PlantUML jar 文件需要存在于 third_party/plantuml/ 目录

示例:
  graffito cli tools plantuml
  # 在 data/ 目录创建或修改 .txt 文件，会自动生成图片
`,
		Run: func(cmd *cobra.Command, args []string) {
			r := buildRuntime()
			if checkErr := checkEnv(r); checkErr != nil {
				logging.Fatalf("环境检查失败: %v", checkErr)
				return
			}

			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				logging.Fatalf("创建文件监听器失败: %v", err)
				return
			}
			defer func() {
				if closeErr := watcher.Close(); closeErr != nil {
					logging.Errorf("关闭文件监听器失败: %v", closeErr)
				}
			}()

			logging.Infof("开始监听目录: %s", r.dataDir)
			logging.Infof("图片输出目录: %s", r.imageDir)

			done := make(chan bool)
			go func() {
				for {
					select {
					case event, ok := <-watcher.Events:
						if !ok {
							return
						}
						if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
							logging.Debugf("检测到文件变更: %s", event.Name)
							generateCmd := exec.Command("java", "-jar", r.plantumlLib, event.Name, "-o", r.imageDir)
							if runErr := generateCmd.Run(); runErr != nil {
								logging.Errorf("生成图片失败: %v", runErr)
							} else {
								logging.Infof("成功生成图片: %s", event.Name)
							}
						}
					case watchErr, ok := <-watcher.Errors:
						if !ok {
							return
						}
						logging.Errorf("文件监听错误: %v", watchErr)
					}
				}
			}()

			if err = watcher.Add(r.dataDir); err != nil {
				logging.Fatalf("添加监听目录失败: %v", err)
				return
			}

			logging.Info("PlantUML监听服务已启动，按 Ctrl+C 退出")
			<-done
		},
	}
	return cmd
}

func buildRuntime() *runtime {
	projectRoot, _ := fs.GetProjectRootPath()

	plantumlPath := fmt.Sprintf("%s/third_party/plantuml/plantuml-1.2022.7.jar", projectRoot)

	appPath := fs.GetCurrentAbsPath()

	dataDirName := "data"
	imagesDirName := "images"

	dataDirPath := fmt.Sprintf("%s/%s", appPath, dataDirName)
	imagesDirPath := fmt.Sprintf("%s/%s", appPath, imagesDirName)

	return &runtime{
		projectRootDir: projectRoot,
		appDir:         appPath,
		dataDir:        dataDirPath,
		imageDir:       imagesDirPath,
		plantumlLib:    plantumlPath,
	}
}

func checkEnv(r *runtime) error {
	if _, plantumlLibErr := os.Stat(r.plantumlLib); plantumlLibErr != nil {
		return fmt.Errorf("PlantUML jar文件不存在: %v", plantumlLibErr)
	}

	// 检查 Java 环境
	javaCmd := exec.Command("java", "-version")
	output, err := javaCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Java环境未安装或不可用: %v", err)
	}

	logging.Debugf("Java环境检查通过:\n%s", string(output))
	return nil
}
