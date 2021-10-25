# Allocamelus

[![DeepSource](https://deepsource.io/gh/Allocamelus/Allocamelus.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/Allocamelus/Allocamelus/?ref=repository-badge)

## Description

Allocamelus is a (work in progress) self hostable social media platform

### Features:

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
    - [x] Api
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
  - [ ] Move crypto to client
    - Preferably not a web app
    - But, seeing Proton Mail's browser crypto...
- [x] User 
  - [x] Profile
    - [x] Bio
    - [x] Avatar
    - [x] Public/Private (Private by default)
    - [x] Feed (Post only)
      - [ ] Comments?
  - [x] Follow
- [ ] Post
  - [x] Create
    - [x] Content
    - [x] Images
      - [x] Alts
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
- [ ] Api
  - [ ] Specification
    - [ ] Open api
  - [ ] Rate limiting (Test [fiber/middleware/limiter](https://docs.gofiber.io/api/middleware/limiter))
