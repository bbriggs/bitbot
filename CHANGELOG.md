# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [UNRELEASED] 
### Added
   - Prometheus endpoint and metrics (#70, bbriggs)
   - Help field added to NamedTriggers
   - New !help command lists all registered triggers (#69, bbriggs)
   - !help command also supports positional arguments and returns contents of Help field
   - Let bitbot help you decide things! Use `{nick} choose {option} or {option} [or {option}]...` 
   - Bitbot now adds an underscore to his nick when needed (#60, m-242)
   - Bitbot can now flip and unflip the table for you ! (#68, m-242)
   - Better tests for all triggers with conditions based on inbound messages

### Fixed
   - Fixed the plugin links's tests (#53, @m-242)
   - Specify plugins to load using the config file instead of modifying the source code. Whodathunkit.
     - When running as a bot (and not a library), all plugins are loaded by default.
     - Plugins must be of `bitbot.NamedTrigger` type. Sorry not sorry. (#69, bbriggs)
       - We depend on NamedTrigger for loading and unloading and and some other metadata about plugins

### Changed
   - Config struct updated to include plugins and a mutex
   - Plugin map type is now `map[string]NamedTrigger` instead of `sync.Map` and guarded with a `sync.RWMutexv`

## [1.1.0] - 2018-11-24
### Added
   - ACL type to manage allowed and rejected users
   - Admins list (using ACL type) and `--admins` flag
   - Dynamic loading and unloading of modules (guarded by admins list)
       - Modules that are loadable must be registered with the local triggers `sync.Map`
   - `!triggers` command to list registered modules
   - NamedTrigger type that extends the new interface used for bitbot handlers

### Changed
   - Running on fork of bitbot until work around new style of triggers/interfaces is resolved

## [1.0.0] - 2018-11-18
### Added
   - Skip processing of a message by prefixing your message with !skip
   - Support to logging into nickserv, server oper via flag configuration
   - 0.0
   - Config struct added

### Removed
   - Removed `version` subcommand in favor of a `--version` flag

### Fixed
   - No more panics when Bitbot can't read a web page.

### Changed
   - `Run` method's signature has changed and only accepts the newly-added Config struct
     - Yes, this is a breaking change (hence the version bump)
     - Accepting a config struct means fewer changes to the function signature in the future
   - Dockerfile is waaaay slimmer 
     - Runs in `scratch` container
     - Runs as limited user
     - Totally statically compiled

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
