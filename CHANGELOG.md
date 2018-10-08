# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2018-10-08
### Added
   - Shruggies! (#15, @wadadli)
   - CircleCI support for automating tests (#8, @wadadli)
   - Added a changelog

### Changed

### Fixed
  - Fix crashes on very large titles (#22, @jrwren)
  - Trim spaces from incoming messages when matching single word commands (#17, @bbriggs)
  - Fix build time variable injection for seting version, git SHA, and branch (#6, #7, #10, #11, #12, @bbriggs)
  - Certain URL schemes were breaking and crashing the bot (#3, @bbriggs)
  - Properly vendor _all_ dependencies (#4, #5, @bbriggs)
  - Bot no longer crashes on bad HTML/URLs because of missing error handler (#2, @bbriggs)

## [0.1.0] - 2018-09-28
### Added
  - Initial semver'd release
  - Spelling fix in readme (#1)
  - Report titles of URLs posted in chat
  - Basic idle tracker
  - !info command displays build, semver, and branch

### Changed

### Fixed
