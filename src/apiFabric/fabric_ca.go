/*
Copyright SecureKey Technologies Inc. All Rights Reserved.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at


      http://www.apache.org/licenses/LICENSE-2.0


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/rand"
	"strconv"
	_"testing"
	"time"
	
	"github.com/hyperledger/fabric/bccsp"
	_"github.com/hyperledger/fabric-sdk-go/fabric-client/keyvaluestore"
	bccspFactory "github.com/hyperledger/fabric/bccsp/factory"
	
	config "github.com/hyperledger/fabric-sdk-go/config"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	kvs "github.com/hyperledger/fabric-sdk-go/fabric-client/keyvaluestore"
	
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
)

// This test loads/enrols an admin user
// Using the admin, it registers, enrols, and revokes a test user
func main() {
	testSetup := BaseSetupImpl{
		ConfigFile: "config_test.yaml",
	}

	testSetup.InitConfig()
	client := fabricClient.NewClient()

	err := bccspFactory.InitFactories(&bccspFactory.FactoryOpts{
		ProviderName: "SW",
		SwOpts: &bccspFactory.SwOpts{
			HashFamily: config.GetSecurityAlgorithm(),
			SecLevel:   config.GetSecurityLevel(),
			FileKeystore: &bccspFactory.FileKeystoreOpts{
				KeyStorePath: config.GetKeyStorePath(),
			},
			Ephemeral: false,
		},
	})
	if err != nil {
		fmt.Printf("Failed getting ephemeral software-based BCCSP [%s]\n", err)
		
	}
	var cryptoSuite  bccsp.BCCSP
	cryptoSuite = bccspFactory.GetDefault()

	client.SetCryptoSuite(cryptoSuite)
	stateStore, err := kvs.CreateNewFileKeyValueStore("/tmp/enroll_user")
	if err != nil {
		fmt.Printf("CreateNewFileKeyValueStore return error[%s]\n", err)
	}
	client.SetStateStore(stateStore)

	caClient, err := fabricCAClient.NewFabricCAClient()
	if err != nil {
		fmt.Printf("NewFabricCAClient return error: %v\n", err)
	}

	// Admin user is used to register, enrol and revoke a test user
	adminUser, err := client.LoadUserFromStateStore("admin")

	if err != nil {
		fmt.Printf("client.LoadUserFromStateStore return error: %v\n", err)
	}
	if adminUser == nil {
		key, cert, err := caClient.Enroll("admin", "adminpw")
		if err != nil {
			fmt.Printf("Enroll return error: %v\n", err)
		}
		if key == nil {
			fmt.Printf("private key return from Enroll is nil\n")
		}
		if cert == nil {
			fmt.Printf("cert return from Enroll is nil\n")
		}
		
		fmt.Printf("\ntest____1________________\n")
		certPem, _ := pem.Decode(cert)
		if err != nil {
			fmt.Printf("pem Decode return error: %v\n", err)
		}
		
		fmt.Printf("\ntest____________________\n")
		cert509, err := x509.ParseCertificate(certPem.Bytes)
		if err != nil {
			fmt.Printf("x509 ParseCertificate return error: %v\n", err)
		}
		if cert509.Subject.CommonName != "admin" {
			fmt.Printf("CommonName in x509 cert is not the enrollmentID\n")
		}

		keyPem, _ := pem.Decode(key)
		if err != nil {
			fmt.Printf("pem Decode return error: %v\n", err)
		}
		adminUser = fabricClient.NewUser("admin")
		k, err := client.GetCryptoSuite().KeyImport(keyPem.Bytes, &bccsp.ECDSAPrivateKeyImportOpts{Temporary: false})
		if err != nil {
			fmt.Printf("KeyImport return error: %v\n", err)
		}
		adminUser.SetPrivateKey(k)
		adminUser.SetEnrollmentCertificate(cert)
		err = client.SaveUserToStateStore(adminUser, false)
		if err != nil {
			fmt.Printf("client.SaveUserToStateStore return error: %v\n", err)
		}
		adminUser, err = client.LoadUserFromStateStore("admin")
		if err != nil {
			fmt.Printf("client.LoadUserFromStateStore return error: %v\n", err)
		}
		if adminUser == nil {
			fmt.Printf("client.LoadUserFromStateStore return nil\n")
		}
	}

	// Register a random user
	userName := createRandomName()
	registerRequest := fabricCAClient.RegistrationRequest{Name: userName, Type: "user", Affiliation: "org1.department1"}
	enrolmentSecret, err := caClient.Register(adminUser, &registerRequest)
	if err != nil {
		fmt.Printf("Error from Register: %s\n", err)
	}
	fmt.Printf("Registered User: %s, Secret: %s\n", userName, enrolmentSecret)
	// Enrol the previously registered user
	ekey, ecert, err := caClient.Enroll(userName, enrolmentSecret)
	if err != nil {
		fmt.Printf("Error enroling user: %s\n", err.Error())
	}
	//re-enroll
	fmt.Printf("** Attempt to re-enrolled user:  '%s'\n", userName)
	keyPem, _ := pem.Decode(ekey)
	if err != nil {
		fmt.Printf("pem Decode return error: %v\n", err)
	}
	//convert key to bccsp
	k, err := client.GetCryptoSuite().KeyImport(keyPem.Bytes, &bccsp.ECDSAPrivateKeyImportOpts{Temporary: false})
	if err != nil {
		fmt.Printf("KeyImport return error: %v\n", err)
	}
	//create new user object and set certificate and private key of the previously enrolled user
	enrolleduser := fabricClient.NewUser(userName)
	enrolleduser.SetEnrollmentCertificate(ecert)
	enrolleduser.SetPrivateKey(k)
	//reenroll
	_, reenrollCert, err := caClient.Reenroll(enrolleduser)
	if err != nil {
		fmt.Printf("Error Reenroling user: %s\n", err.Error())
	}
	fmt.Printf("** User '%s' was re-enrolled \n", userName)
	if bytes.Equal(ecert, reenrollCert) {
		fmt.Printf("Error Reenroling user. Enrollmet and Reenrollment certificates are the same\n")
	}

	revokeRequest := fabricCAClient.RevocationRequest{Name: userName}
	err = caClient.Revoke(adminUser, &revokeRequest)
	if err != nil {
		fmt.Printf("Error from Revoke: %s\n", err)
	}

}

func createRandomName() string {
	rand.Seed(time.Now().UnixNano())
	return "user" + strconv.Itoa(rand.Intn(500000))
}
