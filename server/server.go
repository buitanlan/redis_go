package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/buitanlan/redis_go/config"
)

func readCommand(c net.Conn) (string, error) {
	// TODO: max read in one shot is 512 bytes
	// todo: implement read until command is complete
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}

func RunSyncTCPServer() {
	log.Printf("Starting a sync TCP server on %s:%d", config.Host, config.Port)

	var con_clients int = 0

	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))

	if err != nil {
		panic(err)
	}

	for {
		c, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}

		con_clients++
		log.Printf("New connection from %s", c.RemoteAddr(), "concurrent clients: ", con_clients)

		for {
			// over the socket, continuely the command and read it  out
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients--
				log.Print("Client closed the connection", c.RemoteAddr(), "concurrent clients: ", con_clients)

				if err == io.EOF {
					break
				}

				log.Print("Error reading command: ", err)
			}

			log.Print("Command read: ", cmd)
			if err := respond(cmd, c); err != nil {
				log.Print("Error responding to command: ", err)
			}
		}

	}
}
