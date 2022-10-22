package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

var startHourWorkWeek uint32 = 17
var startMinuteWorkWeek uint32 = 0
var endHourWorkWeek uint32 = 9
var endMinWorkWeek uint32 = 0

func getFirstMonday(year int, month time.Month) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	firstMonday := ((8-int(t.Weekday()))%7 + 1)

	return firstMonday
}

func getNumberOfWeeks(year int, month time.Month) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	_, numberOfWeeks := t.ISOWeek()

	return numberOfWeeks
}

func scheduleCreator(scheduleClient schedule.Client, scheduleName string, scheduleTimezone string, scheduleTeam string, scheduleEnabled bool) schedule.CreateResult {
	scheduleResult, err := scheduleClient.Create(nil, &schedule.CreateRequest{
		Name:     scheduleName,
		Timezone: scheduleTimezone,
		Enabled:  &scheduleEnabled,
		OwnerTeam: &og.OwnerTeam{
			Name: scheduleTeam,
		},
	})

	if err != nil {
		fmt.Printf("Error in scheduleCreator create: %d", err)
	}

	return *scheduleResult
}

func restrictionCreator(scheduleClient schedule.Client, scheduleName string, year int) {
	month := time.Month(1)
	firstMonday := getFirstMonday(year, month)
	numberOfWeeks := getNumberOfWeeks(year, month)

	nextMonday := time.Date(year, month, int(firstMonday), 1, 0, 0, 0, time.UTC)
	for week := 1; week <= numberOfWeeks; week++ {
		monday := nextMonday
		nextMonday = nextMonday.AddDate(0, 0, 7)
		weekName := fmt.Sprintf("w%d-%d.%d-%d.%d", week, monday.Day(), monday.Month(), nextMonday.Day(), nextMonday.Month())

		scheduleClient.CreateRotation(nil, &schedule.CreateRotationRequest{
			Rotation: &og.Rotation{
				Name:      weekName,
				StartDate: &monday,
				EndDate:   &nextMonday,
				Type:      og.Weekly,
				Participants: []og.Participant{
					{
						Type: og.None,
					},
				},
				TimeRestriction: &og.TimeRestriction{
					Type: og.WeekdayAndTimeOfDay,
					RestrictionList: []og.Restriction{
						{
							StartDay:  "monday",
							EndDay:    "tuesday",
							EndHour:   &endHourWorkWeek,
							EndMin:    &endMinWorkWeek,
							StartHour: &startHourWorkWeek,
							StartMin:  &startMinuteWorkWeek,
						},
						{
							StartDay:  "tuesday",
							EndDay:    "wednesday",
							EndHour:   &endHourWorkWeek,
							EndMin:    &endMinWorkWeek,
							StartHour: &startHourWorkWeek,
							StartMin:  &startMinuteWorkWeek,
						},
						{
							StartDay:  "wednesday",
							EndDay:    "thursday",
							EndHour:   &endHourWorkWeek,
							EndMin:    &endMinWorkWeek,
							StartHour: &startHourWorkWeek,
							StartMin:  &startMinuteWorkWeek,
						},
						{
							StartDay:  "thursday",
							EndDay:    "friday",
							EndHour:   &endHourWorkWeek,
							EndMin:    &endMinWorkWeek,
							StartHour: &startHourWorkWeek,
							StartMin:  &startMinuteWorkWeek,
						},
						{
							StartDay:  "friday",
							EndDay:    "monday",
							EndHour:   &endHourWorkWeek,
							EndMin:    &endMinWorkWeek,
							StartHour: &startHourWorkWeek,
							StartMin:  &startMinuteWorkWeek,
						},
					},
				},
			},
			ScheduleIdentifierType:  schedule.Name,
			ScheduleIdentifierValue: scheduleName,
		})

		fmt.Printf("Rotation %s has been created for schedule %s.\n", weekName, scheduleName)
	}
}

func main() {

	apiKey := flag.String("apiKey", "", "# ApiKey for use in that script")
	scheduleName := flag.String("scheduleName", "Test Schedule", "# Name of schedule")
	scheduleTimezone := flag.String("scheduleTimezone", "Europe/Warsaw", "# Timezone of the schedule")
	scheduleTeam := flag.String("scheduleTeam", "TestTeam", "# Name of the team in the schedule")
	scheduleYear := flag.Int("scheduleYear", 2022, "# Year of the schedule")
	scheduleEnabled := flag.Bool("scheduleEnabled", true, "# Schedule is enabled")
	flag.Parse()

	if (*apiKey == "") || (apiKey == nil) {
		fmt.Printf("Empty apiKey... Please use -apiKey \n")
		os.Exit(1)
	}

	scheduleClient, err := schedule.NewClient(&client.Config{
		ApiKey: *apiKey,
	})

	if err != nil {
		fmt.Printf("Error in scheduleClient create: %d", err)
	}

	createdSchedule := scheduleCreator(*scheduleClient, *scheduleName, *scheduleTimezone, *scheduleTeam, *scheduleEnabled)
	restrictionCreator(*scheduleClient, createdSchedule.Name, *scheduleYear)
}
