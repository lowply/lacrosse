# lacrosse

[lacrosse](https://en.wikipedia.org/wiki/La_Crosse,_Wisconsin) - A simple bash script to update DNS records on Amazon Route 53

## Usage

```bash
lacrosse [domain] [type] [record] [TTL] [awscli profile]
```
Example:

```bash
lacrosse test.lowply.com A 192.168.1.0 300 default
```

## Requirements

- [awscli](https://aws.amazon.com/cli/) with the profile configured to have Route 53 permission
- [jq](https://stedolan.github.io/jq/)

## Installation

```bash
$ git clone https://github.com/lowply/lacrosse.git
$ cp lacrosse/lacrosse /usr/local/bin/
```

## Logs

```bash
~/.cache/lacrosse.log
```

## Errors

If you see this error:

```
ValueError: unknown locale: UTF-8
```

Try

```
export LC_ALL=en_US.UTF-8
```

## License

MIT

## Author

Sho Mizutani
