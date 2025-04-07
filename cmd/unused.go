package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ak4bento/code-function-watcher/pkg/scanner"
	"github.com/ak4bento/code-function-watcher/pkg/unused"
	"github.com/ak4bento/code-function-watcher/pkg/utils"
)

var unusedCmd = &cobra.Command{
	Use:   "unused <path>",
	Short: "Find unused functions in the project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := args[0]

		// 🧠 Scan all functions in the target directory
		functions, err := scanner.Scan(targetPath)
		if err != nil {
			fmt.Println("❌ Failed to scan functions:", err)
			os.Exit(1)
		}

		// 📥 Load ignore list from .funcignore
		ignoreList, err := utils.LoadIgnoreList(".funcignore")
		if err != nil {
			fmt.Println("❌ Failed to load .funcignore:", err)
			os.Exit(1)
		}

		// ➕ Load additional ignores from flag --ignore
		flagIgnores, _ := cmd.Flags().GetStringSlice("ignore")
		for _, name := range flagIgnores {
			ignoreList[name] = struct{}{}
		}

		// 🔍 Find unused
		unusedFuncs, err := unused.FindUnusedFunctions(targetPath, functions, ignoreList)
		if err != nil {
			fmt.Println("❌ Failed to analyze unused functions:", err)
			os.Exit(1)
		}

		// 🖨️ Output
		if len(unusedFuncs) == 0 {
			fmt.Println("✅ No unused functions found.")
		} else {
			fmt.Printf("⚠️  Found %d unused function(s):\n", len(unusedFuncs))
			for _, fn := range unusedFuncs {
				fmt.Printf("- %s (%s)\n", fn.Name, fn.File)
			}
		}
	},
}

func init() {
	unusedCmd.Flags().StringSliceP("ignore", "i", []string{}, "Function names to ignore (comma-separated or repeated)")
	rootCmd.AddCommand(unusedCmd)
}

