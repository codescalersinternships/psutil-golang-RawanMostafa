# Psutils

This repository implements a human-friendly lib for querying processes, memory info and cpu info

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

Get the package

   ```bash
     go get github.com/codescalersinternships/psutil-golang-RawanMostafa
   ```
Import the package

  ```go
     import github.com/codescalersinternships/psutil-golang-RawanMostafa/pkg
  ```

## Usage

  The psutils package offers different APIs like:

### 1. CPU related APIs

```go
  cpuInfo, err := psutils.GetCpuInfo()
    if err != nil {
        //handle the error
    }
```

### 2. Memory related APIs

```go
   memInfo, err := psutils.GetMemInfo()
    if err != nil {
        //handle the error
    }
```
### 3. Processes related APIs

1. ```go
	procs, err := psutils.GetProcessList()
    if err != nil {
        //handle the error
    }
   ```

2. ```go
	procDetails, err := psutils.GetProcessDetails(1)
    if err != nil {
        //handle the error
    }
   ```