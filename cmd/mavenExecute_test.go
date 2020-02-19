package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/SAP/jenkins-library/pkg/command"
)

 type myWriter struct (
	func Write(p []byte)  {

	}
)

func TestMavenExecute(t *testing.T) {
	t.Run("maven version", func(t *testing.T) {
		opts := mavenExecuteOptions{}
		c := command.Command{}
		stdOutBuf := new(bytes.Buffer)
		stdOut := io.MultiWriter(os.Stderr, stdOutBuf)
		c.Stdout(stdOut)
		t.Log("test")
		err := runMavenExecute(opts, &c)
		if err != nil {
			fmt.Println("err")
		}
		t.Logf ("my buffer: %v",string(stdOutBuf.Bytes()))

	})
}
