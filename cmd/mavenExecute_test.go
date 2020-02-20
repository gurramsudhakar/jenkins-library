package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/SAP/jenkins-library/pkg/command"
)

func TestMavenExecute(t *testing.T) {
	t.Run("maven version", func(t *testing.T) {
		opts := mavenExecuteOptions{}
		c := command.Command{}
		stdOutBuf := new(bytes.Buffer)

		outfile, err := os.Create("test.txt")
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		defer outfile.Close()

		stdOut := io.MultiWriter(outfile, stdOutBuf)
		c.Stdout(stdOut)
		err = runMavenExecute(&opts, &c)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		t.Logf("my buffer: %v", string(stdOutBuf.Bytes()))

	})
}
