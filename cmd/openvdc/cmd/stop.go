package cmd

import (
	log "github.com/Sirupsen/logrus"

	"github.com/axsh/openvdc/api"
	"github.com/axsh/openvdc/cmd/openvdc/internal/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var stopCmd = &cobra.Command{
	Use:     "stop [Instance ID]",
	Short:   "Stop a running instance",
	Long:    "Stop a running instance.",
	Example: "openvdc stop i-0000000001",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			log.Fatalf("Please provide an instance ID.")
		}

		instanceID := args[0]

		req := &api.StopRequest{
			InstanceId: instanceID,
		}

		return util.RemoteCall(func(conn *grpc.ClientConn) error {
			c := api.NewInstanceClient(conn)

			_, err := c.Stop(context.Background(), req)
			if err != nil {
				log.WithError(err).Fatal("Disconnected abnormally")
				return err
			}
			return err
		})
	},
}
