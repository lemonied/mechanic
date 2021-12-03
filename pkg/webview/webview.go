package webview

import (
	"fmt"
	"mechanic/src/config"
	"os/exec"
)

/*
Run open browser
*/
func Run() {
	website := fmt.Sprintf("http://localhost:%d", config.PORT)
	fmt.Println(website)
	cmd := exec.Command("cmd", "/c", "start", website)
	cmd.Start()
}
