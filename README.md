# Go JWT Transform

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/go-jwt-transform?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/restuwahyu13/go-jwt-transform)](https://goreportcard.com/report/github.com/restuwahyu13/go-jwt-transform)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/restuwahyu13/go-jwt-transform/blob/main/CONTRIBUTING.md)

**go-jwt-transform** is a simple utility tool for your transforming a real _jwt token_ into a fake _jwt token_, because if you
store _jwt token_ into a `cookie` or `local storage` very unsafe, your _jwt token_ can be seen data using [jwt.io](https://jwt.io)
website or chrome extension, if you use **go-jwt-transform** you real _jwt token_ cannot seen using [jwt.io](https://jwt.io)
website or chrome extension, because what you save is fake _jwt token_, you can get back real _jwt token_ using decrypt method for
parse fake _jwt token_.

## Table Of Content

- [Go JWT Transform](#go-jwt-transform)
  - [Table Of Content](#table-of-content)
  - [Installation](#installation)
  - [API Reference](#api-reference)
  - [Example Usage](#example-usage)
  - [Testing](#testing)
  - [Bugs](#bugs)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

```bash
go get github.com/restuwahyu13/go-jwt-transform
```

## API Reference

- #### Encrypt(token: string, rotate: uint): (string, error)

  encrypt jwt token using caesar cipher cryptography from real jwt token into fake jwt token

- #### Decrypt(token: string, rotate: uint): (string, error)

  decrypt jwt token using caesar cipher cryptography from fake jwt token into real jwt token

## Example Usage

Follow this **[express tutorial](https://github.com/restuwahyu13/express-rest-api-clean-architecture)** for example usage using
express, make this as middleware for transform your fake jwt token to real token, because jwt .verify need real token, if you pass
fake token jwt.verify identification your token is not valid and if you not using express, make this as middleware.

- ##### Example Usage Using CommonJs With JavaScript
```go
  package main

  import (
    transform "github.com/restuwahyu13/go-jwt-transform"
    "fmt"
  )

  func main() {
     const accessToken string ="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    const rotate uint = 15

    res, err := transform.Encrypt(accessToken, rotate)

    if err != nil {
      fmt.Error(err)
    }

    fmt.Println(res)
    // fake jwt token
    // tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r

    res, err := transform.Decrypt(accessToken, rotate)

    if err != nil {
      fmt.Error(err)
    }

    fmt.Println(res)
    // real jwt token
    // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
  }
  ```

## Testing

- Testing Via Local

  ```sh
  go test .
  ```

- Testing Via Docker

  ```sh
  docker build -t go-jwt-transform && docker run go-jwt-transform go test .
  ```

## Bugs

For information on bugs related to package libraries, please visit [here](https://github.com/restuwahyu13/go-jwt-transform/issues)

## Contributing

Want to make **jwt-transform** more perfect ? Let's contribute and follow the
[contribution guide.](https://github.com/restuwahyu13/go-jwt-transform/blob/main/CONTRIBUTING.md)

## License

- [MIT License](https://github.com/restuwahyu13/go-jwt-transform/blob/main/LICENSE.md)

<p align="right" style="padding: 5px; border-radius: 100%; background-color: red; font-size: 2.5rem;">
  <b><a href="#jwt-transform">BACK TO TOP</a></b>
</p>
