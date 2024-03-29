# EvilTransmission

## 1. Introduction

This is a side project in my leisure weekend time. It aims at developing full-stack application that is able to exchange data (mostly _evil_ images and videos).

### 1.1. System architecture

![Basic blocks](pictures/eviltrans-general-block.png "EvilTrans basic blocks")

### 1.2. System sequence diagram

![Sequence diagram](pictures/eviltrans-sequence-diagram.png)

## 2. Client

**Flutter** is used to develop smartphone application.
Please visit this branch [Client](https://github.com/tommyjohn1001/EvilTransmission/tree/FE) for more information.

## 3. Server

At first, **Python** is used to develop server side, but as I was introduced about _Golang_, I was totally persuaded. Not only the elegance but also the speed, the siplicity surprise me a lot.

Please visit this branch [Server](https://github.com/tommyjohn1001/EvilTransmission/tree/BE/server-golang) for more information.

## 4. Features

### 5.1. Stage 1: Basic components

This stage include very basic functions for both FE and BE.

- [x] Design API endpoints
- [x] Back-end

  - [x] Implement endpoints
  - [x] Implement server using Postgres

- [ ] Front-end
  - [ ] Implement simple UI that selects file from smartphone to upload
  - [ ] Fetch thumbnail, fetch video

### 5.2. Stage 2: Security improvement

This stage is expected to leverage security of system including protocol, data storage...

- 2-step encryption: encrypt payload and encrypt message itself
- Resources stored in system are encrypted also

- [ ] Back-end

  - [ ] Implement encrypt all data, decrypt on-demand

- [ ] Front-end
  - [ ] Exchange keys
  - [ ] Encrypt/Decrypt message before/after sending
  - [ ] Decrypt payload

### 5.3. Stage 3: UI and performance improvement

This stage aims at improving UI at front-end and performance in genenral

- [ ] Front-end
  - [ ]
- [ ] Back-end
  - [ ] Send file by chunks
  - [x] Apply DL-based feature extractor to filter out similar image/video
