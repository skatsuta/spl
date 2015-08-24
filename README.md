# spl

 [![wercker status](https://app.wercker.com/status/d799f872500c49dd0f91696215d1b30f/s/master "wercker status")](https://app.wercker.com/project/bykey/d799f872500c49dd0f91696215d1b30f)

Command-line splitter for strings. Mainly useful to show each element of `$PATH` separately.

```bash
$ spl $PATH
 1  /Users/skatsuta/.pyenv/shims
 2  /Users/skatsuta/.rbenv/shims
 3  /Users/skatsuta/bin
 4  /Users/skatsuta/.cabal/bin
 5  /usr/local/bin
 6  /usr/local/sbin
 7  /usr/bin
 8  /bin
 9  /usr/sbin
10  /sbin
11  /Users/skatsuta/.fzf/bin
```

## Description

`spl` is a command-line splitter for strings that are combined with a delimiter.

I made this tool to mainly show each element of `$PATH` separately.
Of course I could do it using shell script, Ruby, Python, etc, but I just want to write Go! ;)

## Install

To install, use `go get`:

```bash
$ go get github.com/skatsuta/spl
```

## Usage

You just pass strings or environment variables to `spl`. `spl` shows the lists of elements of each argument.

For example:

```bash
$ echo $PATH $GOPATH
/Users/skatsuta/.pyenv/shims:/Users/skatsuta/.rbenv/shims:/Users/skatsuta/bin:/Users/skatsuta/.cabal/bin:/usr/local/bin:/usr/local/sbin:/usr/bin:/bin:/usr/sbin:/sbin:/Users/skatsuta/.fzf/bin /Users/skatsuta

$ spl $PATH $GOPATH
 1  /Users/skatsuta/.pyenv/shims
 2  /Users/skatsuta/.rbenv/shims
 3  /Users/skatsuta/bin
 4  /Users/skatsuta/.cabal/bin
 5  /usr/local/bin
 6  /usr/local/sbin
 7  /usr/bin
 8  /bin
 9  /usr/sbin
10  /sbin
11  /Users/skatsuta/.fzf/bin

1  /Users/skatsuta
```

### Flags

#### `-nonum`: hide line numbers

By default `spl` shows line numbers. To hide them, you can pass `-nonum` flag.

```bash
$ spl -nonum $PATH $GOPATH
/Users/skatsuta/.pyenv/shims
/Users/skatsuta/.rbenv/shims
/Users/skatsuta/bin
/Users/skatsuta/.cabal/bin
/usr/local/bin
/usr/local/sbin
/usr/bin
/bin
/usr/sbin
/sbin
/Users/skatsuta/.fzf/bin

/Users/skatsuta
```

#### `-delim`: specify a delimiter

By default `spl` treats `:` as a delimiter. To change it, you can pass `-delim <DELIM>` flag. 

```bash
$ spl -delim , foo,bar,baz
1  foo
2  bar
3  baz
```


## Contribution

1. Fork ([https://github.com/skatsuta/spl/fork](https://github.com/skatsuta/spl/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Soshi Katsuta](https://github.com/skatsuta)
