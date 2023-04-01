package gin_dj

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func moduleName() string {
	file, err := os.Open("go.mod")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	module := strings.TrimPrefix(firstLine, "module ")
	return module
}

func read_and_parse(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("failed to read file: %w", err)
		return []byte(err.Error())
	}
	return content
}

func create_file(filename string, content string) {
	newFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	// Write the modified content to the new Go file
	if _, err := newFile.WriteString(content); err != nil {
		panic(err)
	}
}

func execGoModInit(projectName string) error {
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func execGoModGet() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func TemplateInit(cmd *cobra.Command, args []string) {
	projectName := args[0]
	err := os.Mkdir(projectName, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	parentDir := filepath.Dir(dir)
	// Read the file contents into a string
	err = os.Mkdir(projectName+"/core", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	settings_name := "settings.go"
	settings_content := read_and_parse(parentDir + "/core/settings.go")
	// Create the file
	create_file(projectName+"/core/"+settings_name, string(settings_content))

	// Read the file contents into a string
	main_name := "main.go"
	main_content := read_and_parse(parentDir + "/gin_dj/run.go")
	modifiedContent := strings.ReplaceAll(string(main_content), "github.com/Zncl2222/gin-dj", projectName)
	modifiedContent = strings.ReplaceAll(string(modifiedContent), "gin_dj", "main")
	// Create the file
	create_file(projectName+"/"+main_name, string(modifiedContent))

	// Read the file contents into a string
	urls_name := "urls.go"
	urls_content := read_and_parse(parentDir + "/gin_dj/urls.go")
	modifiedContent = strings.ReplaceAll(string(urls_content), "gin_dj", "main")
	// Create the file
	create_file(projectName+"/"+urls_name, string(modifiedContent))

	err = os.Chdir(projectName)
	if err != nil {
		panic(err)
	}
	if err := execGoModInit(projectName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("go mod init successful")

	if err := execGoModGet(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func CreateApp(cmd *cobra.Command, args []string) {
	appName := args[0]
	module_name := moduleName()
	err := os.MkdirAll(args[0], os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	parentDir := filepath.Dir(dir)
	// Read the file contents into a string
	serializers_name := "serializers.go"
	serializers_content := read_and_parse(parentDir + "/app/serializers.go")
	modifiedContent := strings.ReplaceAll(string(serializers_content), "app", appName)
	// Create the file
	create_file(appName+"/"+serializers_name, modifiedContent)

	// Read the file contents into a string
	models_name := "models.go"
	models_content := read_and_parse(parentDir + "/app/models.go")
	modifiedContent = strings.ReplaceAll(string(models_content), "app", appName)
	modifiedContent = strings.ReplaceAll(string(modifiedContent), "github.com/Zncl2222/gin-dj", module_name)
	// Create the file
	create_file(appName+"/"+models_name, modifiedContent)

	// Read the file contents into a string
	urls_name := "urls.go"
	urls_content := read_and_parse(parentDir + "/app/urls.go")
	modifiedContent = strings.ReplaceAll(string(urls_content), "app", appName)
	// Create the file
	create_file(appName+"/"+urls_name, modifiedContent)

	// Read the file contents into a string
	views_name := "views.go"
	views_content := read_and_parse(parentDir + "/app/views.go")
	modifiedContent = strings.ReplaceAll(string(views_content), "app", appName)
	// Create the file
	create_file(appName+"/"+views_name, modifiedContent)

	fmt.Printf("Successfully created folder structure for %s!\n", appName)
}
