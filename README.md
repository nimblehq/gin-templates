<p align="center">
  <img alt="Nimble logo" src="https://assets.nimblehq.co/logo/light/logo-light-text-320.png" />
</p>

<p align="center">
  <strong>Gin Templates</strong>
</p>

---

Our templates offer a rich boilerplate to jump start Go Gin-based application development.

## Get Started

### Prerequisite 
- [Python interpreter](https://docs.python.org/3/using/index.html)
- Cookiecutter
  - Mac: `brew install cookiecutter`
  - Ubuntu: `brew install cookiecutter`
- Go version >= 1.15

## Usage
- Download **latest** version of gin-templates.
  ```
  go get github.com/nimblehq/gin-templates
  ```
- Build the binary file.
  ```
  go build -o $GOPATH/bin/nimble-gin github.com/nimblehq/gin-templates
  ```
- Create the project from a template on a specifying branch.
  ```
  nimble-gin -b {{git branch}} create
  ```
  **Note:** `-b` can be omitted, branch will be `main` by default.
- Follow the instructions in the terminal.
- Your new application is created 🎉 .

## Test

Execute all unit tests:

```sh
BRANCH={{git branch}} make test
```

Test command will use the template from specifying `BRANCH` to create the temporary testing project (default: `main`).

## License

This project is Copyright (c) 2014 and onwards Nimble. It is free software,
and may be redistributed under the terms specified in the [LICENSE] file.

[LICENSE]: /LICENSE

## About

![Nimble](https://assets.nimblehq.co/logo/dark/logo-dark-text-160.png)

This project is maintained and funded by Nimble.

We love open source and do our part in sharing our work with the community!
See [our other projects][community] or [hire our team][hire] to help build your product.

[community]: https://github.com/nimblehq
[hire]: https://nimblehq.co/
