package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

var ErrKeyringOperationDenied = fmt.Errorf("keyring operation denied")

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
