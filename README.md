# Allocamelus

[![DeepSource](https://deepsource.io/gh/Allocamelus/Allocamelus.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/Allocamelus/Allocamelus/?ref=repository-badge)

## Description

Allocamelus is a (work in progress) self hostable social media platform

## Running

See [Basic.md](./examples/Basic.md) or [DockerCompose.md](./examples/DockerCompose.md)

## Features

- [ ] Account
  - [x] Signup
    - [x] Account creation
    - [x] Email verification
  - [x] Login
    - [x] Password auth (Argon2id)
      - Posable alternative (Augmented PAKE?)
    - [x] Remember me
    - [ ] Session private key restore (w/ password)
  - [x] Logout
  - [ ] Delete
  - [ ] Password reset
    - [ ] Api
    - [ ] Web app
  - [ ] Recovery
    - [ ] Old data (w/ old password or backup key)
    - [ ] Backup key (30 days?)
  - [ ] Key management
    - [x] Private key stored encrypted (w/ user password (Argon2id KDF))
    - [ ] View/Manage keys
      - [ ] Api
      - [ ] Web app
  - [ ] Events
    - [x] Logged
      - [x] Failed Login
      - [ ] Successful Login
      - [x] Password reset
  - [x] Move crypto to client
    - [ ] Alterative to web app
- [x] User
  <details>
  <summary>Click to expand</summary>
  - [x] Profile
    - [x] Bio
    - [x] Avatar
    - [x] Public/Private (Private by default)
    - [x] Feed (Post only)
      - Possible
        - [ ] Comments?
  - [x] Follow

  </details>

- [ ] Post
  - [x] Create
    <details>
    <summary>Click to expand</summary>
    - [x] Content
    - [x] Images
      - [x] Alts
    - Possible
      - [ ] Video

    </details>

  - [x] Publish
  - [ ] Update
    - [ ] Content
      - [x] Api
      - [ ] Web app
    - [ ] Images
  - [x] Comments
  - Possible
    - [ ] Mentions/Tagging
    - [ ] Hashtags

- [ ] Feed
  - [x] Followed users'
    - [x] Posts
    - [ ] Comments
  - Possible
    - [ ] Explore public users
- [ ] App
  - [ ] Api
    - [ ] Api Specification
      - [ ] Open api (WIP [Repo](https://github.com/Allocamelus/api-description))
    - [x] Rate limiting (Sliding Window | 690 request / 420 sec)
  - [ ] Federation
