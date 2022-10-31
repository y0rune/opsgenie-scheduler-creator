package main

import (
	"flag"
	"fmt"
	"os/exec"
	"testing"

	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

var scheduleID string
var scheduleClient *schedule.Client
var scheduleTest schedule.CreateResult
var apiKey *string = flag.String("apiKey", "", "# ApiKey for use in that script.\n# You can use the     export OPSGENIE_API_KEY=\"XXXXXXXXXXXXXXX\"")

const scheduleName string = "Testing_Schedule"
const scheduleTimezone string = "Europe/Warsaw"
const scheduleTeam string = "TestTeam"
const scheduleYear int = 2022
const scheduleEnabledFlag bool = false
const expetedNameOfRotation string = "w21-23.5-30.5"

func TestCreateApiClient(t *testing.T) {
	flag.Parse()
	scheduleClient = createApi(*apiKey)
}

func TestCreateSchedule(t *testing.T) {
	scheduleTest = scheduleCreator(*scheduleClient, scheduleName,
		scheduleTimezone, scheduleTeam,
		scheduleEnabledFlag)

	if scheduleTest.Name != scheduleName {
		t.Fatalf("Schedule has been NOT created correctly.")
	}
}

func TestCreateRestriction(t *testing.T) {
	restrictionCreator(*scheduleClient, scheduleTest.Id, scheduleYear)

	listRotation := getListRotation(*scheduleClient, scheduleTest.Id)
	if (listRotation.Rotations[20].Name) != expetedNameOfRotation {
		t.Fatalf("Schedule has been NOT created correctly.")
	}
}

func TestDeleteSchedule(t *testing.T) {
	deleteSchedule(*scheduleClient, scheduleTest.Id)
}

func TestFailedCreateScheduleCommand(t *testing.T) {
	cmd := exec.Command(
		"go", "run",
		"--scheduleTeam", scheduleTeam,
		"--scheduleName", scheduleName,
		"--scheduleYear", fmt.Sprint(scheduleYear),
	)

	err := cmd.Run()
	if err == nil {
		t.Fatalf("Command has been failed")
	}
}

func TestCreateViaGoRunScheduleCommand(t *testing.T) {
	cmd := exec.Command(
		"go", "run",
		"main.go",
		"--apiKey", *apiKey,
		"--scheduleTeam", scheduleTeam,
		"--scheduleName", scheduleName,
		"--scheduleYear", fmt.Sprint(scheduleYear),
		"--delete",
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}

func TestBuildCommand(t *testing.T) {
	cmd := exec.Command("make", "build")

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command has been failed")
	}
}

func TestCreateScheduleCommand(t *testing.T) {
	cmd := exec.Command(
		"./opsgenie-scheduler-rotation",
		"--apiKey", *apiKey,
		"--scheduleTeam", scheduleTeam,
		"--scheduleName", scheduleName,
		"--scheduleYear", fmt.Sprint(scheduleYear),
		"--delete",
	)

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}
