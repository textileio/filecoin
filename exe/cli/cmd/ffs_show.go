package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/caarlos0/spin"
	"github.com/ipfs/go-cid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	ffsShowCmd.Flags().StringP("cid", "c", "", "cid of the data to pin")
	ffsShowCmd.Flags().StringP("token", "t", "", "FFS auth token")

	ffsCmd.AddCommand(ffsShowCmd)
}

var ffsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show pinned cid data",
	Long:  `Show pinned cid data`,
	PreRun: func(cmd *cobra.Command, args []string) {
		err := viper.BindPFlags(cmd.Flags())
		checkErr(err)
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()

		cidString := viper.GetString("cid")

		if cidString == "" {
			Fatal(errors.New("store command needs a cid"))
		}

		c, err := cid.Parse(cidString)
		checkErr(err)

		s := spin.New("%s Getting info from cid...")
		s.Start()
		info, err := fcClient.Ffs.Show(authCtx(ctx), c)
		s.Stop()
		checkErr(err)

		buf, err := json.MarshalIndent(info, "", "  ")
		checkErr(err)
		Message("%s", buf)

	},
}
