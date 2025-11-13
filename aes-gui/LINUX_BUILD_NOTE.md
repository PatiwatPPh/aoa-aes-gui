# Linux Build Notes

## Issue with webkit2gtk-4.0 vs webkit2gtk-4.1

The current Ubuntu version (plucky/25.04) has migrated from `webkit2gtk-4.0` to `webkit2gtk-4.1`, but Wails v2.10.2 still references the older version.

### Workaround Options:

#### Option 1: Install webkit2gtk-4.0 compatibility package

If available on your system:
```bash
sudo apt install libwebkit2gtk-4.0-dev
```

#### Option 2: Build on an older Ubuntu version

Use Ubuntu 22.04 or 24.04 which still has webkit2gtk-4.0:
```bash
# On Ubuntu 22.04/24.04
sudo apt install -y build-essential pkg-config libgtk-3-dev libwebkit2gtk-4.0-dev
cd /path/to/aes-gui
wails build
```

#### Option 3: Build on macOS

Since you need the macOS ARM64 version anyway, the easiest approach is to:

1. Transfer the entire `aes-gui/` directory to your macOS machine
2. Install prerequisites on macOS:
   ```bash
   # Install Homebrew if not already installed
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

   # Install dependencies
   brew install go node

   # Install Wails CLI
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

3. Build the app:
   ```bash
   cd /path/to/aes-gui

   # For development/testing
   wails dev

   # For production build (macOS ARM64)
   wails build -platform darwin/arm64
   ```

4. The app will be at: `build/bin/aes-gui.app`

#### Option 4: Use Docker with Ubuntu 22.04

Create a Dockerfile:
```dockerfile
FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    curl \
    git

# Install Go
RUN curl -OL https://go.dev/dl/go1.21.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz && \
    rm go1.21.0.linux-amd64.tar.gz

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/root/go"

# Install Wails
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

WORKDIR /app
```

Then build:
```bash
docker build -t wails-builder .
docker run -v $(pwd)/aes-gui:/app wails-builder wails build
```

## Current Status

The application code is complete and ready to build. The only issue is the webkit2gtk version mismatch on newer Ubuntu versions.

**Recommended approach**: Build on macOS since that's your target platform anyway.

## Application Features

The GUI application is fully functional with:
- ✅ AES-256-GCM encryption backend (Go)
- ✅ Modern web-based UI (HTML/CSS/JavaScript)
- ✅ Key generation functionality
- ✅ Encrypt/Decrypt operations
- ✅ Error handling and validation
- ✅ Auto-fill between encrypt and decrypt sections

## Testing Without Building

You can test the functionality using the CLI version:
```bash
cd /path/to/aes-cli
./aes-cli genkey
./aes-cli encrypt -key <your-key> -text "Hello World"
./aes-cli decrypt -key <your-key> -text <encrypted-output>
```

The GUI uses the exact same encryption logic.
