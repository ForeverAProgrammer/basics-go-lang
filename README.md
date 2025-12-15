# Go Basics

This repository contains basic Go programming examples and demonstrations.

## What is Go?

Go (also known as Golang) is an open-source programming language created by Google in 2009. It's designed for simplicity, efficiency, and excellent support for concurrent programming. Go is statically typed, compiled, and known for its fast compilation times and built-in concurrency features.

### Key Features
- Fast compilation and execution
- Simple and clean syntax
- Built-in concurrency with goroutines
- Strong standard library
- Garbage collection
- Cross-platform support

## How to Install Go

### Ubuntu/Debian
```bash
sudo apt install golang-go
```

### Verify Installation
After installation, verify that Go is installed correctly:
```bash
go version
```

### macOS
Using Homebrew:
```bash
brew install go
```

### Windows
Download the installer from [golang.org/dl](https://golang.org/dl) and run it.

### Alternative: Official Binary
For the latest version or other platforms, download from the official site:
- Visit [golang.org/dl](https://golang.org/dl)
- Download the appropriate package for your system
- Follow the installation instructions for your platform

## Getting Started

Once Go is installed, you can run any of the example programs in this repository:

```bash
cd println
go run println-demo.go
```

## Repository Structure

- [println/](println/) - Demonstrates the `fmt.Println()` function

## Resources

- [Official Go Documentation](https://golang.org/doc)
- [Go by Example](https://gobyexample.com)
- [A Tour of Go](https://tour.golang.org)
