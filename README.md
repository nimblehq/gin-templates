<p align="center">
  <img alt="Nimble logo" src="https://assets.nimblehq.co/logo/light/logo-light-text-320.png" />
</p>

<p align="center">
  <strong>Gin Templates</strong>
</p>

---

Our templates offer a rich boilerplate to jump start Go [Gin-based](https://github.com/gin-gonic/gin) application development.

Check this [wiki](https://github.com/nimblehq/gin-templates/wiki/Directories) for more information about the template structure.

## Get Started

### Prerequisite

- [Python interpreter](https://docs.python.org/3/using/index.html)
- Cookiecutter
  - Mac: `brew install cookiecutter`
  - Debian/Ubuntu: `sudo apt-get install cookiecutter`
- Go version >= 1.20

## Usage

- Download **latest** version of gin-templates.

  ```
  go get github.com/nimblehq/gin-templates
  ```

- Build the binary file.

  ```
  go build -o $GOPATH/bin/nimble-gin github.com/nimblehq/gin-templates
  ```

- Create the project using gin-templates. Note that the **main** branch is being used by default. Refer to [this wiki page](https://github.com/nimblehq/gin-templates/wiki/Commands) for instructions on how to use a different branch.

  ```
  nimble-gin create
  ```

- Follow the instructions in the terminal.

- Your new application is created 🎉 .

## Test

```sh
make test
```

## License

This project is Copyright (c) 2014 and onwards Nimble. It is free software,
and may be redistributed under the terms specified in the [LICENSE] file.

[license]: /LICENSE

## About

![Nimble](https://assets.nimblehq.co/logo/dark/logo-dark-text-160.png)

This project is maintained and funded by Nimble.

We love open source and do our part in sharing our work with the community!
See [our other projects][community] or [hire our team][hire] to help build your product.

[community]: https://github.com/nimblehq
[hire]: https://nimblehq.co/
