package common

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"strings"

	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	"google.golang.org/grpc/credentials"
)

func LoadTlsCert(path string, serverName string) credentials.TransportCredentials {
	if path == "" {
		return nil
	}

	// build cert obj
	rootCert, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err, "cannot load tls cert from path")
		return nil
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(rootCert) {
		fmt.Println(err, "failed to add server CA's certificate")
		return nil
	}
	// get domain from tcp://domain:port
	domain := strings.Split(serverName, ":")[1][2:]
	config := &tls.Config{
		RootCAs:    certPool,
		ServerName: domain,
	}
	return credentials.NewTLS(config)
}

func MsgResponse(data []byte) []*chaintypes.TxResponseGenericMessage {
	response := chaintypes.TxResponseData{}
	err := response.Unmarshal(data)
	if err != nil {
		panic(err)
	}
	return response.Messages
}
