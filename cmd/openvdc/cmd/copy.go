package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/axsh/openvdc/api"
	"github.com/axsh/openvdc/cmd/openvdc/internal/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
)

var copyCmd = &cobra.Command{
	Use:   "copy [File src path] [Instance ID]:[file dest path]",
	Short: "Copy files to and between instances",
	Long:  "Copy files to and between instances",
	Example: `
	% openvdc copy 1.txt i-xxxxxxx:/tmp/1.txt
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Fatalf("Please provide a source path.")
		}
		if len(args) == 1 {
			log.Fatalf("Please provide a destination path.")
		}

		src := args[0]
		dest := args[1]

		_, err := os.Stat(src)
		if os.IsNotExist(err) {
			log.Fatalf("Source file not found")
		}

		p := strings.Split(dest, ":")
        	if len(p) < 2 { 
			log.Fatalf("Invalid destination path. Please use this format: [Instance ID]:[file dest path]")
		}

        	instanceID, instanceDir := p[0], p[1]
        	if instanceID == "" {
               		log.Fatalf("Invalid Instance ID")
        	}
        	if instanceDir == "" {
                	log.Fatalf("Invalid destination path")
        	}

		req := &api.CopyRequest{
			InstanceId: instanceID,
		}

		var res *api.CopyReply

		err = util.RemoteCall(func(conn *grpc.ClientConn) error {
			c := api.NewInstanceClient(conn)
			var err error
			res, err = c.Copy(context.Background(), req)
			return err
		})

		if err != nil {
			log.WithError(err).Fatal("Disconnected abnormally")
                }

		host, port, err := net.SplitHostPort(res.GetAddress())
               	if err != nil {
                log.Fatal("Invalid ssh host address: ", res.GetAddress())
                }

		clientConfig := &ssh.ClientConfig{
                	User: instanceID,
        	}
		
		client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), clientConfig)

		if err != nil {
			log.WithError(err).Fatal("ssh.Dial failed")
		}
		
		session, err := client.NewSession()
		if err != nil {
			log.WithError(err).Fatal("Failed to create session")
		}
		defer session.Close()

		file, _ := os.Open(src)
		defer file.Close()

		b, _ := ioutil.ReadAll(file)
		br := bytes.NewReader(b)	

		srcFilename := path.Base(src)
		//srcDir := path.Dir(src)

		go func() {
                	w, _ := session.StdinPipe()
			defer w.Close()
			fmt.Fprintln(w, "C0655", int64(len(b)), srcFilename)
			io.Copy(w, br)
			fmt.Fprintln(w, "\x00")
        	}()

		session.Run("/usr/bin/scp -t ./")

		return nil
	},
}
