package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/sterchelen/hssp/internal/status"
)

const (
	appName = "hssp"
)

// print rfc flag
var print bool

var (
	rootCmd = &cobra.Command{
		Use:   appName,
		Short: "",
		Long:  "",
	}

	codeCmd = &cobra.Command{
		Use:   "code CODE [...]",
		Short: "Displays http code meaning",
		Long: `This command displays the description for the given http code
with its corresponding class and its RFC.`,
		Args: cobra.MinimumNArgs(1),
		RunE: codeRun,
	}

	classCmd = &cobra.Command{
		Use:   "class [CLASS | NAME] [...]",
		Short: "Displays http codes corresponding to a given class",
		Long: `This command displays the list of http status codes corresponding
to the given class, which may be specified as a number (1-5),
a class category string (1xx, 2xx, 3xx, 4xx, 5xx),
or the class name, i.e. informational, successful, redirect, clienterror,
or servererror`,
		Args: cobra.MinimumNArgs(1),
		RunE: classRun,
	}
)

func init() {
	codeCmd.PersistentFlags().Bool("print", false, "Prints respective rfc")
}
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
		class, ok := status.CodeClassFromArg(arg)
		if !ok {
			fmt.Printf("%s: Not a known class or code\n", arg)
			continue
		}
		var tableData [][]string

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

	// This variable will be used if `--print` be false (if it doesn't exist)
	tableData := [][]string{}

	// A dictionary of rfcs (key: rfc link, value rfc text)
	// This variable will be use if `--print` be true (if it does exist)
	var rfcs = make(map[string]string)

	// Check for existence of --print flag
	printRFC, _ := cmd.Flags().GetBool("print")

	for _, arg := range args {
		code, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s: Not a numeric code\n", arg)
			continue
		}

		var statuses status.Statuses
		statuses, err = s.FindStatusesByCode(code)
		if err != nil {
			fmt.Printf("%s: No such code\n", arg)
			continue
		}

		for _, status := range statuses {
			if printRFC {
				rfcTxt, err := getRFCText(status.RFCLink)
				if err != nil {
					fmt.Printf("%s: Error occurred during fetching rfc for print", appName)
					continue
				}
				rfcs[status.RFCLink] = rfcTxt
			} else {
				tableData = append(tableData, []string{strconv.Itoa(status.Code), status.GiveClassName(),
					status.Description, status.RFCLink})
			}
		}

	}

	// Ok now we should print rfc or table
	if printRFC {
		// Print rfcs

		count := 0
		rfcsLen := len(rfcs)
		for _, rfcTxt := range rfcs {
			fmt.Println(rfcTxt)

			count += 1

			// If it's the last rfc we don't need to print `-` character to seperate next current rfc from next rfc (there isn't next one)
			if count != rfcsLen {
				fmt.Println("----------------------------------------------------------------------")
			}
		}
	} else {
		renderTable(tableData)
	}

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

func getRFCText(rfccode string) (string, error) {
	// rfccode is something like "rfc7231"

	url := fmt.Sprintf("https://www.rfc-editor.org/rfc/%s.txt", strings.ToLower(rfccode))
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	rfcBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	rfcTxt := string(rfcBytes)
	return rfcTxt, nil
}
