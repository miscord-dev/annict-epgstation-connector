# annict-epgstation-connector

A service that automatically creates recording rules in EPGStation based on your anime watch list from Annict.

## Overview

This connector synchronizes your anime watch status from [Annict](https://annict.com/) (a Japanese anime tracking service) to [EPGStation](https://github.com/l3tnun/EPGStation) (a Japanese TV recording system). It automatically creates recording rules for anime you want to watch, are currently watching, or have on hold, while optionally excluding shows available on streaming services.

### Key Features

- **Automatic Recording Rules**: Creates EPGStation recording rules based on your Annict watch status
- **VOD Service Filtering**: Skip recording for anime available on streaming services like Netflix, Amazon Prime, Hulu, etc.
- **Daemon Mode**: Continuous synchronization with configurable intervals
- **Rule Removal**: Automatically removes recording rules when you mark anime as completed or dropped
- **Prometheus Metrics**: Built-in monitoring and metrics collection
- **Deduplication**: Prevents duplicate recording rules using local database storage

## Prerequisites

- Go 1.24 or later
- Running EPGStation instance
- Annict account with API access

## Installation

### From Source

```bash
git clone https://github.com/miscord-dev/annict-epgstation-connector
cd annict-epgstation-connector
go build ./cmd/annict-epgstation-connector
```

### Using Go Install

```bash
go install github.com/miscord-dev/annict-epgstation-connector/cmd/annict-epgstation-connector@latest
```

## Setup

### 1. Get Annict API Token

1. Visit [Annict Developer Console](https://annict.com/settings/apps)
2. Create a new application or use an existing one
3. Copy your API token

### 2. Configure EPGStation

Ensure your EPGStation instance is running and accessible. The default API endpoint is typically `http://localhost:8888/api`.

### 3. Set Environment Variables

```bash
export ANNICT_API_TOKEN="your-annict-api-token"
export EPGSTATION_ENDPOINT="http://localhost:8888/api"
```

## Usage

### Basic Usage (One-time Sync)

```bash
./annict-epgstation-connector sync
```

### Daemon Mode (Continuous Sync)

```bash
# Run with default 60-second interval
./annict-epgstation-connector sync --daemon

# Custom interval (5 minutes)
./annict-epgstation-connector sync --daemon --interval 300

# With custom database path and metrics endpoint
./annict-epgstation-connector sync --daemon \
  --interval 300 \
  --db-path /path/to/database \
  --metrics-listen-address :9090
```

### VOD Service Filtering

Exclude recording rules for anime available on streaming services:

```bash
# Exclude Netflix and Amazon Prime
./annict-epgstation-connector sync --exclude-vod-services netflix,amazon-prime

# With fallback detection (searches all page links when specific VOD section not found)
./annict-epgstation-connector sync \
  --exclude-vod-services netflix,hulu \
  --enable-vod-fallback
```

**Supported VOD Services:**

- `netflix` - Netflix
- `amazon-prime` - Amazon Prime Video
- `hulu` - Hulu
- `disney` - Disney+
- `abema` - AbemaTV
- `crunchyroll` - Crunchyroll
- `funimation` - Funimation
- `dazn` - DAZN
- `bandai` - Bandai Channel
- `nico` - Niconico
- `danime` - dAnime Store

### Automatic Rule Removal

Enable automatic removal of recording rules when anime is marked as completed or dropped:

```bash
./annict-epgstation-connector sync --enable-rule-removal
```

### Complete Example

```bash
./annict-epgstation-connector sync \
  --daemon \
  --interval 300 \
  --exclude-vod-services netflix,amazon-prime,hulu \
  --enable-vod-fallback \
  --enable-rule-removal \
  --metrics-listen-address :9090 \
  --db-path /opt/annict-connector/db
```

## Configuration

### Command Line Options

| Flag | Description | Default |
|------|-------------|---------|
| `--daemon, -d` | Run as a daemon | `false` |
| `--interval, -i` | Sync interval in seconds (daemon mode) | `60` |
| `--metrics-listen-address` | Metrics HTTP server address | `:8080` |
| `--db-path, --db` | Database storage path | `/tmp/annict-epgstation-connector` |
| `--exclude-vod-services` | Comma-separated list of VOD services to exclude | `none` |
| `--enable-vod-fallback` | Enable fallback VOD detection | `false` |
| `--enable-rule-removal` | Enable automatic rule removal | `false` |

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `ANNICT_API_TOKEN` | Your Annict API token | Yes |
| `EPGSTATION_ENDPOINT` | EPGStation API endpoint | Yes |

## Monitoring

When running in daemon mode, Prometheus metrics are exposed at the configured metrics endpoint (default: `:8080/metrics`).

### Available Metrics

The service exposes various metrics for monitoring synchronization status, recording rule creation, and error rates.

## Development

### Building

```bash
go build ./cmd/annict-epgstation-connector
```

### Code Generation

```bash
make generate  # Runs go generate ./...
```

### Testing

```bash
# Run E2E tests
ginkgo tests/e2e/

# Or using Go test runner
go test ./tests/e2e/
```

### Linting

The project uses golangci-lint for code quality checks:

```bash
golangci-lint run
```

## Architecture

- **CLI**: Built with urfave/cli/v2
- **Annict Client**: GraphQL client using genqlient
- **EPGStation Client**: REST client using oapi-codegen
- **Database**: CockroachDB Pebble for local key-value storage
- **VOD Checker**: Web scraper for streaming service availability
- **Metrics**: Prometheus metrics collection

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run tests and linting
6. Submit a pull request

## Support

For issues and questions, please use the GitHub issue tracker.
