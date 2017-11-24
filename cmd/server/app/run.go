package app

import(
	"net/http"
	"github.com/spf13/cobra"
	"github.com/gorilla/mux"
	"github.com/gevlee/storage-proxy/internal/pkg/handlers"
)

var(
	port, scope string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs storage proxy server",
	Long: `Runs storage proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer();
	 },
}

func runServer() {
	handler := createHandler();
	r := mux.NewRouter()
	r.HandleFunc("/files/{context}/{fileid}", handler.Upload).Methods("POST")
	r.HandleFunc("/files/{context}/{fileid}", handler.Download).Methods("GET")
	r.HandleFunc("/files/{context}/{fileid}", handler.Remove).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":" + port, nil)
}

func createHandler() handlers.IHandler{
	return handlers.NewFileSystemHandler(scope)
}

func init(){
	runCmd.PersistentFlags().StringVar(&port, "port", "8000", "Server port")
	runCmd.PersistentFlags().StringVar(&scope, "scope", ".", "Server port")
}