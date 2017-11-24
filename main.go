package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yut-kt/goholiday"

	"github.com/skitn/duty/config"
)

func main() {

	var (
		version    bool
		configPath string
		startDate  string
		conf       config.Config
	)

	flag.BoolVar(&version, "version", false, "Print version information and quit")
	flag.StringVar(&configPath, "config", "", "Target toml file path for Config")
	flag.StringVar(&startDate, "start-date", "", "Rotation start date. format YYYY-MM-DD")
	flag.Parse()

	if version {
		showVersion()
		os.Exit(0)
	}

	conf, err := config.Load(configPath)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err))
		os.Exit(1)
	}

	holidays := []time.Time{}
	for i := 0; i < len(conf.CustomHolidays); i++ {
		datetime, err := time.Parse("2006-01-02", conf.CustomHolidays[i])
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
			os.Exit(1)
		}
		holidays = append(holidays, datetime)
	}
	goholiday.SetUniqueHolidays(holidays)

	var currentDatePointer *time.Time
	currentDate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err))
		os.Exit(1)
	}
	currentDatePointer = &currentDate

	var dutyDays = []string{}
	for i := 0; i < len(conf.Members); i++ {
		dutyDays = nil
		var j int
		for j < conf.DutyCount {
			if !goholiday.IsHoliday(currentDate) {
				dutyDays = append(dutyDays, currentDate.Format("01/02"))
				j++
			}
			*currentDatePointer = currentDatePointer.AddDate(0, 0, 1)
		}
		fmt.Printf("%s：%s\n", conf.Members[i], strings.Join(dutyDays, ","))
	}

	os.Exit(0)
}

func showVersion() {
	fmt.Println("duty version 0.0.1")
}
