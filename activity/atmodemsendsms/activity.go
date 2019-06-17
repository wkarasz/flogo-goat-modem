package atmodemsendsms

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
        "github.com/wkarasz/goat-modem/at"
	"github.com/wkarasz/goat-modem/serial"
//	"github.com/wkarasz/goat-modem/trace"

        "flag"
//	"log"
	"fmt"
	"io"
//	"os"
	"context"
	"time"
)

// log is the default logger which we'll use to log
var log = logger.GetLogger("activity-at-modem-send-sms")

// String to hold the pointer for serial flag object
var serialPathP string
var timeout time.Duration
var baud int

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
func (a *MyActivity) Eval(contextf activity.Context) (done bool, err error)  {

	// do eval
	device := contextf.GetInput("devicePath").(string)
	cmd := contextf.GetInput("directCmd").(string)
	phoneNumber := contextf.GetInput("recipientMobile").(string)
	message := contextf.GetInput("message").(string)
        log.Infof("Device path capture [%s]", device)

        if flag.Lookup("serial") == nil {
                flag.StringVar(
                        &serialPathP,
                        "serial",
                        device,
			"Path to the serial device to use",
                )
		flag.IntVar(&baud, "baud", int(115200), "baud rate")
		flag.DurationVar(&timeout, "t", 400*time.Millisecond, "command timeout period")
        }
	//verbose := flag.Bool("v", false, "log modem interactions")
	flag.Parse()
	
	m, err := serial.New(device, baud)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer m.Close()
	var mio io.ReadWriter = m
	
	//if *verbose {
	//	mio = trace.New(m, fmt.New(os.Stdout, "", log.LstdFlags))
	//}
	b := at.New(mio)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	err = b.Init(ctx)
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), timeout)
        info, err := b.SMSCommand(ctx, "+cmgs=\"" + phoneNumber + "\"", message)
        log.Infof("AT" + cmd)
        if err != nil {
                log.Infof(" %s\n", err)
        } else {
                for _, l := range info {
        		log.Infof(" %s\n", l)
                }
        }

	contextf.SetOutput("result", info)

	return true, nil
}
