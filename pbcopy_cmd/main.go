package main

import "os/exec"

func main() {
	err := exec.Command("sh", "-c", `echo "hello world" | pbcopy`).Run()
	if err != nil {
		return
	}
}
