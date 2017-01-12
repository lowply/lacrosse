# lacrosse

[![GitHub release](http://img.shields.io/github/release/lowply/lacrosse.svg?style=flat-square)][release]
[![Travis](https://img.shields.io/travis/lowply/lacrosse.svg?style=flat-square)][travis]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/lowply/lacrosse/releases
[travis]: https://travis-ci.org/lowply/lacrosse
[license]: https://github.com/lowply/lacrosse/blob/master/LICENSE

[La Crosse](https://en.wikipedia.org/wiki/U.S._Route_53#Major_intersections) - A simple CLI tool to update DNS records on Amazon Route 53.

## Usage

```bash
lacrosse [domain] [type] [record] [TTL] [awscli profile]
```
Example:

```bash
lacrosse test.lowply.com A 192.168.1.0 300 default
```

## Requirements

- macOS or Linux based platform
- [awscli](https://aws.amazon.com/cli/) with the profile configured to have Route 53 permission

## Installation

- Download the latest version from the [release page](https://github.com/lowply/lacrosse/releases), unarchive it and put the binary into `/usr/local/bin`.
- Or build by yourself:

```bash
$ git clone https://github.com/lowply/lacrosse.git
$ cd lacrosse
$ make build
$ cp bin/lacrosse /usr/local/bin/
```

## Logs

Logs will be recorded in JSON format.

```bash
~/.cache/lacrosse.log
```

## Development

```bash
$ git clone https://github.com/lowply/lacrosse.git
$ cd lacrosse
$ make deps
```

## License

MIT

## Author

Sho Mizutani
