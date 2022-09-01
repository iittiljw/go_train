package lib

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

func executorCmd(cmd *cobra.Command) func(in string) {
	return func(in string) {
		in = strings.TrimSpace(in)
		blocks := strings.Split(in, " ")
		switch blocks[0] {
		case "list":
			if err := listCmd.RunE(cmd, []string{}); err != nil {
				log.Fatalln(err)
			}
		case "test":
			fmt.Println("this is only test!")
		case "exit":
			fmt.Println("Bye!")
			os.Exit(0)
		}
	}
	
}

var suggestions = []prompt.Suggest{
	// Command
	{"test", "测试啦~"},
	{"list", "列出小pod~"},
	{"exit", "Exit prompt"},
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

var promptCmd = &cobra.Command{
	Use:          "prompt",
	Short:        "prompt pods ",
	Example:      "kubectl pods prompt",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		p := prompt.New(
			executorCmd(c),
			completer,
			prompt.OptionPrefix(">>> "),
		)
		p.Run()
		return nil
	},
}
