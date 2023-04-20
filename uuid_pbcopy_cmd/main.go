package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"os/exec"
)

func main() {
	uuid1, err1 := uuid.NewV4()
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(uuid1.String() + " 가 복사되었습니다.")

	err2 := exec.Command("sh", "-c", `echo `+uuid1.String()+` | pbcopy`).Run()
	if err2 != nil {
		return
	}
}
