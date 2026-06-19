# Purpose

Run on AWS Linux and reports if free memory is near zero.

An example might be when free memory is nearly exhausted during PDF generation.

This would log to CloudWatch which could then make a decision to scale up by a single instance.

# Install

## Package managers

### Amazon Linux 2023

Amazon Linux 2023 is RPM-based and uses DNF. Add the RPM repo and install:

```bash
sudo curl -fsSL https://kevinpinscoe.github.io/rpm/kevinpinscoe.repo \
  -o /etc/yum.repos.d/kevinpinscoe.repo
sudo dnf install aws-linux-memory-tools
```

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

Deployed as part of EC2 User Data?

## Contributing & Reporting Issues

Bug reports, feature requests, security disclosures, and contributions are all
welcome. I keep these guidelines in one place for all my projects:

- **How to contribute or report an issue:** https://github.com/kevinpinscoe/how-to-contribute
- **Report a security vulnerability:** do not open a public issue. Use the
  **"Report a vulnerability"** button on this repository's **Security** tab, or
  see the [security policy](https://github.com/kevinpinscoe/how-to-contribute/blob/main/SECURITY.md).
