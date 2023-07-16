#!/bin/sh
if [ -z "$SSH_AUTH_SOCK" ]; then
  SSH_AUTH_SOCK=/var/run/go-ssh-keyholder.sock
  readonly SSH_AUTH_SOCK
  export SSH_AUTH_SOCK
fi
