# Arch Linux Development Environment Setup

This guide helps you set up a local development environment for the `mono-repo-release` project on Arch Linux.

## Prerequisites
- Arch Linux (or compatible, e.g., Manjaro)
- sudo privileges

## 1. Update System
```bash
sudo pacman -Syu
```

## 2. Essential Packages
Install the following packages:
```bash
sudo pacman -S \
  zsh starship \
  vlc brave \
  go go-tools delve gopls \
  code \
  zsh-autosuggestions zsh-syntax-highlighting fzf \
  goreleaser lcov genhtml golangci-lint \
  bazel git make
```

## 3. Optional: Font Installer (getnf)
```bash
paru -S getnf
# or via script:
curl -fsSL https://raw.githubusercontent.com/getnf/getnf/main/install.sh | bash
```

## 4. Oh My Zsh (optional, for zsh users)
```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

## 5. GitHub CLI (optional)
```bash
sudo pacman -S github-cli
```

## 6. Project Setup
Clone your repository and install Go dependencies:
```bash
git clone git@github.com:omargallob/mono-repo-release.git
cd mono-repo-release
make check-deps
```

## 7. Build & Test
```bash
make build
make test
```

---

**Tip:**
- Use `zsh` as your shell for a better experience.
- Use `starship` for a modern prompt.
- Use `code` (VS Code) for editing.

---

This setup covers all major tools and packages found in your bash history for local development on Arch Linux.
