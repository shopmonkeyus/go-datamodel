package cmd

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"path"

// 	v3 "github.com/shopmonkeyus/go-datamodel/v3"
// 	"github.com/spf13/cobra"
// )

// var decodeCmd = &cobra.Command{
// 	Use:   "decode [name] [input-file]",
// 	Short: "Decode a datamodel JSON",
// 	Long:  `This command decode a datamodel JSON document and validate it`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if len(args) == 0 {
// 			cmd.Help()
// 			os.Exit(0)
// 		}
// 		name := args[0]
// 		fn := args[1]
// 		if _, err := os.Stat(fn); err != nil {
// 			fmt.Printf("File %s doesn't exist\n", fn)
// 			os.Exit(1)
// 		}
// 		file, err := os.Open(fn)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
// 		defer file.Close()
// 		buf, err := io.ReadAll(file)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
// 		m, err := v3.NewFromChangeEvent(name, buf, path.Ext(fn) == ".gz")
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
// 		fmt.Println(m)
// 		// j, err := json.MarshalIndent(m, "", " ")
// 		// fmt.Println(string(j), err)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(decodeCmd)
// }
