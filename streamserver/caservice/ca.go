package caservice


import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	config "github.com/hyperledger/fabric-sdk-go/config"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	kvs "github.com/hyperledger/fabric-sdk-go/fabric-client/keyvaluestore"
	bccspFactory "github.com/hyperledger/fabric/bccsp/factory"

	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
)

type BaseSetupImpl struct {
	Client             fabricClient.Client
	OrdererAdminClient fabricClient.Client
	Chain              fabricClient.Chain
	EventHub           events.EventHub
	ConnectEventHub    bool
	ConfigFile         string
}

type ca struct{
	CaService services
}

func (c *ca)InitCaServer() {
	testSetup := BaseSetupImpl{
		ConfigFile: "../fixtures/config/config_test.yaml",
	}
}

func (c *ca)RegisterUser() {

}

func (c *ca)LoadUser() {

}