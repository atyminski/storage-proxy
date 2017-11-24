package app

import(
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Storage proxy server",
	Long: `Storage proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
		 cmd.Help();
	 },
}

func init(){
	rootCmd.AddCommand(runCmd)
}

func Execute() {
	rootCmd.Execute()
}