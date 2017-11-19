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
		customHolidayPath string
	)

	flag.BoolVar(&version, "version", false, "Print version information and quit")
	flag.StringVar(&memberFilePath, "member", "", "Target json file path for Rotation member")
	flag.StringVar(&customHolidayPath, "holiday", "", "Target json file path for customize holiday")
	flag.Parse()

	if version {
		showVersion()
		os.Exit(0)
	}

	// TODO: require member json

	if customHolidayPath != "" {
		buf, err := ioutil.ReadFile(customHolidayPath)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}

		var custom_holiday_config config.CustomHolidayConfig
		err = json.Unmarshal(buf, &custom_holiday_config)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}

		holidays := []time.Time{}
		for i := 0; i < len(custom_holiday_config); i++ {
			datetime, err := time.Parse("2006-01-02", custom_holiday_config[i])
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
