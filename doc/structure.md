- cmd
  - ctftk: Main executable
- config
  - root_config.go: Root config parser
  - challenge_config.go: Challenge config parser
- providers
  - hosting: Challenge hosting providers
    - gcloud: Google Cloud
    - kubernetes: Raw kubernetes
  - scoreboard: Scoreboard providers
    - ctfscoreboard: support for google.com/ctfscoreboard

## TODO

- Sidecars?
- Non-challenge containers?
- Network Policy?
