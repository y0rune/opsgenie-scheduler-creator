package main

import (
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"testing"

	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
)

var scheduleID string
var scheduleClient *schedule.Client
var scheduleTest schedule.CreateResult
var apiKey *string = flag.String("apiKey", "", "# ApiKey for use in that script.\n# You can use the     export OPSGENIE_API_KEY=\"XXXXXXXXXXXXXXX\"")

var teamID string
var teamClient *team.Client
var teamTest *team.CreateTeamResult

const scheduleName string = "Testing_Schedule"
const scheduleTimezone string = "Europe/Warsaw"
const scheduleTeam string = "TestTeam"
const scheduleYear int = 2022
const scheduleEnabledFlag bool = false
const expetedNameOfRotation string = "w21-23.5-30.5"

const teamName string = "TestTeam"
const teamDesc string = "Test"

func TestCleningApp(t *testing.T) {
	cmd := exec.Command("make", "clean")

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}

func TestBuilingApp(t *testing.T) {
	cmd := exec.Command("make", "build")

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}

// Creating the Clients
func TestCreateApiClient(t *testing.T) {
	flag.Parse()
	scheduleClient = createApi(*apiKey)
}

func TestCreateTeamClient(t *testing.T) {
	teamClient = createTeamClient(*apiKey)
}

// Test One:
// - Create a testTeam via function
// - Create a testSchedule via function
// - Create a restriction via function
// - Delete a testschedule via function
// - Delete a testTeam via function

func TestOneCreateTestTeam(t *testing.T) {
	teamTest = teamCreator(*teamClient, teamName, teamDesc)
	if teamTest.Name != teamName {
		t.Fatalf("Team has been NOT created correctly.")
	}
}

func TestOneCreateSchedule(t *testing.T) {
	scheduleTest = scheduleCreator(*scheduleClient, scheduleName,
		scheduleTimezone, scheduleTeam,
		scheduleEnabledFlag)

	if scheduleTest.Name != scheduleName {
		t.Fatalf("Schedule has been NOT created correctly.")
	}
}

func TestOneCreateRestriction(t *testing.T) {
	restrictionCreator(*scheduleClient, scheduleTest.Id, scheduleYear)

	listRotation := getListRotation(*scheduleClient, scheduleTest.Id)
	if (listRotation.Rotations[20].Name) != expetedNameOfRotation {
		t.Fatalf("Schedule has been NOT created correctly.")
	}
}

func TestOneDeleteSchedule(t *testing.T) {
	deleteSchedule(*scheduleClient, scheduleTest.Id)
}

func TestOneDeleteTeam(t *testing.T) {
	deleteTeam(*teamClient, teamTest.Id)
}

// Test Two:
// - Create a testTeam via go run
// - Create a testSchedule via go run
// - Delete a testschedule via go run
// - Delete a testTeam via go run

func TestTwoCreateTestTeam(t *testing.T) {
	cmd := exec.Command(
		"go", "run",
		"main.go",
		"--apiKey", *apiKey,
		"--teamName", teamName,
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}

	r, _ := regexp.Compile("[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+")
	output, _ := cmd.CombinedOutput()
	teamID = r.FindString(string(output))
}

func TestTwoCreateSchedule(t *testing.T) {
	cmd := exec.Command(
		"go", "run",
		"main.go",
		"--apiKey", *apiKey,
		"--scheduleTeam", scheduleTeam,
		"--scheduleName", scheduleName,
		"--scheduleYear", fmt.Sprint(scheduleYear),
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}

	r, _ := regexp.Compile("[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+")
	output, _ := cmd.CombinedOutput()
	scheduleID = r.FindString(string(output))
}

func TestTwoDeleteSchedule(t *testing.T) {
	cmd := exec.Command(
		"go", "run",
		"main.go",
		"--apiKey", *apiKey,
		"--scheduleID", scheduleID,
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}

func TestTwoDeleteTeam(t *testing.T) {
	cmd := exec.Command(
		"go", "run",
		"main.go",
		"--apiKey", *apiKey,
		"--teamID", teamID,
		"--delete",
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}

// Test Three:
// - Create a testTeam via builded app
// - Create a testSchedule via builded app
// - Delete a testschedule via builded app
// - Delete a testTeam via builded app

func TestThreeCreateTestTeam(t *testing.T) {
	apiKey := checkApiKey(*apiKey)

	cmd := exec.Command(
		"./opsgenie-scheduler-creator",
		"--apiKey", apiKey,
		"--teamName", teamName,
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}

	r, _ := regexp.Compile("[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+")
	output, _ := cmd.CombinedOutput()
	teamID = r.FindString(string(output))
}

func TestThreeCreateSchedule(t *testing.T) {
	apiKey := checkApiKey(*apiKey)
	cmd := exec.Command(
		"./opsgenie-scheduler-creator",
		"--apiKey", apiKey,
		"--scheduleTeam", scheduleTeam,
		"--scheduleName", scheduleName,
		"--scheduleYear", fmt.Sprint(scheduleYear),
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}

	r, _ := regexp.Compile("[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+")
	output, _ := cmd.CombinedOutput()
	scheduleID = r.FindString(string(output))
}

func TestThreeDeleteSchedule(t *testing.T) {
	apiKey := checkApiKey(*apiKey)
	cmd := exec.Command(
		"./opsgenie-scheduler-creator",
		"--apiKey", apiKey,
		"--scheduleID", scheduleID,
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}

func TestThreeDeleteTeam(t *testing.T) {
	apiKey := checkApiKey(*apiKey)
	cmd := exec.Command(
		"go", "run",
		"main.go",
		"--apiKey", apiKey,
		"--teamID", teamID,
		"--delete",
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}

func TestFourCreateSchedule(t *testing.T) {
	apiKey := checkApiKey(*apiKey)
	cmd := exec.Command(
		"go", "run",
		"main.go",
		"--apiKey", apiKey,
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

// Test Five:
// - Create and delete testTeam via builded app
// - Create and delete a testSchedule via builded app

func TestFiveCreateTestTeam(t *testing.T) {
	apiKey := checkApiKey(*apiKey)
	cmd := exec.Command(
		"./opsgenie-scheduler-creator",
		"--apiKey", apiKey,
		"--teamName", teamName,
		"--delete",
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}

func TestFiveCreateSchedule(t *testing.T) {
	apiKey := checkApiKey(*apiKey)
	cmd := exec.Command(
		"./opsgenie-scheduler-creator",
		"--apiKey", apiKey,
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

// Test Six:
// - Create and delete testTeam and scheduleTest via go run

func TestSixCreateDeleteTestTeamScheduleTest(t *testing.T) {
	apiKey := checkApiKey(*apiKey)
	cmd := exec.Command(
		"go", "run",
		"main.go",
		"--apiKey", apiKey,
		"--teamName", teamName,
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

// Test Seven:
// - Create and delete testTeam and scheduleTest via builded app

func TestSevenCreateDeleteTestTeamScheduleTest(t *testing.T) {
	apiKey := checkApiKey(*apiKey)
	cmd := exec.Command(
		"./opsgenie-scheduler-creator",
		"--apiKey", apiKey,
		"--teamName", teamName,
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

// Test eight:
// - Failed test via go run
func TestFailedCreateScheduleCommand(t *testing.T) {
	cmd := exec.Command(
		"go", "run",
		"--scheduleTeam", scheduleTeam,
		"--scheduleName", scheduleName,
		"--scheduleYear", fmt.Sprint(scheduleYear),
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		t.Fatalf("Command has been failed.\nCommand: %s", err)
	}
}
