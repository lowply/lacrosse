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
lacrosse [domain] [type] [record] [TTL] [aws profile]
```

Example:

```bash
lacrosse example.com A 93.184.216.34 300 default
```

## Requirements

- macOS or Linux based platform
- aws profile granted Route 53 read/write permission

## Installation

```
go get github.com/lowply/lacrosse
```

## Logs

Logs will be recorded in `~/.cache/lacrosse.log` file in JSON format.

## Development

```bash
$ git clone https://github.com/lowply/lacrosse.git
$ cd lacrosse
$ make
```

## License

MIT

## Author

Sho Mizutani
