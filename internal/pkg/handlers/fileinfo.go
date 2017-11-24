package handlers

import(
	"net/http"
	"github.com/gorilla/mux"
)

type FileInfo struct {
	Extension string
	Context string
	Identifier string
}

func createFileInfoFromRequest(r *http.Request) *FileInfo {
	vars := mux.Vars(r)
	ext := r.URL.Query().Get("ext")

	return &FileInfo{
		Context : vars["context"],
		Identifier : vars["fileid"],
		Extension : ext,
	}
}

func (info FileInfo) HasExtension() bool {
	return info.Extension != ""
}

func (info FileInfo) FileNameWithExt() string {
	if info.HasExtension() {
		return info.Identifier + "." + info.Extension
	}

	return info.Identifier
}