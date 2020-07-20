# 🔗 AutoCmd

[![Go Report Card](https://goreportcard.com/badge/github.com/hleejr/makesite)](https://goreportcard.com/report/github.com/hleejr/makesite)
*[AutoCmd Details & Demo](https://docs.google.com/presentation/d/1w7uBAsoznALcbNteGXf-v_b4sSUu6_At69wOEWYC7Yg/edit?usp=sharing)*

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Utility Usage](#utility-usage)

## Project Structure

```bash
📂 Autocmd
├── README.md
├── autocmd.go
├── git.txt
├── go.sum
└── go.mod
```

## Getting Started

Run each command line-by-line in your terminal to set up the utility:

```bash
$ cd ~/go/src
$ git clone https://github.com/hleejr/Auto-Commit.git
$ cd makesite
$ git remote rm origin
```
Then add the link to your repository as the new origin

## Utility Usage
When you have multiple shell commands needed for various setups, AutoCmd makes use of textfiles to run multiple commands with just one shell entry. The one example I included is a text file for git commands to create a git commit. However, by creating a new ".txt" file and adding the desired demands, with newlines after each entry, any combination of shell commands should be compatible.
