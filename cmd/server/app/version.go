package app

import(
	"fmt"
	"github.com/spf13/cobra"
)

var Version string = "0.0.0.0";

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows version of app",
	Long: `Shows version of app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	 },
}