# go-ssh-keyholder

Go implementation of the MediaWiki KeyHolder for securely sharing ssh a secured and system wide SSH agent among groups of users.

# Installation

The go-ssh-keyholder only works under UNIX-like OS'es. Currently supported are Linux `SO_PEERCRED` and FreeBSD `LOCAL_PEERCRED` for reading the Uid and Gid of the unix socket ssh agent client connection.

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
