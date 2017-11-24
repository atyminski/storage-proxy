package app

import(
	"io"
	"net/http"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Storage proxy server",
	Long: `Storage proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", hello)
		http.ListenAndServe(":8000", nil)
	},
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func Execute() {
	rootCmd.Execute()
}