package app

import(
	"io"
	"net/http"
	"github.com/spf13/cobra"
)

var(
	port string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs storage proxy server",
	Long: `Runs storage proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
	 	http.HandleFunc("/", hello)
		http.ListenAndServe(":" + port, nil)
	 },
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func init(){
	runCmd.PersistentFlags().StringVar(&port, "port", "8000", "Server port")
}