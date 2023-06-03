package cmd

import (
	"errors"
	"fmt"
	"github.com/jason0x43/go-toggl"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use: "show",
	// 一ヶ月の稼働時間を表示します。
	Short: "Show the working hours of a month",
	/*
		togglでの一ヶ月の稼働時間を稼働単位で表示します。
		例えば、2022年5月の稼働時間を表示する場合は、
		$ toggl show 2022 5
		と入力します。
		表示時間のタイムゾーンはJSTで表示するformatはDateTimeです。
	*/
	Long: `Show the working hours of a month in toggl.
For example, if you want to show the working hours of May 2022, you can enter:
$ toggl show 2022 5
The time zone of the display time is JST and the format is DateTime.`,
	Run: func(cmd *cobra.Command, args []string) {
		var errs error
		title, err := cmd.Flags().GetString("content")
		if err != nil {
			errs = errors.Join(errs, err)
		} else if title == "" {
			errs = errors.Join(errs, errors.New("content is required"))
		}
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			errs = errors.Join(errs, err)
		} else if token == "" {
			errs = errors.Join(errs, errors.New("token is required"))
		}
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			errs = errors.Join(errs, err)
		}
		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			errs = errors.Join(errs, err)
		}
		projectID, err := cmd.Flags().GetString("project_id")
		if err != nil {
			errs = errors.Join(errs, err)
		} else if projectID == "" {
			errs = errors.Join(errs, errors.New("project_id is required"))
		}
		if errs != nil {
			log.Fatalln(errs)
		}
		session := toggl.OpenSession(token)
		s := &session
		pID, err := strconv.Atoi(projectID)
		if err != nil {
			log.Fatalln(err)
		}
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		entries, err := s.GetTimeEntries(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, jst), time.Date(year, time.Month(month), 31, 23, 59, 59, 0, jst))
		if err != nil {
			log.Fatalln(err)
		}

		for _, entry := range entries {
			if entry.Pid == pID {
				fmt.Printf("%s ~ %s: %s\n", entry.Start.In(jst).Format(time.DateTime), entry.Stop.In(jst).Format(time.DateTime), title)
			}
		}
	},
}

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}
	now := time.Now().In(jst)
	showCmd.Flags().StringP("content", "c", "", "working content")
	showCmd.Flags().StringP("token", "t", "", "toggl api token")
	showCmd.Flags().IntP("year", "y", now.Year(), "target year")
	showCmd.Flags().IntP("month", "m", int(now.Month()), "target month")
	showCmd.Flags().StringP("project_id", "p", "", "project id")
	rootCmd.AddCommand(showCmd)
}
