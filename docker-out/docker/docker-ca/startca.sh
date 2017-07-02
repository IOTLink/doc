#!/bin/bash
export FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
export FABRIC_CA_SERVER_CA_NAME=ca-org1
export FABRIC_CA_SERVER_TLS_ENABLED=false
export FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.org1.example.com-cert.pem
export FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/15209b469925589e07f5005d768169e2652a8e4227cd670176cc5ac10afa3e70_sk
sh -c 'fabric-ca-server start --ca.certfile /etc/hyperledger/fabric-ca-server-config/ca.org1.example.com-cert.pem --ca.keyfile /etc/hyperledger/fabric-ca-server-config/15209b469925589e07f5005d768169e2652a8e4227cd670176cc5ac10afa3e70_sk -b admin:adminpw -d'
