[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/hleejr/Auto-Cmd">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Auto Command</h3>

  <p align="center">
    Golang CLI which makes use of text files to run multiple commands within the terminal at once
    <br />
    <a href="https://github.com/hleejr/Auto-Cmd"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://docs.google.com/presentation/d/1w7uBAsoznALcbNteGXf-v_b4sSUu6_At69wOEWYC7Yg/edit?usp=sharing">View Demo</a>
    ·
    <a href="https://github.com/hleejr/Auto-Cmd/issues">Report Bug</a>
    ·
    <a href="https://github.com/hleejr/Auto-Cmd/issues">Request Feature</a>
  </p>
</p>

[![Go Report Card](https://goreportcard.com/badge/github.com/hleejr/Auto-Cmd)](https://goreportcard.com/report/github.com/hleejr/Auto-Cmd)

*[AutoCmd Details & Demo](https://docs.google.com/presentation/d/1w7uBAsoznALcbNteGXf-v_b4sSUu6_At69wOEWYC7Yg/edit?usp=sharing)*

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Utility Usage](#utility-usage)

## Project Structure

```bash
📂 Auto-Cmd
├── 📂 images
├── README.md
├── autocmd.go
├── git.txt
├── go.sum
└── go.mod
```

## Getting Started

Run each command line-by-line in your terminal to set up the utility:
```
$ git clone https://github.com/hleejr/Auto-Cmd.git
$ cd Auto-Cmd
$ git remote rm origin
```
Then add the link to your repository as the new origin

## Utility Usage
When you have multiple shell commands needed for various setups, AutoCmd makes use of text files to run multiple commands with just one shell entry. The one example I included is a text file for git commands to create a git commit. However, by creating a new ".txt" file and adding the desired commands with newlines after each entry, any combination of shell commands should be compatible.

Once you have the utility installed it can be used by running the following commands within the terminal from inside the necessary directory:
```
go build
./autocmd --file --msg --branch
```
Alternatively you can download the utility to be used anywhere on your system and not just in one specific directory:
```
go install
autocmd --file --msg --branch
```
- The **file** flag is used to determine which ".txt" file is used to pull the desired commands
- The **msg** flag is used to pass a message along with the git commit command
- The **branch** flag is used to determine which branch of the repository you want to push your commit to


As these flags may not be needed and other flags may be useful depending on the particular commands being used, the flags within program can easily be altered to best suite your purposes.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/hleejr/Auto-Cmd.svg?style=for-the-badge
[contributors-url]: https://github.com/hleejr/Auto-Cmd/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/hleejr/Auto-Cmd.svg?style=for-the-badge
[forks-url]: https://github.com/hleejr/Auto-Cmd/network/members
[stars-shield]: https://img.shields.io/github/stars/hleejr/Auto-Cmd.svg?style=for-the-badge
[stars-url]: https://github.com/hleejr/Auto-Cmd/stargazers
[issues-shield]: https://img.shields.io/github/issues/hleejr/Auto-Cmd.svg?style=for-the-badge
[issues-url]: https://github.com/hleejr/Auto-Cmd/issues
[license-shield]: https://img.shields.io/github/license/hleejr/Auto-Cmd.svg?style=for-the-badge
[license-url]: https://github.com/hleejr/Auto-Cmd/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/henry-bowe-jr-31498916a/