# xtime-conv

A small CLI tool that converts Unix timestamps to human-readable time. Automatically detects the timestamp unit: seconds, milliseconds, microseconds, or nanoseconds.

## Installation

### From source

```bash
go install github.com/MichurinDev/xtime-conv/cmd/xtime-conv@latest
```

### From releases

Download a binary for your platform from [GitHub Releases](https://github.com/MichurinDev/xtime-conv/releases).

### Local build

```bash
git clone https://github.com/MichurinDev/xtime-conv.git
cd xtime-conv
make install
```

## Usage

```bash
xtime-conv -t <timestamp> [options]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-t` | Unix timestamp (required) | — |
| `-tz` | Timezone offset in hours | `0` (UTC) |
| `-f` | Output time format (Go layout) | `02.01.2006 15:04:05` |
| `-v`, `--version` | Print version and exit | — |

### Examples

```bash
# Seconds
xtime-conv -t 1710000000
# UTC:  09.03.2024 10:40:00

# Milliseconds (auto-detected)
xtime-conv -t 1710000000000
# UTC:  09.03.2024 10:40:00

# With timezone offset UTC+3
xtime-conv -t 1710000000 -tz 3
# UTC:    09.03.2024 10:40:00
# Local:  09.03.2024 13:40:00

# Custom format (Go reference time: Mon Jan 2 15:04:05 MST 2006)
xtime-conv -t 1710000000 -f "2006-01-02 15:04:05"
# UTC:  2024-03-09 10:40:00

# Print version
xtime-conv -v
```

## Development

```bash
make help    # list available commands
make build   # build to bin/xtime-conv
make test    # run tests
make run ARGS="-t 1710000000"   # run without installing
```

## Contributing

This project uses [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add new feature
fix: bug fix
docs: documentation changes
chore: maintenance tasks
```

## License

MIT
