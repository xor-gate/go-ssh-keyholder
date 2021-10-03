# go-ssh-keyholder

Go implementation of the MediaWiki KeyHolder for securely sharing ssh agents among groups of users.

# Installation

The go-ssh-keyholder only works under linux due to Unix Socket `SO_PEERCRED` usage for reading the Uid and Gid of the client connection.

`go get github.com/xor-gate/go-ssh-keyholder`

# Usage

```
go-ssh-keyholder -config /path/to/go-keyholder.yml
export SSH_AUTH_SOCK=/path/to/go-keyholder.agent.sock
ssh-add
ssh <host>
```

# Documentation

* https://blog.wikimedia.org/2017/03/22/keyholder/
* https://github.com/wikimedia/keyholder
* http://www.unixwiz.net/techtips/ssh-agent-forwarding.html

# See also

* https://github.com/cptpcrd/unix_cred

# License

[MIT](LICENSE)
