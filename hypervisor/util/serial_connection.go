package util

import (
	"net"
	"time"
	"sync"
	"bytes"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"github.com/axsh/openvdc/hypervisor"
)

type SerialConnection struct {
	SerialConn net.Conn
}

func (sc *SerialConnection) AttachSerialConsole(param *hypervisor.ConsoleParam, waitClosed *sync.WaitGroup, errc chan error) error {
	waitClosed.Add(1)
	go func() {
		b := make([]byte, 8192) // 8 kB is the default page size for most modern file systems
		for {
			select {
			case err := <-errc:
				 errc <- err
				break
			default:
				n, err := param.Stdin.Read(b)
				if err != nil {
					errc <- errors.Wrap(err, "\nFailed to read from the from the console input buffer\n\n")
				}

				if bytes.Contains(b[0:n], []byte{0x11}) {
					log.Info("Received exit from stdin")
					errc <- errors.Wrap(err, "\nConsole exited by ctrl-q\n\n")
				}
				_, err = sc.SerialConn.Write(b[0:n])
				if err != nil {
					errc <- errors.Wrap(err, "\nFailed to write to the connection from the buffer\n\n")
				}
			}
		}
		defer waitClosed.Done()
	}()

	waitClosed.Add(1)
	go func() {
		b := make([]byte, 8192)
		for {
			sc.SerialConn.SetDeadline(time.Now().Add(time.Second))
			n, err := sc.SerialConn.Read(b)
			select {
			case err := <-errc:
				if _, e := param.Stdout.Write([]byte{0x0A}); e != nil {
					errc <- errors.Wrap(e, "\nFailed to write the linefeed character on exit\n\n")
				} else {
					errc <- err
				}
				break
			default:
				if err != nil && !err.(net.Error).Timeout() {
					errc <- errors.Wrap(err, "\nFailed to read the connection buffer from socket ")
				}
				// exit on ctrl + q
				if bytes.Contains(b[0:n], []byte{0x11}) {
					errc <- nil
				}
				_, err = param.Stdout.Write(b[0:n])
				if err != nil {
					errc <- errors.Wrap(err, "\nFailed to write to the console stdout buffer\n\n")
				}
			}
		}
		defer waitClosed.Done()
	}()

	return nil
}
