package caller

import (
	"fmt"
	"log"
	"os"
)

const (
	defaultLogOutputPath = `coredns/core.log`
)

var (
	oldStdout = os.Stdout
	oldStderr = os.Stderr
)

type Logger struct {
	log *os.File
}

func (c *Logger) Printf(frmt string, args ...interface{}) {
	fmt.Fprintf(oldStdout, frmt, args...)
}

func (c *Logger) Log(msg string) {
	log.Println(msg)
}

// SetLogOutput function sets the ouput file for the logging instead of defined
// stdout/stderr.
func (c *Logger) SetLogOutput(p string) {

	var err error

	if p == "" {
		p = defaultLogOutputPath
	}

	if c.log, err = os.OpenFile(p, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644); err != nil {
		panic(err)
	}

	os.Stdout, os.Stderr = c.log, c.log
	log.SetOutput(c.log)
}

// ResetLogOutput resets the output file for the log by reverting back stdout/stderr.
func (c *Logger) ResetLogOutput() {
	if c.log != nil {
		c.log.Close()
	}
	os.Stdout, os.Stderr = oldStdout, oldStderr
	log.SetOutput(os.Stdout)
}
