# Purpose

Run on AWS Linux and reports if free memory is near zero.

An example might be when free memory is nearly exhausted during PDF generation.

This would log to CoudWatch which could then make a decison to scale up by a single instance.

# Install

* ```go mod init linux-memory-tool```
* ```go mod tidy```

# Deployment

Deployed as a part of ec2 User Agent?
