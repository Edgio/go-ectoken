# go-ectoken
> _golang implementation of EdgeCast token (`ectoken`)_


## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)


## Background

golang implementation of the "EdgeCast Token" (`ectoken`) -see main repo [ectoken](https://github.com/VerizonDigital/ectoken) for more details.

## Install


### Build with Make
```sh
go-ectoken>make
go build ./cmd/ectoken/ectoken.go
```

### Running tests
```sh
go-ectoken>make test
go clean -testcache
go test ./...
ok  	_/go-ectoken	0.002s
?   	_/go-ectoken/cmd/ectoken	[no test files]
go test ./cmd/ectoken/...
?   	_/go-ectoken/cmd/ectoken	[no test files]
```

## Usage

### Help
```sh
~>./ectoken 
Usage:
 To Encrypt:
     ec_encrypt <key> <text>
 or:
     ec_encrypt encrypt <key> <text>
 To Decrypt:
     ec_encrypt decrypt <key> <text>
```

### Encrypt

Encrypt clear text token `<token>` with key: `<key>`:
```sh
~>./ectoken encrypt MY_SECRET_KEY MY_COOL_TOKEN
aaOUHrqxEmCLpnSAwTD06I54ffrtaa0K-3P7KKbosRChSubvbtWq5VY
```

### Decrypt

Decrypt ciphertext token `<token>` with key: `<key>`:
```sh
~>./ectoken decrypt MY_SECRET_KEY lm6b156_M3XkS8SWOdJ_D1UBPZZzX7cDxh6aCb9kJ7pox0eco9XZqSk
MY_COOL_TOKEN
```


## Contribute

- We welcome issues, questions and pull requests.


## License

This project is licensed under the terms of the Apache 2.0 open source license. Please refer to the `LICENSE-2.0.txt` file for the full terms.
