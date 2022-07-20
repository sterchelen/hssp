package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/sterchelen/hssp/internal/status"
)

const (
	appName = "hssp"
)

var (
	rootCmd = &cobra.Command{
		Use:   appName,
		Short: "",
		Long:  "",
	}

	codeCmd = &cobra.Command{
		Use:   "code",
		Short: "Displays http code meaning",
		Long: `This command displays the given http code description 
with its corresponding class and its RFC.`,
		Args: cobra.MinimumNArgs(1),
		RunE: codeRun,
	}

	classCmd = &cobra.Command{
		Use:   "class",
		Short: "Displays http codes corresponding to a given class",
		Long: `This command displays the list of http status codes corresponding
to the given class number (1,2,3,4,5).`,
		Args: cobra.MinimumNArgs(1),
		RunE: classRun,
	}
)

func Execute() error {
	rootCmd.AddCommand(codeCmd)
	rootCmd.AddCommand(classCmd)

	return rootCmd.Execute()
}

func classRun(cmd *cobra.Command, args []string) error {
	s, err := status.Initialize()
	if err != nil {
		return fmt.Errorf("class: Unable to initialize status due to: %s", err)
	}

	for _, arg := range args {
		var tableData [][]string
		class, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s: Not a numeric code\n", arg)
			continue
		}

		statuses, err := s.StatusesByClass(class)
		if err != nil {
			fmt.Printf("%s: No such class\n", arg)
			continue
		}

		for _, status := range statuses {
			tableData = append(tableData, []string{strconv.Itoa(status.Code), status.GiveClassName(),
				status.Description, status.RFCLink})
		}
		renderTable(tableData)
	}

	return nil
}

func codeRun(cmd *cobra.Command, args []string) error {
	s, err := status.Initialize()
	if err != nil {
		return fmt.Errorf("code: Unable to initialize status due to: %s", err)
	}
	tableData := [][]string{}

	for _, arg := range args {
		code, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s: Not a numeric code\n", arg)
			continue
		}

		sCode, err := s.FindStatusByCode(code)
		if err != nil {
			fmt.Printf("%s: No such code\n", arg)
			continue
		}

		tableData = append(tableData,
			[]string{strconv.Itoa(sCode.Code), sCode.GiveClassName(), sCode.Description, sCode.RFCLink},
		)
	}
	renderTable(tableData)

	return nil
}

func renderTable(tableData [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Code", "Class", "Description", "RFC"})

	if len(tableData) > 0 {
		for _, v := range tableData {
			table.Append(v)
		}
		table.Render()
	}
}
