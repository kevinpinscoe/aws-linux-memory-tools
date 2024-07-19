# Purpose

Run on AWS Linux and reports if free memory is near zero.

An example is if PDF generation is going on free meory is nearly or is exhausted.

This would log to CoudWatch which could then make a decison to scale up by a single instance.

# Install

* ```go mod init linux-memory-tool```
* ```go mod tidy```

# Deployment

Deployed as a part of ec2 User Agent?
