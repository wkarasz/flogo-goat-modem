package atmodemdirect

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
        "github.com/wkarasz/goat-modem/at"
	"github.com/wkarasz/goat-modem/serial"
	"github.com/wkarasz/goat-modem/trace"

        "flag"
	"log"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

// log is the default logger which we'll use to log
var log = logger.GetLogger("activity-at-modem-direct")

// String to hold the pointer for serial flag object
var serialPathP string

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	device := context.GetInput("devicePath").(string)
	cmd := context.GetInput("directCmd").(string)
        log.Infof("Device path capture [%s]", device)

        if flag.Lookup("serial") == nil {
                flag.StringVar(
                        &serialPathP,
                        "serial",
                        device,
			"Path to the serial device to use",
                )
        }
        baud := flag.Int("b", 115200, "baud rate")
	verbose := flag.Bool("v", false, "log modem interactions")
	flag.Parse()
	
	m, err := serial.New(*device, baud)
	if err != nil {
		log.Println(err)
		return
	}
	defer m.Close()
	var mio io.ReadWriter = m
	
	if *verbose {
		mio = trace.New(m, log.New(os.Stdout, "", log.LstdFlags))
	}
	a := at.New(mio)
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	err = a.Init(ctx)
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	cmds := []string{
		"I",
		"+GCAP",
		"+CMEE=2",
		"+CGMI",
		"+CGMM",
		"+CGMR",
		"+CGSN",
		"+CSQ",
		"+CIMI",
		"+CREG?",
		"+CNUM",
		"+CPIN?",
		"+CEER",
		"+CSCA?",
		"+CSMS?",
		"+CSMS=?",
		"+CPMS=?",
		"+CNMI?",
		"+CNMI=?",
		"+CNMA=?",
		"+CMGF=?",
	}
	for _, cmd := range cmds {
		ctx, cancel := context.WithTimeout(context.Background(), *timeout)
		info, err := a.Command(ctx, cmd)
		cancel()
		fmt.Println("AT" + cmd)
		if err != nil {
			fmt.Printf(" %s\n", err)
		} else {
			for _, l := range info {
				fmt.Printf(" %s\n", l)
			}
		}
	}

	context.SetOutput("result", info)

	return true, nil
}
