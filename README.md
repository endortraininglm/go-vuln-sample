# go-vuln-sample

This is a simple Go project that demonstrates the use of:
1. The [gosnowflake](https://github.com/snowflakedb/gosnowflake) package at version v1.6.18, which is vulnerable to [CVE-2023-34231](https://nvd.nist.gov/vuln/detail/CVE-2023-34231).
2. The [etcd](https://github.com/etcd-io/etcd) client package at version v3.5.9, which is vulnerable to [GO-2024-2528](https://pkg.go.dev/vuln/GO-2024-2528) with alias [GHSA-j86v-2vjr-fg8f](https://github.com/advisories/GHSA-j86v-2vjr-fg8f).

## Project Overview

This project serves as a demonstration of how to reference specific vulnerable dependencies in a Go application. It's designed to be minimal but functional, allowing security scanning tools to detect the vulnerabilities. The project consists of:

- `main.go`: A simple Go application that imports and uses the vulnerable gosnowflake and etcd packages
- `go.mod`: The Go module file that specifies the dependencies on gosnowflake v1.6.18 and etcd client v3.5.9
- `go.sum`: The checksum file that ensures dependency integrity

## Vulnerability Information

### CVE-2023-34231 (gosnowflake)
- **CVE ID**: CVE-2023-34231
- **Affected Version**: gosnowflake < v1.6.19
- **Description**: The vulnerability is related to improper validation of server certificates in the gosnowflake package.
- **Impact**: This vulnerability could potentially allow attackers to perform man-in-the-middle attacks due to improper certificate validation.

### GO-2024-2528 (etcd)
- **GO ID**: GO-2024-2528
- **GitHub Security Advisory**: GHSA-j86v-2vjr-fg8f
- **Affected Version**: go.etcd.io/etcd/client/v3 < v3.5.10
- **Description**: The vulnerability is related to improper handling of certain requests in the etcd client package.
- **Impact**: This vulnerability could potentially allow attackers to cause denial of service or other security issues.

## Build Instructions

### Prerequisites

- Go 1.20 or later - This project requires Go version 1.20 or newer to ensure compatibility with all dependencies.
- Git - Required for cloning the repository (if not downloading directly).
- Internet connection - Needed to download the dependencies during the build process.

### Steps to Build

1. Clone the repository:
   ```
   git clone https://github.com/example/go-vuln-sample.git
   cd go-vuln-sample
   ```
   This step downloads the source code to your local machine and navigates to the project directory.

2. Download dependencies:
   ```
   go mod tidy
   ```
   This command analyzes the project's imports, downloads all required dependencies (including the vulnerable gosnowflake package), and updates the go.sum file with checksums to ensure dependency integrity.

3. Build the application:
   ```
   go build -o vuln-sample
   ```
   This compiles the Go code into an executable binary named 'vuln-sample'. The compilation process will include the vulnerable gosnowflake package.

4. Run the application (optional):
   ```
   ./vuln-sample
   ```
   Running the application will demonstrate that the code successfully imports and uses the gosnowflake package. The program doesn't actually connect to a Snowflake database but shows that the dependency is properly referenced.

## Scanning the Project

You can use various SCA (Software Composition Analysis) tools to scan this project and detect the vulnerable dependency. The scanning process typically involves analyzing the go.mod and go.sum files to identify dependencies and their versions.

### Recommended Scanning Tools:

- **Endor SCA**: Specifically mentioned by the client for inventory tracking of dependencies.
- **Snyk**: Offers comprehensive vulnerability scanning for Go projects.
- **OWASP Dependency-Check**: An open-source solution for identifying project dependencies and checking for known vulnerabilities.
- **GoSec**: A Golang security checker that can identify potential security issues in your code.

### Expected Scan Results:

The scan should identify the following vulnerable dependencies:
1. gosnowflake v1.6.18 - CVE-2023-34231
2. etcd client v3.5.9 - GO-2024-2528 (GHSA-j86v-2vjr-fg8f)

This demonstrates that the security scanning tools are correctly identifying vulnerable dependencies in your Go projects.
