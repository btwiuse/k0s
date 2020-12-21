package gitd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	service := FindService(r)
	execPath := "."
	if ok := IsExistingRepository(execPath); !ok {
		log.Println("repository not found")
		http.Error(w, "repository not found", http.StatusNotFound)
		return
	}

	cmd := exec.Command("git", service, "--stateless-rpc", "--advertise-refs", execPath)
	_, stdout, stderr, ok := GetChildPipes(cmd, w)
	if !ok {
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println("Error while spawning:", err)
		http.Error(w, "Error while spawning", http.StatusInternalServerError)
		return
	}

	contentType := fmt.Sprintf("application/x-git-%s-advertisement", service)
	SetHeader(w, contentType)

	w.Write([]byte(CreateFirstPKTLine(service)))
	go io.Copy(w, stdout)
	go io.Copy(w, stderr)
	if err := cmd.Wait(); err != nil {
		log.Println("Error while waiting:", err)
		return
	}
}
