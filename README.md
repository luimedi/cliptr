# Cliptr

Cliptr is a lightweight terminal-based utility (TUI) to dynamically monitor and transform text copied to your clipboard.

## Features

- **TUI Controls**: Easily activate or deactivate transformation rules via an interactive terminal checklist.
- **Auto-trim**: Automatically removes leading and trailing whitespaces from copied text.
- **Capitalization**: Capitalizes words if the copied text is fully uppercase.
- **Live Monitoring**: Listens to clipboard changes in the background and applies active transformations instantly.

## Installation

For **Windows** and **macOS**, no external system dependencies are required. You can install it directly.

### Prerequisites (Linux only)

Since Cliptr uses X11 to communicate with the clipboard on Linux, you must install the X11 development libraries first:

- **Ubuntu / Debian / Mint:**
  ```bash
  sudo apt-get install libx11-dev
  ```
- **Fedora / RHEL:**
  ```bash
  sudo dnf install libX11-devel
  ```

### Install with Go

To install the latest version globally on your system, run:

```bash
go install github.com/luimedi/cliptr@latest
```

> [!NOTE]
> Make sure your `PATH` environment variable includes your Go binary path (typically `$HOME/go/bin` or `~/go/bin`). If not, you can add it to your `~/.bashrc` or `~/.zshrc`:
> ```bash
> export PATH=$PATH:$(go env GOPATH)/bin
> ```

## Usage

Simply run the command in your terminal:

```bash
cliptr
```

Use the **Spacebar** or **Enter** key to toggle rules, **Up/Down** arrows to navigate, and **Q** to exit.

## License

Cliptr is open-sourced software licensed under the [GPL-3.0 license](LICENSE).
