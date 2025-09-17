# system-monitor

Production-ready system metrics REST API (Go).

## Requirements
- Go 1.20+
- Optional: Docker socket (/var/run/docker.sock) for container info
- Optional: `nvidia-smi` in PATH for GPU info
- Optional: lm-sensors / kernel sensors for temperatures

## Build & Run
```bash
make build
./bin/sysinfo
```

Or run in Docker:
```bash
docker build -t sysinfo:latest .
docker run -p 3000:3000 --rm \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  --pid=host \
  --cap-add SYS_ADMIN \
  sysinfo:latest
```

## API
- GET /healthz
- GET /api/system/cpu
- GET /api/system/memory
- GET /api/system/disk
- GET /api/system/network
- GET /api/system/processes
- GET /api/system/info
- GET /api/system/gpu
- GET /api/system/containers

## Notes
- Some features require host-level permissions/tools; see Requirements.
- Replace module name in go.mod if you want a proper module path.
