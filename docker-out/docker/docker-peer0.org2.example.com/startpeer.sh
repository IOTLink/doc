#!/bin/bash
export CORE_PEER_ID=peer0.org2.example.com
export CORE_PEER_ADDRESS=peer0.org2.example.com:7051
export CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.org2.example.com:7051
export CORE_PEER_GOSSIP_BOOTSTRAP=peer0.org2.example.com:7051
export CORE_PEER_LOCALMSPID=Org2MSP
export CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
export CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=linuxamd64_default
export CORE_LOGGING_LEVEL=DEBUG
export CORE_PEER_TLS_ENABLED=false
export CORE_PEER_ENDORSER_ENABLED=true
export CORE_PEER_GOSSIP_USELEADERELECTION=true
export CORE_PEER_GOSSIP_ORGLEADER=false
export CORE_PEER_PROFILE_ENABLED=true
export CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
export CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
export CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt
export FABRIC_CFG_PATH=/etc/hyperledger/fabric
sh -c 'peer node start'
