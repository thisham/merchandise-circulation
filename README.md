<p align="center">
  <h1 align="center">Merchandise API</h1>

  <p align="center">
    REST API of Clinic Application
  </p>
</p>

<div align="center">
  
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![semantic-release: angular][semantic-badge]][semantic-url]
[![codecov][codecov-shield]][codecov-url]
  
</div>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#deployment-guide">Deployment Guide</a></li>
      </ul>
    </li>
    <li><a href="#given-rules">Given Rules</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project
As a merchant, we need to register and check the transaction of merchandise 
status. With merchandise management API, this will be possible to list and data 
the merchandise transactions. And, we also can check what are the merchandise
in the warehouse.

This project was given as an assignment and an exercise instead of an assignment 
on the alta.id platform, because the expert class time was running out.

## Getting Started

### Prerequisites

When you're going to contribute or compile, you'll need at least:
  - Go 1.17.5+
  - MySQL 8.0.x
  - MongoDB 5.x

### Installation or Configure

```bash
# clone if you don't have the code base
$ git clone git@github.com:thisham/merchandise-circulation.git

# tidy up and run
$ go mod tidy

$ go run main.go
```

### Deployment Guide

#### Self-Build

```sh
# clone if you don't have the code base
$ git clone git@github.com:thisham/merchandise-circulation.git

# install the dependencies
$ go mod download

# compile 
$ go build -o main
```

## Given Rules

[Translated into English]

Merchandise Stock System, can be written in MVC or Clean Architecture. The 
assessment of these points will replace the assessment of the REST API and 
Middleware assignment.
- Incoming Transaction (50%)
  - Validate if the merchandise record is available 
    - The stock will be increased (15%)
    - Update the merchandise data if changed (20%)
  - ... if not available
    - Add new merchandise record (15%)
- Outcoming Transaction (50%)
  - Check a valid transaction, only accept the transaction of amount of 
    merchandise less than the available stock (25%)
  - Count total price of outcoming merchandise, and add the service fee by 20%
    (25%)

The assessment of these points will replace the assessment of the Middleware 
assignment.
- Authentication (50%)
- Logging (50%)

The assessment of these points will replace the assessment of the Unit Test 
assignment.
- Unit test (at least 80% coverage)

These assignment will be presented on next Monday, April 18th 2022. Because, on 
that day there will be 2 sessions, the total will be 3 hours.

## Features

### 🛒 Merchandise
The added merchandise will be listed here, so you can have an eye in the 
warehouse to check the merchandise statuses. The views can be detailed or will 
be listed sortly.

### 💳 Transaction
When the merchandise are going out or going in, there will be a note for doing 
the transaction. So, the merchandise can be fully tracked and will be sortlisted 
completely.

## Licenses

All of this source code are licensed under the MIT licenses. Although this 
source code is related to the Alterra Academy.

## Contact

- Hamdan Yuwafi - <yuwafi.hamdan365@gmail.com>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[stars-shield]: https://img.shields.io/github/stars/thisham/merchandise-circulation.svg?style=for-the-badge
[stars-url]: https://github.com/thisham/merchandise-circulation/stargazers
[issues-shield]: https://img.shields.io/github/issues/thisham/merchandise-circulation.svg?style=for-the-badge
[issues-url]: https://github.com/thisham/merchandise-circulation/issues
[license-shield]: https://img.shields.io/github/license/thisham/merchandise-circulation.svg?style=for-the-badge
[license-url]: https://github.com/thisham/merchandise-circulation/blob/master/LICENSE
[semantic-badge]: https://img.shields.io/badge/semantic--release-angular-e10079?style=for-the-badge&logo=semantic-release
[semantic-url]: https://github.com/semantic-release/semantic-release
[codecov-shield]: https://img.shields.io/codecov/c/gh/thisham/merchandise-circulation?label=CODECOV&logo=codecov&style=for-the-badge&token=YL5V1QWP2I
[codecov-url]: https://codecov.io/gh/thisham/merchandise-circulation
