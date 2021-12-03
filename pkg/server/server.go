package server

import (
	"fmt"
	"mechanic/src/config"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

/*
Run run web server
*/
func Run() {
	box := packr.New("Assets", "../../assets")
	mux := http.NewServeMux()

  mux.Handle("/", http.FileServer(box))
	httpErr := http.ListenAndServe(fmt.Sprintf("localhost:%d", config.PORT), mux)
	if httpErr != nil {
		fmt.Println(httpErr)
	}
}
