package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
	"io/ioutil"

	"github.com/yut-kt/goholiday"

	"github.com/skitn/duty/config"
)

func main() {

	var (
		version bool
		memberFilePath string
		customHolidayFilePath string
		memberConfig config.MemberConfig
	)

	flag.BoolVar(&version, "version", false, "Print version information and quit")
	flag.StringVar(&memberFilePath, "member", "", "Target json file path for Rotation member")
	flag.StringVar(&customHolidayFilePath, "holiday", "", "Target json file path for customize holiday")
	flag.Parse()

	if version {
		showVersion()
		os.Exit(0)
	}

	if memberFilePath == "" {
		fmt.Println("Require member option")
		os.Exit(1)
	} else {
		buf, err := ioutil.ReadFile(memberFilePath)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}

		err = json.Unmarshal(buf, &memberConfig)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}
	}

	if customHolidayFilePath != "" {
		buf2, err := ioutil.ReadFile(customHolidayFilePath)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}

		var customHolidayConfig config.CustomHolidayConfig
		err = json.Unmarshal(buf2, &customHolidayConfig)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}

		holidays := []time.Time{}
		for i := 0; i < len(customHolidayConfig); i++ {
			datetime, err := time.Parse("2006-01-02", customHolidayConfig[i])
			if err != nil {
				fmt.Println(fmt.Errorf("error: %s", err))
				os.Exit(1)
			}
			holidays = append(holidays, datetime)
		}
		goholiday.SetUniqueHolidays(holidays)
	}

	// check holiday
	NYD := "2017-11-17"
	datetime, err := time.Parse("2006-01-02", NYD)
	if err != nil {
		panic(err)
	}
	fmt.Println(goholiday.IsHoliday(datetime))

	// rotation start
	os.Exit(0)
}

func showVersion() {
	fmt.Println("duty version 0.0.1")
}
