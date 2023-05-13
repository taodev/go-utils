package ssh

import (
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	User       string
	KeyFile    string
	RemoteAddr string
	TryTime    int

	conn *ssh.Client
}

func (cli *SSHClient) Dial() (err error) {
	for i := 0; i < cli.TryTime; i++ {
		if err = cli._dial(); err == nil {
			break
		}

		<-time.After(time.Second)
	}

	return
}

func (cli *SSHClient) SSHDial(addr string) (conn net.Conn, err error) {
	if cli.conn == nil {
		if err = cli.Dial(); err != nil {
			return
		}
	}

	conn, err = cli.conn.Dial("tcp", addr)

	return
}

func (cli *SSHClient) _dial() (err error) {
	if cli.conn != nil {
		cli.conn.Close()
		cli.conn = nil
	}

	pemFile, err := os.ReadFile(cli.KeyFile)
	if err != nil {
		return
	}

	signer, err := ssh.ParsePrivateKey(pemFile)
	if err != nil {
		log.Printf("ssh: %v", err)
		return
	}

	config := ssh.ClientConfig{
		User:            cli.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		Timeout:         20 * time.Minute,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		BannerCallback: func(message string) error {
			log.Println("[ERROR] ssh:", message)
			return nil
		},
	}

	cli.conn, err = ssh.Dial("tcp", cli.RemoteAddr, &config)
	return
}

func New(addr string, user string, sshKeyFile string) (cli *SSHClient) {
	cli = new(SSHClient)
	cli.RemoteAddr = addr
	cli.User = user
	cli.KeyFile = sshKeyFile
	cli.TryTime = 6

	return
}
