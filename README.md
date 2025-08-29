# Portfolio

Minimal portfolio with Go backend and Preact frontend.

## Quick Start

### Development

```bash
pnpm install
pnpm go-install

# Source environment variables (required for AWS S3 functionality)
source .env

pnpm dev
```

Access at: **http://localhost:3000** (auto-reload enabled)

### Production Build

```bash
pnpm run build
./run.sh
```

### Docker

```bash
docker-compose up --build
```
