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
	var tableData [][]string

	class, err := strconv.Atoi(args[0])
	if err != nil {
		ok := false
		if class, ok = status.CodeClassFromName(args[0]); !ok {
			return fmt.Errorf("class: please, give a numerical value; %s", err)
		}
	}

	statuses, err := s.StatusesByClass(class)
	if err != nil {
		return err
	}

	for _, status := range statuses {
		tableData = append(tableData, []string{strconv.Itoa(status.Code), status.GiveClassName(),
			status.Description, status.RFCLink})
	}
	renderTable(tableData)
	return nil
}

func codeRun(cmd *cobra.Command, args []string) error {
	s, err := status.Initialize()
	var tableData [][]string

	code, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("code: please, give a numerical value; %s", err)
	}

	sCode, err := s.FindStatusByCode(code)
	if err != nil {
		return err
	}

	tableData = [][]string{
		[]string{strconv.Itoa(sCode.Code), sCode.GiveClassName(), sCode.Description, sCode.RFCLink},
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
