package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

var lintCmd = &cobra.Command{
	Use:   "lint [directory w/ trailing slash]",
	Short: "Validate code fences",
	Long: `Validates each code fence in each Markdown file in the directory.
	
	Currently only supports YAML code fences.`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		directory := args[0]
		foundError := false

		files, err := ioutil.ReadDir(directory)

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {

			if !strings.HasSuffix(file.Name(), ".md") {
				continue
			}

			if checkFile(directory, file.Name()) {
				foundError = true
			}
		}

		if foundError {
			return errors.New("Failed code block validation.")
		} else {
			return nil
		}
	},
}

func checkFile(directory, doc string) bool {

	file, err := os.Open(directory + doc)
	defer file.Close()

	if err != nil {
		log.Fatalf("Could not open file: %s", doc)
	}

	reader := bufio.NewReader(file)

	var line string
	var lineNum int
	var lineStart int
	var block []string

	foundError := false
	scan := false

	for {
		line, err = reader.ReadString('\n')

		if err != nil {
			break
		}

		lineNum += 1

		if strings.HasPrefix(line, "```yaml") {
			scan = true
			lineStart = lineNum
		} else if scan && strings.HasPrefix(line, "```") {
			scan = false
			t := T{}
			snippet := strings.Join(block, "\n")

			err := yaml.Unmarshal([]byte(snippet), &t)

			if err != nil {

				foundError = true
				errString := err.Error()[strings.LastIndex(err.Error(), ":")+2:]
				fmt.Printf("%s:%d => %s\n", doc, lineStart, errString)
			}
		} else if scan && strings.Compare(line, "#...\n") == 0 && lineNum == (lineStart+1) {
			block = append(block, "fake-key:\n")
		} else if scan {
			block = append(block, line)
		}
	}

	return foundError
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
