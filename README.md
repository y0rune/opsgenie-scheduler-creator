# Opsgenie-scheduler-creator

The script `Opsgenie-scheduler-creator` is a automatically creator a schedule
with rotation in OpsGenie. The rotation settings are:
- Workweek (from Monday to Friday): daily - 17:00 - 9:00  (5:00PM - 9:00AM)
- Weekend (from Friday to Monday): all weekend - 17:00 - 9:00  (5:00PM - 9:00AM)

## Instalation

```bash
git clone https://github.com/y0rune/opsgenie-scheduler-creator.git
go get
```

## Arguments

```
  -apiKey string
        # ApiKey for use in that script
  -scheduleEnabled
        # Schedule is enabled (default true)
  -scheduleName string
        # Name of schedule (default "Test Schedule")
  -scheduleTeam string
        # Name of the team in the schedule (default "TestTeam")
  -scheduleTimezone string
        # Timezone of the schedule (default "Europe/Warsaw")
  -scheduleYear int
        # Year of the schedule (default 2022)
```

## Example of usage

### How to use it in the console?

```bash
go run main.go --apiKey XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX --scheduleTeam TestTeam --scheduleName "YEAR_2023" --scheduleYear 2023
```

### Output console

```
INFO[2022-10-23T00:36:15.270221+02:00] Client is configured with ApiUrl: api.opsgenie.com, RetryMaxCount: 4
Rotation w1-2.1-9.1 has been created for schedule YEAR_2023.
Rotation w2-9.1-16.1 has been created for schedule YEAR_2023.
Rotation w3-16.1-23.1 has been created for schedule YEAR_2023.
Rotation w4-23.1-30.1 has been created for schedule YEAR_2023.
Rotation w5-30.1-6.2 has been created for schedule YEAR_2023.
Rotation w6-6.2-13.2 has been created for schedule YEAR_2023.
Rotation w7-13.2-20.2 has been created for schedule YEAR_2023.
Rotation w8-20.2-27.2 has been created for schedule YEAR_2023.
Rotation w9-27.2-6.3 has been created for schedule YEAR_2023.
Rotation w10-6.3-13.3 has been created for schedule YEAR_2023.
Rotation w11-13.3-20.3 has been created for schedule YEAR_2023.
Rotation w11-13.3-20.3 has been created for schedule YEAR_2023.
Rotation w12-20.3-27.3 has been created for schedule YEAR_2023.
Rotation w13-27.3-3.4 has been created for schedule YEAR_2023.
Rotation w14-3.4-10.4 has been created for schedule YEAR_2023.
Rotation w15-10.4-17.4 has been created for schedule YEAR_2023.
Rotation w16-17.4-24.4 has been created for schedule YEAR_2023.
Rotation w17-24.4-1.5 has been created for schedule YEAR_2023.
Rotation w18-1.5-8.5 has been created for schedule YEAR_2023.
Rotation w19-8.5-15.5 has been created for schedule YEAR_2023.
Rotation w20-15.5-22.5 has been created for schedule YEAR_2023.
Rotation w21-22.5-29.5 has been created for schedule YEAR_2023.
Rotation w22-29.5-5.6 has been created for schedule YEAR_2023.
Rotation w23-5.6-12.6 has been created for schedule YEAR_2023.
Rotation w24-12.6-19.6 has been created for schedule YEAR_2023.
Rotation w25-19.6-26.6 has been created for schedule YEAR_2023.
Rotation w26-26.6-3.7 has been created for schedule YEAR_2023.
Rotation w27-3.7-10.7 has been created for schedule YEAR_2023.
Rotation w28-10.7-17.7 has been created for schedule YEAR_2023.
Rotation w29-17.7-24.7 has been created for schedule YEAR_2023.
Rotation w30-24.7-31.7 has been created for schedule YEAR_2023.
Rotation w31-31.7-7.8 has been created for schedule YEAR_2023.
Rotation w32-7.8-14.8 has been created for schedule YEAR_2023.
Rotation w33-14.8-21.8 has been created for schedule YEAR_2023.
Rotation w34-21.8-28.8 has been created for schedule YEAR_2023.
Rotation w35-28.8-4.9 has been created for schedule YEAR_2023.
Rotation w36-4.9-11.9 has been created for schedule YEAR_2023.
Rotation w37-11.9-18.9 has been created for schedule YEAR_2023.
Rotation w38-18.9-25.9 has been created for schedule YEAR_2023.
Rotation w39-25.9-2.10 has been created for schedule YEAR_2023.
Rotation w40-2.10-9.10 has been created for schedule YEAR_2023.
Rotation w41-9.10-16.10 has been created for schedule YEAR_2023.
Rotation w42-16.10-23.10 has been created for schedule YEAR_2023.
Rotation w43-23.10-30.10 has been created for schedule YEAR_2023.
Rotation w44-30.10-6.11 has been created for schedule YEAR_2023.
Rotation w45-6.11-13.11 has been created for schedule YEAR_2023.
Rotation w46-13.11-20.11 has been created for schedule YEAR_2023.
Rotation w47-20.11-27.11 has been created for schedule YEAR_2023.
Rotation w48-27.11-4.12 has been created for schedule YEAR_2023.
Rotation w49-4.12-11.12 has been created for schedule YEAR_2023.
Rotation w50-11.12-18.12 has been created for schedule YEAR_2023.
Rotation w51-18.12-25.12 has been created for schedule YEAR_2023.
Rotation w52-25.12-1.1 has been created for schedule YEAR_2023.
```

### Schedule settings in OpsGenie
![alt text](https://github.com/y0rune/opsgenie-scheduler-creator/blob/main/screenshots/OpsGenieSchedule.png)
![alt text](https://github.com/y0rune/opsgenie-scheduler-creator/blob/main/screenshots/OpsGenieUpdateRotation.png)
