package cmd

import (
	"fmt"
	"io"
	"os"

	v3 "github.com/shopmonkeyus/go-datamodel/v3"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode [name] [input-file]",
	Short: "Decode a datamodel JSON",
	Long:  `This command decode a datamodel JSON document and validate it`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		name := args[0]
		fn := args[1]
		if _, err := os.Stat(fn); err != nil {
			fmt.Printf("File %s doesn't exist\n", fn)
			os.Exit(1)
		}
		jsonFile, err := os.Open(fn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		buf, err := io.ReadAll(jsonFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer jsonFile.Close()
		var enctype v3.EncodingType
		// sniff the payload and see if its json
		if string(buf[0:1]) == "{" && string(buf[len(buf)-2:len(buf)-1]) == "}" {
			enctype = v3.JSONEncoding
		} else {
			enctype = v3.MsgPackEncoding
		}
		m, err := v3.NewFromModel(name, buf, enctype)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(m.String())
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}
