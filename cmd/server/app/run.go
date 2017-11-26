package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/gevlee/storage-proxy/internal/pkg/handlers"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPath string
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs storage proxy server",
	Long:  `Runs storage proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func runServer() {
	port := viper.GetString("port")
	r := mux.NewRouter()
	initRouter(r)
	http.Handle("/", r)
	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Server run")
	logrus.Fatal(http.ListenAndServe(":"+port, nil))
}

func initRouter(r *mux.Router) {
	handler := createHandler()
	r.HandleFunc("/files/{context}/{fileid}", handler.Upload).Methods("POST")
	r.HandleFunc("/files/{context}/{fileid}", handler.Download).Methods("GET")
	r.HandleFunc("/files/{context}/{fileid}", handler.Remove).Methods("DELETE")
}

func createHandler() handlers.IHandler {
	return handlers.NewFileSystemHandler(viper.GetString("scope"))
}

func init() {
	cobra.OnInitialize(initConfig)
	runCmd.Flags().String("port", "8000", "Server port")
	runCmd.Flags().String("scope", ".", "Server port")
	runCmd.Flags().StringVar(&configPath, "config", "", "Path to config file")
	viper.BindPFlag("port", runCmd.Flags().Lookup("port"))
	viper.BindPFlag("scope", runCmd.Flags().Lookup("scope"))
}

func initConfig() {
	if configPath != "" {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}
	}
}
