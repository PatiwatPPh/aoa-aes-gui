# AES-256-GCM GUI Application - Build Instructions

This is a desktop GUI application for AES-256-GCM encryption/decryption built with Wails.

## Features

- Generate secure AES-256 keys
- Encrypt text with AES-256-GCM
- Decrypt encrypted text
- Beautiful, modern user interface
- Cross-platform (Windows, macOS, Linux)

## Prerequisites

### For macOS

1. **Install Xcode Command Line Tools:**
   ```bash
   xcode-select --install
   ```

2. **Install Go (1.19 or later):**
   ```bash
   brew install go
   ```

3. **Install Node.js:**
   ```bash
   brew install node
   ```

4. **Install Wails CLI:**
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

### For Linux

1. **Install dependencies:**
   ```bash
   # Ubuntu/Debian
   sudo apt install build-essential pkg-config libgtk-3-dev libwebkit2gtk-4.0-dev

   # Fedora
   sudo dnf install gtk3-devel webkit2gtk3-devel

   # Arch Linux
   sudo pacman -S webkit2gtk gtk3
   ```

2. **Install Go (1.19 or later):**
   ```bash
   # Download from https://go.dev/dl/
   ```

3. **Install Node.js:**
   ```bash
   sudo apt install nodejs npm  # Ubuntu/Debian
   ```

4. **Install Wails CLI:**
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

### For Windows

1. **Install Go:** Download from https://go.dev/dl/
2. **Install Node.js:** Download from https://nodejs.org/
3. **Install Wails CLI:**
   ```powershell
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

## Building the Application

### Development Mode

Run the app in development mode with hot reload:

```bash
cd /path/to/aes-gui
wails dev
```

### Production Build

#### macOS (ARM64 - M1/M2/M3/M4)

**Note:** Must be built on a macOS ARM64 machine.

```bash
cd /path/to/aes-gui
wails build -platform darwin/arm64
```

The app will be created at: `build/bin/aes-gui.app`

#### macOS (Intel)

```bash
cd /path/to/aes-gui
wails build -platform darwin/amd64
```

#### Windows

```bash
cd /path/to/aes-gui
wails build -platform windows/amd64
```

#### Linux

```bash
cd /path/to/aes-gui
wails build -platform linux/amd64
```

### Universal macOS Build

To create a universal binary that works on both Intel and Apple Silicon:

```bash
wails build -platform darwin/universal
```

## Output Location

After building, the application will be located in:
- **macOS:** `build/bin/aes-gui.app`
- **Windows:** `build/bin/aes-gui.exe`
- **Linux:** `build/bin/aes-gui`

## Running the Application

### macOS
```bash
open build/bin/aes-gui.app
```

Or double-click the app in Finder.

### Windows
Double-click `aes-gui.exe`

### Linux
```bash
./build/bin/aes-gui
```

## Usage

1. **Generate a Key:**
   - Click "Generate New Key" button
   - The key will appear and be automatically filled into the key input field
   - Save this key securely!

2. **Encrypt Text:**
   - Enter your encryption key (or generate one)
   - Type the text you want to encrypt in the "Encrypt" section
   - Click "Encrypt"
   - The encrypted text will appear in base64 format

3. **Decrypt Text:**
   - Enter the same encryption key used for encryption
   - Paste the encrypted text (base64) in the "Decrypt" section
   - Click "Decrypt"
   - The original text will appear

## Security Notes

- **Never share your encryption key**
- Keys are 64 hexadecimal characters (32 bytes for AES-256)
- Each encryption uses a unique random nonce for security
- GCM mode provides authenticated encryption
- Keys are only stored in memory during the session

## Troubleshooting

### macOS: "App is damaged and can't be opened"

This happens because the app is not signed. Run:
```bash
xattr -cr build/bin/aes-gui.app
```

### Linux: Missing GTK dependencies

Install the required dependencies:
```bash
sudo apt install libgtk-3-0 libwebkit2gtk-4.0-37
```

### Cross-compilation limitations

- **macOS apps must be built on macOS** (Wails limitation)
- Windows apps can be built on Windows or Linux with MinGW
- Linux apps can be built on Linux

## Development

To modify the app:

1. **Backend (Go):** Edit `app.go`
2. **Frontend (HTML/CSS/JS):** Edit files in `frontend/src/`
3. **Run in dev mode:** `wails dev`

The app will automatically reload when you make changes.

## Project Structure

```
aes-gui/
├── app.go              # Backend Go code with encryption logic
├── main.go             # Application entry point
├── frontend/
│   ├── index.html      # Main HTML structure
│   └── src/
│       ├── main.js     # Frontend JavaScript logic
│       ├── app.css     # Application styles
│       └── style.css   # Global styles
├── build/              # Build output directory
└── wails.json          # Wails configuration
```

## License

This is a utility tool for encryption/decryption. Use responsibly and keep your keys secure.
