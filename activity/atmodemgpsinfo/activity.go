package atmodemgpsinfo

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
	"strings"
)

// log is the default logger which we'll use to log
var log = logger.GetLogger("activity-at-modem-direct")

// String to hold the pointer for serial flag object
var serialPathP string
var baud int
var timeout time.Duration

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
	cmd := "+cgpsinfo"
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
        info, err := b.Command(ctx, cmd)
        cancel()
        log.Infof("AT" + cmd)
        if err != nil {
                log.Infof(" %s\n", err)
        } else {
                for _, l := range info {
        		log.Infof(" %s\n", l)
                }
        }

	log.Info(info)	
	tmpArray := strings.Split(info[0], ",")
	contextf.SetOutput("latitude", strings.TrimPrefix(tmpArray[0], strings.ToUpper(cmd)+": "))
	contextf.SetOutput("ns-indicator", tmpArray[1])
	contextf.SetOutput("longitude", tmpArray[2]) 
	contextf.SetOutput("ew-indicator", tmpArray[3])
	contextf.SetOutput("date", tmpArray[4])
        contextf.SetOutput("utctime", tmpArray[5])
	contextf.SetOutput("altitude", tmpArray[6])
	contextf.SetOutput("speed", tmpArray[7])
	contextf.SetOutput("course", tmpArray[8])
	return true, nil
}
