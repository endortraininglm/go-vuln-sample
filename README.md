# go-vuln-sample

This is a simple Go project that demonstrates the use of the [gosnowflake](https://github.com/snowflakedb/gosnowflake) package at version v1.6.18, which is vulnerable to [CVE-2023-34231](https://nvd.nist.gov/vuln/detail/CVE-2023-34231).

## Project Overview

This project serves as a demonstration of how to reference a specific vulnerable dependency in a Go application. It's designed to be minimal but functional, allowing security scanning tools to detect the vulnerability. The project consists of:

- `main.go`: A simple Go application that imports and uses the vulnerable gosnowflake package
- `go.mod`: The Go module file that specifies the dependency on gosnowflake v1.6.18
- `go.sum`: The checksum file that ensures dependency integrity

## Vulnerability Information

- **CVE ID**: CVE-2023-34231
- **Affected Version**: gosnowflake < v1.6.19
- **Description**: The vulnerability is related to improper validation of server certificates in the gosnowflake package.
- **Impact**: This vulnerability could potentially allow attackers to perform man-in-the-middle attacks due to improper certificate validation.

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

The scan should identify the vulnerable gosnowflake dependency (v1.6.18) and report CVE-2023-34231. This demonstrates that the security scanning tools are correctly identifying vulnerable dependencies in your Go projects.
