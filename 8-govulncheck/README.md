# Find and fix vulnerable dependencies with govulncheck

This is a basic example of the **govulncheck** tool, which scans a project's dependencies for known vulnerabilities. In this tutorial, we will find a vulnerable dependency and fix it.

---

### Prerequisites
* **Go** latest version 
* Command terminal

---


### Tutorial steps
Download dependencies
```bash
$ go mod tidy
```

Downgrade dependency to vulnerable version
```bash
$ go get golang.org/x/text@v0.3.5
```

Check for vulnerabilities 
```bash
$ govulncheck ./...
```

Update dependency to fixed version
```bash
$ go get golang.org/x/text@v0.3.7
```

Check for absence of vulnerabilities 
```bash
$ govulncheck ./...
```
