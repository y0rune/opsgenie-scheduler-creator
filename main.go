package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"github.com/sirupsen/logrus"
)

var startHourWorkWeek uint32 = 17
var startMinuteWorkWeek uint32 = 0
var endHourWorkWeek uint32 = 9
var endMinWorkWeek uint32 = 0

const staticScheduleName string = "Test Schedule"
const staticScheduleID string = "XXXXXXXXXXXXXXX"
const staticScheduleTimezone string = "Europe/Warsaw"
const staticScheduleTeam string = "TestTeam"
const staticScheduleYear int = 2022
const staticScheduleEnabledFlag bool = true

var defaultSchedule = [...]og.Restriction{
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
}

func createApi(apiKey string) *schedule.Client {
	apiKeyEnv := os.Getenv("OPSGENIE_API_KEY")

	if apiKeyEnv != "" {
		apiKey = apiKeyEnv
	} else if apiKey == "" {
		fmt.Printf("Empty apiKey...\nPlease use -apiKey or    export OPSGENIE_API_KEY=\"XXXXXXXXXXXXXXX\" \n")
		os.Exit(1)
	}

	scheduleClient, err := schedule.NewClient(&client.Config{
		ApiKey: apiKey,
		LogLevel: logrus.ErrorLevel,
	})

	if err != nil {
		fmt.Printf("Error in scheduleClient create: %d", err)
		os.Exit(1)
	}

	return scheduleClient
}

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

func scheduleCreator(scheduleClient schedule.Client, scheduleName string, scheduleTimezone string, scheduleTeam string, scheduleEnabledFlag bool) schedule.CreateResult {
	scheduleResult, err := scheduleClient.Create(nil, &schedule.CreateRequest{
		Name:     scheduleName,
		Timezone: scheduleTimezone,
		Enabled:  &scheduleEnabledFlag,
		OwnerTeam: &og.OwnerTeam{
			Name: scheduleTeam,
		},
	})

	if err != nil {
		fmt.Printf("Schedule %s with id: %s has been NOT created. Error: %d \n", scheduleResult.Name, scheduleResult.Id, err)
		os.Exit(1)
	} else {
		fmt.Printf("Schedule %s with id: %s has been created.\n", scheduleResult.Name, scheduleResult.Id)
	}

	return *scheduleResult
}

func restrictionCreator(scheduleClient schedule.Client, scheduleID string, year int) {
	month := time.Month(1)
	firstMonday := getFirstMonday(year, month)
	numberOfWeeks := getNumberOfWeeks(year, month)

	nextMonday := time.Date(year, month, int(firstMonday), 1, 0, 0, 0, time.UTC)
	for week := 1; week <= numberOfWeeks; week++ {
		monday := nextMonday
		nextMonday = nextMonday.AddDate(0, 0, 7)
		weekName := fmt.Sprintf("w%d-%d.%d-%d.%d", week, monday.Day(), monday.Month(), nextMonday.Day(), nextMonday.Month())

		_, err := scheduleClient.CreateRotation(nil, &schedule.CreateRotationRequest{
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
					Type:            og.WeekdayAndTimeOfDay,
					RestrictionList: defaultSchedule[:],
				},
			},
			ScheduleIdentifierType:  schedule.Id,
			ScheduleIdentifierValue: scheduleID,
		})

		if err != nil {
			fmt.Printf("Rotation %s has been NOT created for schedule %s.\n", weekName, scheduleID)
		} else {
			fmt.Printf("Rotation %s has been created for schedule %s.\n", weekName, scheduleID)
		}
	}
}

func deleteSchedule(scheduleClient schedule.Client, scheduleID string) {
	time.Sleep(10 * time.Second)

	_, err := scheduleClient.Delete(nil, &schedule.DeleteRequest{
		IdentifierType:  schedule.Id,
		IdentifierValue: scheduleID,
	})

	if err != nil {
		fmt.Printf("Schedule %s has been NOT deleted.\n", scheduleID)
	} else {
		fmt.Printf("Schedule %s has been deleted.\n", scheduleID)
	}
}

func getListRotation(scheduleClient schedule.Client, scheduleID string) *schedule.ListRotationsResult {
	scheduleResult, err := scheduleClient.ListRotations(nil, &schedule.ListRotationsRequest{
		ScheduleIdentifierType:  schedule.Id,
		ScheduleIdentifierValue: scheduleID,
	})

	if err != nil {
		fmt.Printf("Schedule %s can not be get.\n", scheduleID)
	}
	return scheduleResult
}

func main() {
	apiKey := flag.String("apiKey", "", "# ApiKey for use in that script.\n# You can use the     export OPSGENIE_API_KEY=\"XXXXXXXXXXXXXXX\"")
	scheduleName := flag.String("scheduleName", staticScheduleName, "# Name of schedule")
	scheduleID := flag.String("scheduleID", staticScheduleID, "# ID of schedule")
	scheduleTimezone := flag.String("scheduleTimezone", staticScheduleTimezone, "# Timezone of the schedule")
	scheduleTeam := flag.String("scheduleTeam", staticScheduleTeam, "# Name of the team in the schedule")
	scheduleYear := flag.Int("scheduleYear", staticScheduleYear, "# Year of the schedule")
	scheduleEnabledFlag := flag.Bool("scheduleEnabledFlag", staticScheduleEnabledFlag, "# Schedule is enabled")
	delete := flag.Bool("delete", false, "# Delete schedule ")
	flag.Parse()

	scheduleClient := createApi(*apiKey)

	if *scheduleName != staticScheduleName {
		createdSchedule := scheduleCreator(*scheduleClient, *scheduleName, *scheduleTimezone, *scheduleTeam, *scheduleEnabledFlag)
		restrictionCreator(*scheduleClient, createdSchedule.Id, *scheduleYear)
		if *delete {
			scheduleID = &createdSchedule.Id
			deleteSchedule(*scheduleClient, *scheduleID)
			os.Exit(0)
		}
	}

	if *delete && *scheduleID != staticScheduleID {
		deleteSchedule(*scheduleClient, *scheduleID)
		os.Exit(0)
	}
}
