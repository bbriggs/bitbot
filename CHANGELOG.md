# Change Log

## [Unreleased](https://github.com/bbriggs/bitbot/tree/HEAD)

[Full Changelog](https://github.com/bbriggs/bitbot/compare/v1.3.1...HEAD)

**Implemented enhancements:**

- \[Reminder trigger\] Indicate timezone. [\#180](https://github.com/bbriggs/bitbot/issues/180)

**Fixed bugs:**

- \\[Reminder trigger\\] Indicate timezone. [\#180](https://github.com/bbriggs/bitbot/issues/180)

**Merged pull requests:**

- Release 1.3.0 [\#185](https://github.com/bbriggs/bitbot/pull/185) ([bbriggs](https://github.com/bbriggs))

## [v1.3.1](https://github.com/bbriggs/bitbot/tree/v1.3.1) (2020-06-19)
[Full Changelog](https://github.com/bbriggs/bitbot/compare/v1.3.0...v1.3.1)

**Implemented enhancements:**

- Add urban dictionary support [\#14](https://github.com/bbriggs/bitbot/issues/14)

**Fixed bugs:**

- Unable to setup bitbot on local system [\#161](https://github.com/bbriggs/bitbot/issues/161)

**Closed issues:**

- contributing.md [\#172](https://github.com/bbriggs/bitbot/issues/172)
- Update or Remove DeepSource configuration file. [\#169](https://github.com/bbriggs/bitbot/issues/169)
- Event Setting Trigger [\#162](https://github.com/bbriggs/bitbot/issues/162)
- Missing error checks \(and a couple more\) [\#157](https://github.com/bbriggs/bitbot/issues/157)
- Update Readme [\#147](https://github.com/bbriggs/bitbot/issues/147)
- Adapt golangci-lint workflow or refactor to comply with it [\#145](https://github.com/bbriggs/bitbot/issues/145)
- bitbot leaves channel upon mentioning it [\#139](https://github.com/bbriggs/bitbot/issues/139)
- Feature Request - Add the PUMP command [\#136](https://github.com/bbriggs/bitbot/issues/136)
- Bitbot does not properly identify with nickserv on Freenode [\#100](https://github.com/bbriggs/bitbot/issues/100)
- `markovInit` reporting failed even though it succeeds [\#94](https://github.com/bbriggs/bitbot/issues/94)
- Bitbot should follow and expand shortened URLs [\#29](https://github.com/bbriggs/bitbot/issues/29)
- !idle crashes when user doesn't have an entry/key in idlers DB [\#13](https://github.com/bbriggs/bitbot/issues/13)

**Merged pull requests:**

- Added the UTC to all the places. :P [\#184](https://github.com/bbriggs/bitbot/pull/184) ([PJ290](https://github.com/PJ290))
- Add docker-compose [\#183](https://github.com/bbriggs/bitbot/pull/183) ([bbriggs](https://github.com/bbriggs))
- Fix broken bot initialization [\#181](https://github.com/bbriggs/bitbot/pull/181) ([bbriggs](https://github.com/bbriggs))
- Improve logging at startup and make DB connection failure non-fatal [\#179](https://github.com/bbriggs/bitbot/pull/179) ([bbriggs](https://github.com/bbriggs))
- Fix triming bitbot [\#177](https://github.com/bbriggs/bitbot/pull/177) ([m-242](https://github.com/m-242))
- Fix triming bitbot [\#176](https://github.com/bbriggs/bitbot/pull/176) ([m-242](https://github.com/m-242))
- Doc/contributing.md [\#174](https://github.com/bbriggs/bitbot/pull/174) ([m-242](https://github.com/m-242))
- Plugins/testing [\#173](https://github.com/bbriggs/bitbot/pull/173) ([m-242](https://github.com/m-242))
- Remove useless deepsource, closes \#169 [\#171](https://github.com/bbriggs/bitbot/pull/171) ([m-242](https://github.com/m-242))
- Update channels.go [\#170](https://github.com/bbriggs/bitbot/pull/170) ([PJ290](https://github.com/PJ290))
- Feature/reminder, closes \#162 [\#163](https://github.com/bbriggs/bitbot/pull/163) ([m-242](https://github.com/m-242))
- Fixed improper definition sanitizing [\#160](https://github.com/bbriggs/bitbot/pull/160) ([m-242](https://github.com/m-242))
- Nolint false positives, closes \#145 [\#159](https://github.com/bbriggs/bitbot/pull/159) ([m-242](https://github.com/m-242))
- removing code anti-patterns [\#158](https://github.com/bbriggs/bitbot/pull/158) ([staticnotdynamic](https://github.com/staticnotdynamic))
- Added urbd trigger, closes \#14 [\#156](https://github.com/bbriggs/bitbot/pull/156) ([m-242](https://github.com/m-242))
- Covid province [\#155](https://github.com/bbriggs/bitbot/pull/155) ([m-242](https://github.com/m-242))
- Fixed total trigger addition loops [\#154](https://github.com/bbriggs/bitbot/pull/154) ([m-242](https://github.com/m-242))
- Get total for country by summing all provinces [\#152](https://github.com/bbriggs/bitbot/pull/152) ([bbriggs](https://github.com/bbriggs))
- Uppercase arguments to covid19 stats function [\#151](https://github.com/bbriggs/bitbot/pull/151) ([bbriggs](https://github.com/bbriggs))
- Updated README.md, closes \#147 [\#150](https://github.com/bbriggs/bitbot/pull/150) ([m-242](https://github.com/m-242))
- Feature/covid data [\#148](https://github.com/bbriggs/bitbot/pull/148) ([rfc2119](https://github.com/rfc2119))
- Fix golangci lint errors [\#146](https://github.com/bbriggs/bitbot/pull/146) ([m-242](https://github.com/m-242))
- bitbot/channels.go: fixed bitbot erronous parting [\#143](https://github.com/bbriggs/bitbot/pull/143) ([rfc2119](https://github.com/rfc2119))
- Add test action [\#142](https://github.com/bbriggs/bitbot/pull/142) ([bbriggs](https://github.com/bbriggs))

## [v1.3.0](https://github.com/bbriggs/bitbot/tree/v1.3.0) (2020-03-23)
[Full Changelog](https://github.com/bbriggs/bitbot/compare/v1.2.2...v1.3.0)

**Merged pull requests:**

- Release 1.3.0 [\#141](https://github.com/bbriggs/bitbot/pull/141) ([bbriggs](https://github.com/bbriggs))
- Migrate from bbolt to Postgres [\#140](https://github.com/bbriggs/bitbot/pull/140) ([bbriggs](https://github.com/bbriggs))
- Declare /data volume in Dockerfile [\#138](https://github.com/bbriggs/bitbot/pull/138) ([bbriggs](https://github.com/bbriggs))
- Minor vanity fix [\#132](https://github.com/bbriggs/bitbot/pull/132) ([PaulGWebster](https://github.com/PaulGWebster))
- Add CODEOWNERS file [\#131](https://github.com/bbriggs/bitbot/pull/131) ([bbriggs](https://github.com/bbriggs))
- undid m242s shinnanigans [\#130](https://github.com/bbriggs/bitbot/pull/130) ([TheRealSuser](https://github.com/TheRealSuser))
- Release 1.2.2 [\#128](https://github.com/bbriggs/bitbot/pull/128) ([bbriggs](https://github.com/bbriggs))

## [v1.2.2](https://github.com/bbriggs/bitbot/tree/v1.2.2) (2020-01-25)
[Full Changelog](https://github.com/bbriggs/bitbot/compare/v1.2.1...v1.2.2)

**Implemented enhancements:**

- Improve ACL checks [\#127](https://github.com/bbriggs/bitbot/pull/127) ([TheRealSuser](https://github.com/TheRealSuser))
- Ipinfo add [\#125](https://github.com/bbriggs/bitbot/pull/125) ([TheRealSuser](https://github.com/TheRealSuser))

**Fixed bugs:**

- suser and m242 resume their epeen contest [\#124](https://github.com/bbriggs/bitbot/pull/124) ([TheRealSuser](https://github.com/TheRealSuser))

**Merged pull requests:**

- Add ipinfo module [\#126](https://github.com/bbriggs/bitbot/pull/126) ([TheRealSuser](https://github.com/TheRealSuser))
- Export name for RaiderQuoteTrigger and add it to default plugin map [\#108](https://github.com/bbriggs/bitbot/pull/108) ([bbriggs](https://github.com/bbriggs))

## [v1.2.1](https://github.com/bbriggs/bitbot/tree/v1.2.1) (2019-12-19)
[Full Changelog](https://github.com/bbriggs/bitbot/compare/v1.2.0-CunningChuckwalla...v1.2.1)

**Fixed bugs:**

- epeen fixed formatting maintainted [\#122](https://github.com/bbriggs/bitbot/pull/122) ([TheRealSuser](https://github.com/TheRealSuser))

**Merged pull requests:**

- Dev [\#123](https://github.com/bbriggs/bitbot/pull/123) ([m-242](https://github.com/m-242))
- Update changelog with new release notes [\#121](https://github.com/bbriggs/bitbot/pull/121) ([bbriggs](https://github.com/bbriggs))

## [v1.2.0-CunningChuckwalla](https://github.com/bbriggs/bitbot/tree/v1.2.0-CunningChuckwalla) (2019-12-17)
[Full Changelog](https://github.com/bbriggs/bitbot/compare/v1.1.0...v1.2.0-CunningChuckwalla)

**Implemented enhancements:**

- Include !tableflip and !unflip commands [\#42](https://github.com/bbriggs/bitbot/issues/42)
- created epeen trigger [\#116](https://github.com/bbriggs/bitbot/pull/116) ([TheRealSuser](https://github.com/TheRealSuser))
- Updated markov help [\#114](https://github.com/bbriggs/bitbot/pull/114) ([m-242](https://github.com/m-242))
- Use new Part method in hbot.Bot [\#112](https://github.com/bbriggs/bitbot/pull/112) ([bbriggs](https://github.com/bbriggs))
- Join and part channels [\#110](https://github.com/bbriggs/bitbot/pull/110) ([bbriggs](https://github.com/bbriggs))
- Prevent MarkovTrainer from training on links and commands [\#107](https://github.com/bbriggs/bitbot/pull/107) ([bbriggs](https://github.com/bbriggs))
- PM tarot reading when \> 5 cards [\#99](https://github.com/bbriggs/bitbot/pull/99) ([skidd0](https://github.com/skidd0))
- Readded snwcrsh source, and converted the array to a slice. \(\#94\) [\#97](https://github.com/bbriggs/bitbot/pull/97) ([parsec](https://github.com/parsec))
- Limit !babble to max 200 chars. [\#96](https://github.com/bbriggs/bitbot/pull/96) ([skidd0](https://github.com/skidd0))

**Fixed bugs:**

- Limit babble size to avoid kick [\#93](https://github.com/bbriggs/bitbot/issues/93)
- Fix docker-build.sh [\#91](https://github.com/bbriggs/bitbot/issues/91)
- Channel Pop Gauge is Registering High Counts [\#76](https://github.com/bbriggs/bitbot/issues/76)
- fmt-ed epeen [\#118](https://github.com/bbriggs/bitbot/pull/118) ([skidd0](https://github.com/skidd0))
- Fix epeen [\#117](https://github.com/bbriggs/bitbot/pull/117) ([skidd0](https://github.com/skidd0))
- Properly set version/commit/branch info at build time [\#113](https://github.com/bbriggs/bitbot/pull/113) ([bbriggs](https://github.com/bbriggs))
- fixed the segfault problem by cleaning the url, closes \#104 [\#106](https://github.com/bbriggs/bitbot/pull/106) ([m-242](https://github.com/m-242))
- Fix improperly bound nickserv config variable [\#102](https://github.com/bbriggs/bitbot/pull/102) ([bbriggs](https://github.com/bbriggs))
- Update nickserv login to include bot nick [\#101](https://github.com/bbriggs/bitbot/pull/101) ([bbriggs](https://github.com/bbriggs))
- Fix docker build.sh, closes \#91 [\#92](https://github.com/bbriggs/bitbot/pull/92) ([m-242](https://github.com/m-242))

**Closed issues:**

- Links shortening segmentation fault [\#104](https://github.com/bbriggs/bitbot/issues/104)
- Fix `version.go` [\#98](https://github.com/bbriggs/bitbot/issues/98)
- Use the log package instead of the fmt packages inside triggers [\#88](https://github.com/bbriggs/bitbot/issues/88)
- Automate the Markov triggers seeding. [\#82](https://github.com/bbriggs/bitbot/issues/82)
- Add controls to re-initialize \(clear\) and bootstrap markov model [\#81](https://github.com/bbriggs/bitbot/issues/81)
- Update references in `go.mod` file [\#78](https://github.com/bbriggs/bitbot/issues/78)
- Anti spam module [\#61](https://github.com/bbriggs/bitbot/issues/61)
- Get a new nick if configured one is taken [\#55](https://github.com/bbriggs/bitbot/issues/55)
- Magic 8-Ball Module [\#46](https://github.com/bbriggs/bitbot/issues/46)
- Troll Launcher Module [\#45](https://github.com/bbriggs/bitbot/issues/45)
- Truncated links should end with ellipses. [\#43](https://github.com/bbriggs/bitbot/issues/43)
- Build/incorporate a URL shortener [\#28](https://github.com/bbriggs/bitbot/issues/28)
- Nickserv authentication [\#26](https://github.com/bbriggs/bitbot/issues/26)
- Add RPG style dice roll function [\#18](https://github.com/bbriggs/bitbot/issues/18)

**Merged pull requests:**

- Prep for release 1.2.0 [\#120](https://github.com/bbriggs/bitbot/pull/120) ([bbriggs](https://github.com/bbriggs))
- Use Changelog Generator to standardize changelog [\#119](https://github.com/bbriggs/bitbot/pull/119) ([bbriggs](https://github.com/bbriggs))
- reduce markov chance [\#90](https://github.com/bbriggs/bitbot/pull/90) ([skidd0](https://github.com/skidd0))
- used log in triggers, closes \#88 [\#89](https://github.com/bbriggs/bitbot/pull/89) ([m-242](https://github.com/m-242))
- Shortening long urls [\#87](https://github.com/bbriggs/bitbot/pull/87) ([m-242](https://github.com/m-242))
- Removed Snwcrsh Source [\#86](https://github.com/bbriggs/bitbot/pull/86) ([parsec](https://github.com/parsec))
- Wrote `!troll` Troll Launcher \(\#45\) [\#85](https://github.com/bbriggs/bitbot/pull/85) ([parsec](https://github.com/parsec))
- Updated `cmd/root.go` to include `markovInit` [\#84](https://github.com/bbriggs/bitbot/pull/84) ([parsec](https://github.com/parsec))
- MarkovInit \(Issue \#81\) [\#83](https://github.com/bbriggs/bitbot/pull/83) ([parsec](https://github.com/parsec))
- Add markov trigger [\#80](https://github.com/bbriggs/bitbot/pull/80) ([bbriggs](https://github.com/bbriggs))
- Add new trigger to watch for RPL\_LIST messages from server. Fixes \#76 [\#77](https://github.com/bbriggs/bitbot/pull/77) ([bbriggs](https://github.com/bbriggs))
- Add Prom gauge to track user count per channel [\#75](https://github.com/bbriggs/bitbot/pull/75) ([bbriggs](https://github.com/bbriggs))
- Add 8ball to trigger condition tests [\#74](https://github.com/bbriggs/bitbot/pull/74) ([bbriggs](https://github.com/bbriggs))
- Add Tarot Trigger [\#73](https://github.com/bbriggs/bitbot/pull/73) ([skidd0](https://github.com/skidd0))
- Add magic 8 ball [\#72](https://github.com/bbriggs/bitbot/pull/72) ([bbriggs](https://github.com/bbriggs))
- Help text [\#71](https://github.com/bbriggs/bitbot/pull/71) ([bbriggs](https://github.com/bbriggs))
- Prometheus [\#70](https://github.com/bbriggs/bitbot/pull/70) ([bbriggs](https://github.com/bbriggs))
- Add a help plugin [\#69](https://github.com/bbriggs/bitbot/pull/69) ([bbriggs](https://github.com/bbriggs))
- Flipping the table [\#68](https://github.com/bbriggs/bitbot/pull/68) ([m-242](https://github.com/m-242))
- Add BeefyTrigger to default trigger map [\#64](https://github.com/bbriggs/bitbot/pull/64) ([bbriggs](https://github.com/bbriggs))
- Enable all plugins in pluginMap by default [\#63](https://github.com/bbriggs/bitbot/pull/63) ([bbriggs](https://github.com/bbriggs))
- no push binary [\#62](https://github.com/bbriggs/bitbot/pull/62) ([m-242](https://github.com/m-242))
- Auto rename [\#60](https://github.com/bbriggs/bitbot/pull/60) ([m-242](https://github.com/m-242))
- Tests for Trigger Conditions [\#59](https://github.com/bbriggs/bitbot/pull/59) ([bbriggs](https://github.com/bbriggs))
- Create map of available plugins to be loaded [\#58](https://github.com/bbriggs/bitbot/pull/58) ([bbriggs](https://github.com/bbriggs))
- BEEF [\#57](https://github.com/bbriggs/bitbot/pull/57) ([bbriggs](https://github.com/bbriggs))
- Update README to new usage patterns [\#56](https://github.com/bbriggs/bitbot/pull/56) ([bbriggs](https://github.com/bbriggs))
- Title shortener working, closes \#43 [\#53](https://github.com/bbriggs/bitbot/pull/53) ([m-242](https://github.com/m-242))
- Fixing circle ci tests [\#52](https://github.com/bbriggs/bitbot/pull/52) ([m-242](https://github.com/m-242))
- Pass plugins in as configuration [\#51](https://github.com/bbriggs/bitbot/pull/51) ([bbriggs](https://github.com/bbriggs))
- Fix wrong spelling in docker-build.sh [\#50](https://github.com/bbriggs/bitbot/pull/50) ([omBratteng](https://github.com/omBratteng))
- Link expansion [\#49](https://github.com/bbriggs/bitbot/pull/49) ([bbriggs](https://github.com/bbriggs))
- Send first 350 chars of title to account for twitter links [\#48](https://github.com/bbriggs/bitbot/pull/48) ([bbriggs](https://github.com/bbriggs))
- Stop hauling deps around like a moron [\#47](https://github.com/bbriggs/bitbot/pull/47) ([bbriggs](https://github.com/bbriggs))
- Finally patched roll.go [\#44](https://github.com/bbriggs/bitbot/pull/44) ([parsec](https://github.com/parsec))
- Fix crashes on empty titles; extend search space for titles to 2^16-1… [\#41](https://github.com/bbriggs/bitbot/pull/41) ([bbriggs](https://github.com/bbriggs))
- Add quotes module [\#40](https://github.com/bbriggs/bitbot/pull/40) ([bbriggs](https://github.com/bbriggs))
- Decisions [\#38](https://github.com/bbriggs/bitbot/pull/38) ([bbriggs](https://github.com/bbriggs))
- Redirect following [\#30](https://github.com/bbriggs/bitbot/pull/30) ([harrywhite4](https://github.com/harrywhite4))

## [v1.1.0](https://github.com/bbriggs/bitbot/tree/v1.1.0) (2018-11-24)
[Full Changelog](https://github.com/bbriggs/bitbot/compare/v1.0.0...v1.1.0)

**Merged pull requests:**

- Hot loading [\#37](https://github.com/bbriggs/bitbot/pull/37) ([bbriggs](https://github.com/bbriggs))

## [v1.0.0](https://github.com/bbriggs/bitbot/tree/v1.0.0) (2018-11-19)
[Full Changelog](https://github.com/bbriggs/bitbot/compare/0.1.1...v1.0.0)

**Fixed bugs:**

- Crashing on short HTTP title elements [\#23](https://github.com/bbriggs/bitbot/issues/23)

**Merged pull requests:**

- Add support for nickserv, server oper, and channel oper [\#36](https://github.com/bbriggs/bitbot/pull/36) ([bbriggs](https://github.com/bbriggs))
- Slim down image size and use scratch container [\#35](https://github.com/bbriggs/bitbot/pull/35) ([bbriggs](https://github.com/bbriggs))
- Set abyss threshold to 2 [\#34](https://github.com/bbriggs/bitbot/pull/34) ([bbriggs](https://github.com/bbriggs))
- Add Abyss simulator [\#33](https://github.com/bbriggs/bitbot/pull/33) ([bbriggs](https://github.com/bbriggs))
- fix panic [\#32](https://github.com/bbriggs/bitbot/pull/32) ([C-Sto](https://github.com/C-Sto))
- Add skip flag [\#31](https://github.com/bbriggs/bitbot/pull/31) ([bbriggs](https://github.com/bbriggs))
- Fix short title problem [\#25](https://github.com/bbriggs/bitbot/pull/25) ([sylviamoss](https://github.com/sylviamoss))

## [0.1.1](https://github.com/bbriggs/bitbot/tree/0.1.1) (2018-10-08)
[Full Changelog](https://github.com/bbriggs/bitbot/compare/0.1.0...0.1.1)

**Fixed bugs:**

- Bitbot crashes on very large titles [\#20](https://github.com/bbriggs/bitbot/issues/20)

**Closed issues:**

- !shrug breaks on a trailing space [\#16](https://github.com/bbriggs/bitbot/issues/16)

**Merged pull requests:**

- Release 0.1.1 [\#24](https://github.com/bbriggs/bitbot/pull/24) ([bbriggs](https://github.com/bbriggs))
- fix crashes on very long titles [\#22](https://github.com/bbriggs/bitbot/pull/22) ([jrwren](https://github.com/jrwren))
- Trim spaces from incoming messages when matching for single-word comm… [\#17](https://github.com/bbriggs/bitbot/pull/17) ([bbriggs](https://github.com/bbriggs))
- Add shrug trigger to bitbot [\#15](https://github.com/bbriggs/bitbot/pull/15) ([wadadli](https://github.com/wadadli))
- Dockerhub sucks [\#12](https://github.com/bbriggs/bitbot/pull/12) ([bbriggs](https://github.com/bbriggs))
- Work around Dockerhub's shallow clone and set a default version [\#11](https://github.com/bbriggs/bitbot/pull/11) ([bbriggs](https://github.com/bbriggs))
- Add compile output information to aid in debugging [\#10](https://github.com/bbriggs/bitbot/pull/10) ([bbriggs](https://github.com/bbriggs))
- fix configuration issues with circle ci [\#9](https://github.com/bbriggs/bitbot/pull/9) ([wadadli](https://github.com/wadadli))
- updates to the circle ci config [\#8](https://github.com/bbriggs/bitbot/pull/8) ([wadadli](https://github.com/wadadli))
- Update info output [\#7](https://github.com/bbriggs/bitbot/pull/7) ([bbriggs](https://github.com/bbriggs))
- Add version output at build time  [\#6](https://github.com/bbriggs/bitbot/pull/6) ([bbriggs](https://github.com/bbriggs))

## [0.1.0](https://github.com/bbriggs/bitbot/tree/0.1.0) (2018-09-28)
**Closed issues:**

- Crashes when getting unsupported URL schemes [\#2](https://github.com/bbriggs/bitbot/issues/2)

**Merged pull requests:**

- Switch back to upstream [\#5](https://github.com/bbriggs/bitbot/pull/5) ([bbriggs](https://github.com/bbriggs))
- Remove the go get ./.. from Dockerfile  [\#4](https://github.com/bbriggs/bitbot/pull/4) ([bbriggs](https://github.com/bbriggs))
- Fix \#2: Move defer after error handling in url lookups [\#3](https://github.com/bbriggs/bitbot/pull/3) ([bbriggs](https://github.com/bbriggs))
- Fix Spelling Error [\#1](https://github.com/bbriggs/bitbot/pull/1) ([zombeej](https://github.com/zombeej))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*
