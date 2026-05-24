# Purpose

Run on AWS Linux and reports if free memory is near zero.

An example might be when free memory is nearly exhausted during PDF generation.

This would log to CoudWatch which could then make a decison to scale up by a single instance.

# Install

## Package managers

### Homebrew (macOS/Linux)

```bash
brew tap kevinpinscoe/homebrew-tap
brew install aws-linux-memory-tools
```

### APT (Debian/Ubuntu)

```bash
curl -sL https://kevinpinscoe.github.io/apt/gpg.key \
  | sudo gpg --dearmor -o /etc/apt/keyrings/kevinpinscoe.gpg

echo "deb [signed-by=/etc/apt/keyrings/kevinpinscoe.gpg] \
  https://kevinpinscoe.github.io/apt stable main" \
  | sudo tee /etc/apt/sources.list.d/kevinpinscoe.list

sudo apt update
sudo apt install aws-linux-memory-tools
```

### DNF (Fedora/RHEL)

```bash
sudo curl -fsSL https://kevinpinscoe.github.io/rpm/kevinpinscoe.repo \
  -o /etc/yum.repos.d/kevinpinscoe.repo
sudo dnf install aws-linux-memory-tools
```

## Compile from source

* ```go mod init linux-memory-tool```
* ```go mod tidy```

# Deployment

Deployed as a part of ec2 User Agent?
