# AES-256-GCM GUI Application

A beautiful, cross-platform desktop application for AES-256-GCM encryption and decryption.

![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Windows%20%7C%20Linux-lightgrey)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue)

## Features

- **Secure Key Generation**: Generate cryptographically secure 256-bit AES keys
- **AES-256-GCM Encryption**: Industry-standard authenticated encryption
- **User-Friendly Interface**: Clean, modern GUI built with Wails
- **Cross-Platform**: Works on macOS (Intel & Apple Silicon), Windows, and Linux
- **Base64 Output**: Encrypted data encoded in base64 for easy storage and transmission

## Screenshots

The application provides a simple interface with sections for:
- Key generation
- Encryption
- Decryption
- Clear error messages

## Quick Start

### Prerequisites

- Go 1.19 or later
- Node.js
- Platform-specific dependencies (see [BUILD_INSTRUCTIONS.md](BUILD_INSTRUCTIONS.md))

### Installation

1. Clone this repository
2. Install Wails CLI:
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

### Development

Run in development mode with hot reload:

```bash
wails dev
```

### Building

For detailed build instructions for all platforms, see [BUILD_INSTRUCTIONS.md](BUILD_INSTRUCTIONS.md).

Quick build for your current platform:

```bash
wails build
```

**Note:** macOS apps must be built on macOS machines due to Wails limitations.

## Usage

1. **Generate a Key**
   - Click "Generate New Key" button
   - Save the generated key securely

2. **Encrypt Text**
   - Enter or generate an encryption key
   - Type your message in the encryption box
   - Click "Encrypt"
   - Copy the base64-encoded result

3. **Decrypt Text**
   - Enter the same encryption key
   - Paste the encrypted base64 text
   - Click "Decrypt"
   - View your original message

## Security

- Uses AES-256-GCM (Galois/Counter Mode) for authenticated encryption
- Unique random nonce generated for each encryption
- Keys are 256-bit (32 bytes, 64 hex characters)
- Never stores keys persistently
- Industry-standard cryptographic implementation

## Technology Stack

- **Backend:** Go with standard crypto libraries
- **Frontend:** Vanilla JavaScript, HTML, CSS
- **Framework:** Wails v2 (Go + Web Technologies)
- **Build Tool:** Vite

## Project Structure

```
aes-gui/
├── app.go              # Backend encryption logic
├── main.go             # Application entry point
├── frontend/
│   ├── index.html      # UI structure
│   └── src/
│       ├── main.js     # Frontend logic
│       ├── app.css     # Styles
│       └── style.css   # Global styles
└── build/              # Build output
```

## Contributing

This is a utility tool. Contributions are welcome for:
- UI improvements
- Additional encryption algorithms
- File encryption support
- Bug fixes

## Troubleshooting

### macOS: App won't open ("damaged")
```bash
xattr -cr build/bin/aes-gui.app
```

### Linux: Missing dependencies
```bash
sudo apt install libgtk-3-0 libwebkit2gtk-4.0-37
```

For more troubleshooting, see [BUILD_INSTRUCTIONS.md](BUILD_INSTRUCTIONS.md).

## Related Projects

- [aes-cli](../aes-cli) - Command-line version of this tool

## License

This is a utility tool for educational and practical encryption/decryption purposes.

**Security Notice:** Keep your encryption keys secure and never commit them to version control.
