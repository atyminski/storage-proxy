package handlers

import (
	"io"
	"net/http"
	"os"
	"path"
)

type FileSystemHandler struct{
	ScopeDir string
}

func NewFileSystemHandler(scopeDir string) *FileSystemHandler {
	err := createDir(scopeDir)
	if err != nil {
		panic(err)
	}
	return &FileSystemHandler{
		ScopeDir : scopeDir,
	}
}

func createDir(dir string) error{
	_, err := os.Stat(dir)

	if err == nil {
		return err
	}

	if os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		return err
	}

	return nil
}

func (handler FileSystemHandler) Upload(w http.ResponseWriter, r *http.Request){
	fileinfo := createFileInfoFromRequest(r)
	destDir := path.Join(handler.ScopeDir, fileinfo.Context)
	err := createDir(destDir)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		NewErrorMessage(err.Error()).AsJsonResponse(w)
	}

	destFile := path.Join(destDir, fileinfo.FileNameWithExt())

	file, err := os.Create(destFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		NewErrorMessage(err.Error()).AsJsonResponse(w)
	}

	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		NewErrorMessage(err.Error()).AsJsonResponse(w)
	}
}
func (FileSystemHandler) Download(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func (FileSystemHandler) Remove(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusMethodNotAllowed)
}