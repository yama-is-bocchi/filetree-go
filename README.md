# filetree-go
A Go-based CLI tool that displays any directory in a tree structure.

## Installation
### Using Go
```bash
go install github.com/yama-is-bocchi/filetree-go@latest
```
Make sure your Go binary directory is included in your PATH.

### Using a Prebuilt Binary
1. Download the appropriate archive for your platform from the [GitHub Releases page](https://github.com/yama-is-bocchi/filetree-go/releases).
2. Extract the archive.
3. Move the binary to a directory included in your PATH.

Example for Linux/macOS:
```bash
tar -xzf filetree-go_linux_amd64.tar.gz
chmod +x filetree-go
sudo mv filetree-go /usr/local/bin/
```

### Verify the installation:
```bash
filetree-go --help
```
