package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func osExecWithTimeout(cmdBin string, args ...string) ([]byte, error) {

	ctxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctxt, cmdBin, args...)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Start(); err != nil {
		return buf.Bytes(), err
	}

	if err := cmd.Wait(); err != nil {
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}

func main() {
	res, err := osExecWithTimeout("/bin/sleep", "100")
	if err != nil {
		fmt.Printf("sleep error :%v", err)
	} else {
		fmt.Printf("sleep res: %v\n", string(res))
	}
}
