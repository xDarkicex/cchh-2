package compile

import (
	"fmt"
	"log"
	"os/exec"
)

//CompileAssets Asset Pipeline
func Assets() {
	compiled := make(chan bool)
	go func() {
		err := exec.Command(
			"sass",
			"--watch",
			"./app/assets/css/:./public/pipeline/stylesheets/", "--style", "compressed").Start()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Sass Assets Compiled")

		compiled <- true
		close(compiled)
	}()
}
