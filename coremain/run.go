// Package coremain contains the functions for starting CoreDNS.
package coremain

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
)

// CoremainStarter runs coremain
type CoremainStarter interface {
	Run()
	Init()
	GetLog() string
}

// NewCoreDns makes a new CoreDns
func NewCoreDns() *CoreDns {
	return &CoreDns{
		// corefilePath: corefilePath,
		status:       "СoreDns was defined",
	}
}

type CoreDns struct {
	// corefilePath string
	status       string
}

func (c *CoreDns) Init() {
	caddy.DefaultConfigFile = "Corefile"
	caddy.Quiet = true // don't show init stuff from caddy
	setVersion()

	// flag.StringVar(&conf, "conf", c.corefilePath, "Corefile to load (default \""+caddy.DefaultConfigFile+"\")")
	flag.BoolVar(&plugins, "plugins", false, "List installed plugins")
	flag.StringVar(&caddy.PidFile, "pidfile", "", "Path to write pid file")
	flag.BoolVar(&version, "version", false, "Show version")
	flag.BoolVar(&dnsserver.Quiet, "quiet", false, "Quiet mode (no initialization output)")

	caddy.RegisterCaddyfileLoader("flag", caddy.LoaderFunc(confLoader))
	caddy.SetDefaultCaddyfileLoader("default", caddy.LoaderFunc(defaultLoader))

	caddy.AppName = coreName
	caddy.AppVersion = CoreVersion
	c.status = "Init of the СoreDns finished"
}

// Run is CoreDNS's main() function.
func (c *CoreDns) Run() {
	c.status = "Run started"
	// time.Sleep(2*time.Second)

	caddy.TrapSignals()

	c.status = "TrapSignals finished"
	// time.Sleep(2*time.Second)

	// Reset flag.CommandLine to get rid of unwanted flags for instance from glog (used in kubernetes).
	// And read the ones we want to keep.
	flag.VisitAll(func(f *flag.Flag) {
		if _, ok := flagsBlacklist[f.Name]; ok {
			return
		}
		flagsToKeep = append(flagsToKeep, f)
	})

	c.status = "Reseted the flag.CommandLine"
	// time.Sleep(2*time.Second)

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	for _, f := range flagsToKeep {
		flag.Var(f.Value, f.Name, f.Usage)
	}

	c.status = "NewFlagSet finished"
	// time.Sleep(2*time.Second)

	flag.Parse()

	c.status = "Flags parsing finished"
	// time.Sleep(2*time.Second)

	if len(flag.Args()) > 0 {
		c.status = "Error: extra command line arguments"
		mustLogFatal(fmt.Errorf("extra command line arguments: %s", flag.Args()))
	}

	c.status = "Checking lenght args finished"
	// time.Sleep(2*time.Second)

	log.SetOutput(os.Stdout)
	log.SetFlags(0) // Set to 0 because we're doing our own time, with timezone

	c.status = "Setting flag to 0 finished"
	// time.Sleep(2*time.Second)

	if version {
		showVersion()
		os.Exit(0)
	}

	c.status = "Show version finished"
	// time.Sleep(2*time.Second)

	if plugins {
		fmt.Println(caddy.DescribePlugins())
		os.Exit(0)
	}

	c.status = "Show plugins finished"
	// time.Sleep(2*time.Second)

	// Get Corefile input
	// corefile, err := caddy.LoadCaddyfile(serverType)
	// if err != nil {
	// 	c.status = "Error: СoreDns didn't get Corefile"
	// 	mustLogFatal(err)
	// }
	// contentProtofile := ".:1253 {\n\t\tforward . 8.8.8.8\n\t\terrors\n\t\tdebug\n\t\tlog\n}"
	// fmt.Printf("contentProtofile: %v\n", []byte(contentProtofile))

	hardCorefile := caddy.CaddyfileInput { 
		Filepath: "corefile",
		Contents: []byte{46, 58, 49, 50, 53, 51, 32, 123, 10, 32, 32, 32, 32, 102, 111, 114, 119, 97, 114, 100, 32, 46,
			32, 56, 46, 56, 46, 56, 46, 56, 10, 10, 32, 32, 32, 32, 101, 114, 114, 111, 114, 115, 10, 32, 32, 32, 32,
			100, 101, 98, 117, 103, 10, 32, 32, 32, 32, 108, 111, 103, 10, 125, 10},
		ServerTypeName: "dns",
	}

	c.status = "СoreDns got Corefile"
	// time.Sleep(2*time.Second)

	// Start your engines
	instance, err := caddy.Start(hardCorefile)
	if err != nil {
		c.status = fmt.Sprintf("caddy.Start failed with error: %v", err)
		time.Sleep(10*time.Second)
		mustLogFatal(err)
	}

	c.status = "Engines started"
	// time.Sleep(1*time.Second)
	
	if !dnsserver.Quiet {
		showVersion()
	}
	c.status = "СoreDns started"

	// Twiddle your thumbs
	instance.Wait()
}

