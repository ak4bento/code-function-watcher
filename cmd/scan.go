package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/ak4bento/code-function-watcher/pkg/scanner"
	"github.com/ak4bento/code-function-watcher/pkg/exporter"
)

var (
	outputPath string
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan directory for Go functions",
	Long:  `Scan the given directory recursively and extract all Go function declarations.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Println("❌ Failed to resolve path:", err)
			os.Exit(1)
		}

		fmt.Println("🔍  Scanning functions in:", absPath)

		functions, err := scanner.Scan(absPath)
		if err != nil {
			fmt.Println("❌ Failed to scan functions:", err)
			os.Exit(1)
		}

		if outputPath == "" {
			outputPath = "data/functions.json"
		}

		fmt.Println("💾  Exporting to:", outputPath)

		if err := exporter.ExportToJSON(functions, outputPath); err != nil {
			fmt.Println("❌ Failed to export:", err)
			os.Exit(1)
		}

		fmt.Printf("✅  Exported %d functions to %s\n", len(functions), outputPath)
	},
}

func init() {
	// rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file path for JSON export")
}

