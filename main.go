package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/yut-kt/goholiday"

	"github.com/skitn/duty/config"
)

func main() {

	var (
		version         bool
		configPath      string
		config          config.Config
	)

	flag.BoolVar(&version, "version", false, "Print version information and quit")
	flag.StringVar(&configPath, "config", "", "Target toml file path for Config")
	flag.Parse()

	if version {
		showVersion()
		os.Exit(0)
	}

	if configPath != "" {
		buf, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}

		if err := toml.Unmarshal(buf, &config); err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}

		holidays := []time.Time{}
		for i := 0; i < len(config.CustomHolidays); i++ {
			datetime, err := time.Parse("2006-01-02", config.CustomHolidays[i])
			if err != nil {
				fmt.Println(fmt.Errorf("error: %s", err))
				os.Exit(1)
			}
			holidays = append(holidays, datetime)
		}
		goholiday.SetUniqueHolidays(holidays)
	}

	var currentDatePointer *time.Time
	currentDate := time.Now()
	currentDatePointer = &currentDate
	for i := 0; i < len(config.Members); i++ {
		*currentDatePointer = currentDatePointer.AddDate(0, 0, 1)
		fmt.Println(config.Members[i])
		fmt.Println(currentDatePointer)
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
