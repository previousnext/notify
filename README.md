Notify
======

Slack cli notification sender.

## Usage

```bash
usage: notify [<flags>] <message> <url>

Flags:
  --help             Show help.
  --name="Notify"    Name your bot.
  --emoji=":slack:"  Give your bot a custom image.
  --channel="#general"
                     Which channel do you wish to post in.

Args:
  <message>  The message you wish to send.
  <url>      The url you wish to post to.
```

## Development

```bash
make
```

### Documentation

See `/docs`

### Resources

* [Dave Cheney - Reproducible Builds](https://www.youtube.com/watch?v=c3dW80eO88I)
* [Bryan Cantril - Debugging under fire](https://www.youtube.com/watch?v=30jNsCVLpAE&t=2675s)
* [Sam Boyer - The New Era of Go Package Management](https://www.youtube.com/watch?v=5LtMb090AZI)
* [Kelsey Hightower - From development to production](https://www.youtube.com/watch?v=XL9CQobFB8I&t=787s)

### Tools

```bash
# Dependency management
go get -u github.com/golang/dep/cmd/dep

# Testing
go get -u github.com/golang/lint/golint

# Release management.
go get -u github.com/tcnksm/ghr

# Build
go get -u github.com/mitchellh/gox
```

### Workflow

**Testing**

```bash
make lint test
```

**Building**

```bash
make build
```

**Releasing**

```bash
make release
```