func (c *CoreDns) GetLog() string {
	return c.status
}

// mustLogFatal wraps log.Fatal() in a way that ensures the
// output is always printed to stderr so the user can see it
// if the user is still there, even if the process log was not
// enabled. If this process is an upgrade, however, and the user
// might not be there anymore, this just logs to the process
// log and exits.
func mustLogFatal(args ...interface{}) {
	if !caddy.IsUpgrade() {
		log.SetOutput(os.Stderr)
	}
	log.Fatal(args...)
}

// confLoader loads the Caddyfile using the -conf flag.
func confLoader(serverType string) (caddy.Input, error) {
	if conf == "" {
		return nil, nil
	}

	if conf == "stdin" {
		return caddy.CaddyfileFromPipe(os.Stdin, serverType)
	}

	contents, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, err
	}
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       conf,
		ServerTypeName: serverType,
	}, nil
}

// defaultLoader loads the Corefile from the current working directory.
func defaultLoader(serverType string) (caddy.Input, error) {
	contents, err := ioutil.ReadFile(caddy.DefaultConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: serverType,
	}, nil
}

// showVersion prints the version that is starting.
func showVersion() {
	fmt.Print(versionString())
	fmt.Print(releaseString())
	if devBuild && gitShortStat != "" {
		fmt.Printf("%s\n%s\n", gitShortStat, gitFilesModified)
	}
}

// versionString returns the CoreDNS version as a string.
func versionString() string {
	return fmt.Sprintf("%s-%s\n", caddy.AppName, caddy.AppVersion)
}

// releaseString returns the release information related to CoreDNS version:
// <OS>/<ARCH>, <go version>, <commit>
// e.g.,
// linux/amd64, go1.8.3, a6d2d7b5
func releaseString() string {
	return fmt.Sprintf("%s/%s, %s, %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version(), GitCommit)
}

// setVersion figures out the version information
// based on variables set by -ldflags.
func setVersion() {
	// A development build is one that's not at a tag or has uncommitted changes
	devBuild = gitTag == "" || gitShortStat != ""

	// Only set the appVersion if -ldflags was used
	if gitNearestTag != "" || gitTag != "" {
		if devBuild && gitNearestTag != "" {
			appVersion = fmt.Sprintf("%s (+%s %s)", strings.TrimPrefix(gitNearestTag, "v"), GitCommit, buildDate)
		} else if gitTag != "" {
			appVersion = strings.TrimPrefix(gitTag, "v")
		}
	}
}

// Flags that control program flow or startup
var (
	conf    string
	version bool
	plugins bool
)

// Build information obtained with the help of -ldflags
var (
	appVersion = "(untracked dev build)" // inferred at startup
	devBuild   = true                    // inferred at startup

	buildDate        string // date -u
	gitTag           string // git describe --exact-match HEAD 2> /dev/null
	gitNearestTag    string // git describe --abbrev=0 --tags HEAD
	gitShortStat     string // git diff-index --shortstat
	gitFilesModified string // git diff-index --name-only HEAD

	// Gitcommit contains the commit where we built CoreDNS from.
	GitCommit string
)

// flagsBlacklist removes flags with these names from our flagset.
var flagsBlacklist = map[string]struct{}{
	"logtostderr":      {},
	"alsologtostderr":  {},
	"v":                {},
	"stderrthreshold":  {},
	"vmodule":          {},
	"log_backtrace_at": {},
	"log_dir":          {},
}

var flagsToKeep []*flag.Flag
