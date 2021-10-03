package main

import (
	"fmt"
	"log"

	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/dsa"
	"crypto/ed25519"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

var ErrKeyringOperationDenied = fmt.Errorf("keyring operation denied")

type TrustedKeyring struct {
	agent.Agent
}

func (kr *TrustedKeyring) Add(key agent.AddedKey) error {
	var err error
	var sshPublicKey ssh.PublicKey
	var keyType string

	switch pkey := key.PrivateKey.(type) {
	case *ecdsa.PrivateKey:
		p := pkey.Public()
		sshPublicKey, err = ssh.NewPublicKey(p)
		if err != nil {
			return err
		}
		keyType = "ECDSA"
	case *rsa.PrivateKey:
		p := pkey.Public()
		sshPublicKey, err = ssh.NewPublicKey(p)
		if err != nil {
			return err
		}
		keyType = "RSA"
	case *dsa.PrivateKey:
		sshPublicKey, err = ssh.NewPublicKey(pkey.PublicKey)
		if err != nil {
			return err
		}
		keyType = "DSA"
	case *ed25519.PrivateKey:
		p := pkey.Public()
		sshPublicKey, err = ssh.NewPublicKey(p)
		if err != nil {
			return err
		}
		keyType = "ED25519"
	}

	log.Printf("Identity added: %s %s (%s)\n", ssh.FingerprintSHA256(sshPublicKey), key.Comment, keyType)

	return kr.Agent.Add(key)
}

type SignOnlyKeyring struct {
	agent.Agent
}

func (kr *SignOnlyKeyring) Add(key agent.AddedKey) error {
	return ErrKeyringOperationDenied
}

func (kr *SignOnlyKeyring) Remove(key ssh.PublicKey) error {
	return ErrKeyringOperationDenied
}

func (kr *SignOnlyKeyring) RemoveAll() error {
	return ErrKeyringOperationDenied
}

func (kr *SignOnlyKeyring) Lock(passphrase []byte) error {
	return ErrKeyringOperationDenied
}

func (kr *SignOnlyKeyring) Unlock(passphrase []byte) error {
	return ErrKeyringOperationDenied
}

type NoopKeyring struct{}

func (kr *NoopKeyring) List() ([]*agent.Key, error) {
	return nil, ErrKeyringOperationDenied
}

func (kr *NoopKeyring) Sign(key ssh.PublicKey, data []byte) (*ssh.Signature, error) {
	return nil, ErrKeyringOperationDenied
}

func (kr *NoopKeyring) Add(key agent.AddedKey) error {
	return ErrKeyringOperationDenied
}

func (kr *NoopKeyring) Remove(key ssh.PublicKey) error {
	return ErrKeyringOperationDenied
}

func (kr *NoopKeyring) RemoveAll() error {
	return ErrKeyringOperationDenied
}

func (kr *NoopKeyring) Lock(passphrase []byte) error {
	return ErrKeyringOperationDenied
}

func (kr *NoopKeyring) Unlock(passphrase []byte) error {
	return ErrKeyringOperationDenied
}

func (kr *NoopKeyring) Signers() ([]ssh.Signer, error) {
	return nil, ErrKeyringOperationDenied
}
