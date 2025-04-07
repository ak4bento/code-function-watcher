package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ak4bento/code-function-watcher/pkg/exporter"
	"github.com/ak4bento/code-function-watcher/pkg/compare"
)

var threshold int

var compareCmd = &cobra.Command{
	Use:   "compare [old.json] [new.json]",
	Short: "Compare two sets of functions and detect similar ones",
	Long:  `Compare previously scanned functions and detect potentially duplicated or similar functions.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		oldPath := args[0]
		newPath := args[1]

		oldFuncs, err := exporter.LoadFromJSON(oldPath)
		if err != nil {
			fmt.Println("❌ Failed to load old file:", err)
			os.Exit(1)
		}

		newFuncs, err := exporter.LoadFromJSON(newPath)
		if err != nil {
			fmt.Println("❌ Failed to load new file:", err)
			os.Exit(1)
		}

		dupes := compare.Compare(oldFuncs, newFuncs, threshold)

		if len(dupes) == 0 {
			fmt.Println("✅ No similar functions found.")
		} else {
			fmt.Printf("⚠️  Found %d potentially similar functions:\n", len(dupes))
			for _, d := range dupes {
				fmt.Printf("🔁  %s:%d  ⇄  %s:%d  (%.2f%%)\n",
					d.FuncA.File, d.FuncA.Position,
					d.FuncB.File, d.FuncB.Position,
					d.Similarity*100,
				)
			}
		}
	},
}

func init() {
	// rootCmd.AddCommand(compareCmd)

	compareCmd.Flags().IntVarP(&threshold, "threshold", "t", 70, "Similarity threshold percentage")
}

