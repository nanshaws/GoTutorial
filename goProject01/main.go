package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "rename-tool",
	Short: "A tool to rename files based on prefixes or serial numbers",
}

var changePrefixCmd = &cobra.Command{
	Use:   "cp",
	Short: "Change the prefix of filenames in a directory",
	Run: func(cmd *cobra.Command, args []string) {
		directoryPath, _ := cmd.Flags().GetString("directory")
		oldPrefix, _ := cmd.Flags().GetString("old-prefix")
		newPrefix, _ := cmd.Flags().GetString("new-prefix")

		// 这里是原代码中的逻辑，现在封装在一个函数中
		changeFilePrefixes(directoryPath, oldPrefix, newPrefix)
	},
}

var changeNameBySerialCmd = &cobra.Command{
	Use:   "cn",
	Short: "Change filenames based on serial numbers",
	Run: func(cmd *cobra.Command, args []string) {
		directoryPath, _ := cmd.Flags().GetString("directory")
		baseNewName, _ := cmd.Flags().GetString("base-new-name")
		serialNumbers, _ := cmd.Flags().GetIntSlice("serial-numbers")

		// 这里是原代码中的逻辑，现在封装在一个函数中
		changeFileNamesBySerial(directoryPath, baseNewName, serialNumbers)
	},
}

func init() {
	rootCmd.AddCommand(changePrefixCmd)
	rootCmd.AddCommand(changeNameBySerialCmd)

	changePrefixCmd.Flags().StringP("directory", "d", "", "The directory to search")
	changePrefixCmd.Flags().StringP("old-prefix", "o", "", "The old prefix to replace")
	changePrefixCmd.Flags().StringP("new-prefix", "n", "", "The new prefix to use")

	changeNameBySerialCmd.Flags().StringP("directory", "d", "", "The directory to search")
	changeNameBySerialCmd.Flags().StringP("base-new-name", "b", "", "The base name to use for renaming")
	changeNameBySerialCmd.Flags().IntSliceP("serial-numbers", "s", []int{}, "The serial numbers to rename")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func changeFilePrefixes(directoryPath, oldPrefix, newPrefix string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入文件夹的名字:")
	scanner.Scan()
	directoryPath = scanner.Text()

	fmt.Println("请输入旧的前缀:")
	scanner.Scan()
	oldPrefix = scanner.Text()

	fmt.Println("请输入新的前缀:")
	scanner.Scan()
	newPrefix = scanner.Text()

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println(directoryPath, "不是一个目录。")
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), oldPrefix) {
			newName := strings.Replace(file.Name(), oldPrefix, newPrefix, 1)
			oldPath := filepath.Join(directoryPath, file.Name())
			newPath := filepath.Join(directoryPath, newName)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("重命名失败:", file.Name())
			} else {
				fmt.Println("重命名:", file.Name(), "->", newName)
			}
		}
	}
}

func changeFileNamesBySerial(directoryPath string, baseNewName string, serialNumbers []int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入文件夹的名字:")
	scanner.Scan()
	directoryPath = scanner.Text()

	fmt.Println("请输入要修改的个数:")
	var count int
	fmt.Scanln(&count)

	s := make([]int, count)
	fmt.Println("请依次输入要改的序列号:")
	for i := 0; i < count; i++ {
		fmt.Scanln(&s[i])
	}

	fmt.Println("请输入新的名字:")
	scanner.Scan()
	baseNewName = scanner.Text()

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println(directoryPath, "不是一个目录。")
	}

	for _, file := range files {
		for _, num := range s {
			suffix := fmt.Sprintf("%02d.png", num)
			if strings.HasSuffix(file.Name(), suffix) {
				newName := fmt.Sprintf("%s_%d.png", baseNewName, num)
				oldPath := filepath.Join(directoryPath, file.Name())
				newPath := filepath.Join(directoryPath, newName)
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("重命名失败:", file.Name())
				} else {
					fmt.Println("重命名:", file.Name(), "->", newName)
				}
			}
		}
	}
}
