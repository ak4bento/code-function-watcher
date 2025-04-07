# 🧠 code-function-watcher

A powerful CLI tool for **Go developers** to scan, detect, and prevent duplicated or unused functions in your codebase.  
Designed for **code quality automation**, especially useful in **team collaboration**, **code review**, and **CI/CD pipelines**.

---

## ✨ Features

- ✅ Detect all exported functions in a directory
- 🔁 Compare old and new snapshots to find **duplicated or similar functions**
- 🧹 Detect **unused functions** that are defined but never called
- 🚫 Ignore list support to skip known functions or third-party calls
- 🤖 GitHub CI/CD & PR auto-comment support

---

## 📦 Installation

### 1. Clone the Repo

```bash
git clone https://github.com/ak4bento/code-function-watcher.git
cd code-function-watcher
```

### 2. Install Dependencies

```bash
go mod tidy
```

---

## 🚀 Usage

### 🔍 Scan Functions

```bash
go run main.go scan <path> [-o output.json]
```

Example:

```bash
go run main.go scan ./ -o data/functions.json
```

> This will scan all Go files recursively and extract exported function names + locations.

---

### 🔁 Compare Function Snapshots

```bash
go run main.go compare <old.json> <new.json>
```

Example:

```bash
go run main.go compare data/functions-old.json data/functions-new.json
```

> It will print potentially duplicated or very similar functions with similarity percentage (e.g., 94.5%).

---

### 🧹 Detect Unused Functions

```bash
go run main.go unused <path> --defined <functions.json> [--ignore ignore.txt]
```

Example:

```bash
go run main.go unused ./ --defined data/functions.json --ignore ignore.txt
```

> Lists exported functions that were never called anywhere in your project.

---

## 🚫 Ignore List

Use a plain text file (e.g. `ignore.txt`) to exclude functions from being detected as "unused":

```
log.Println
fmt.Errorf
main
```

---

## 🧪 GitHub Actions (CI/CD)

Create a file `.github/workflows/scan.yml`:

```yaml
name: Code Function Watcher

on:
  pull_request:
    paths:
      - '**/*.go'

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.21

      - name: Install Dependencies
        run: go mod tidy

      - name: Run Function Comparison
        run: |
          go run main.go scan ./ -o data/functions-new.json
          go run main.go compare data/functions-old.json data/functions-new.json > result.txt

      - name: Upload Results
        uses: actions/upload-artifact@v3
        with:
          name: compare-result
          path: result.txt
```

---

## 💬 Auto-Comment to Pull Request (Bonus)

Install [peter-evans/create-or-update-comment](https://github.com/peter-evans/create-or-update-comment)

Extend your workflow:

```yaml
      - name: Post PR Comment
        uses: peter-evans/create-or-update-comment@v3
        with:
          issue-number: ${{ github.event.pull_request.number }}
          body: |
            🔁 **Function Comparison Report**
            ```
            $(cat result.txt)
            ```
```

---

## 📁 Project Structure

```
.
├── cmd/
│   ├── compare.go
│   ├── scan.go
│   └── unused.go
├── pkg/
│   ├── compare/
│   ├── exporter/
│   ├── scanner/
│   └── unused/
├── data/
│   └── functions-old.json
│   └── functions-new.json
├── ignore.txt
└── main.go
```

---

## 🧑‍💻 Author

Made with ❤️ by [@ak4bento](https://github.com/ak4bento)