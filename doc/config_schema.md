## Top-Level Config

* `version`: A string version that ensures the repository is compatible with
  this version of the tooling.
* `hosting`: A dictionary of options for the provider.
  * `name`: The provider for the hosting.  See the providers list for
    supported providers and per-provider options.
* `scoreboard`: A dictionary of options for the scoreboard provider.
  * `name`: The name of the scoreboard provider.  The only supported option is
    `ctfscoreboard` as in `github.com/google/ctfscoreboard`
  * `url`: The base URL of the scoreboard.
* `challenge_domain`: Base domain for challenges.
* `hostname_secret`: Secret for generating hashed hostnames.
* `default_replicas`: Default replicas for each challenge.
* `max_replicas`: Maximum number of replicas.
* `tags`: Array/list of tags for challenge tagging.
  * `name`: Name of tag as displayed.
  * `description`: Description of tag.
  * `aliases`: Optional array of tag strings to recognize as aliases.

## Provider Config (gcloud)

* `project_id`: GCP Project ID
* `region`: Region for regional commands
* `zone`: Zone for zonal commands
* `service_account`: Service Account for K8S instances
* `challenge_dns_zone`: Name of DNS zone instance for DNS management.

## Per-Challenge Config

* `version`: A string version for the challenge.  Update this to a new string
  to trigger builds/updates.
* `name`: Name of the challenge, as displayed on the scoreboard.
* `type`: Type of the challenge, currently supported:
  * `offline`: No hosted components, just upload to scoreboard.
  * `tcp`: Runs a TCP socket.
  * `http`: Runs a web service behind an HTTP/HTTPS aware load balancer.
* `author`: Name of the challenge author.
* `description`: Description of challenge for player consumption.
* `points`: Point value of challenge.
* `flag`: A string flag using the default validator.  This is shorthand for a
  single instance of `flags` below.
* `flags`: An array of objects to support multiple flag challenges.  If this is
  specified, then `points` and `flag` are ignored at the top level. Each
  object's keys are:
  * `name`: Challenge name associated with the flag (to distringuish on
    scoreboard)
  * `flag`: Flag value for this flag.
  * `points`: Points for this flag.
  * `validator`: Which validator to use (for scoreboards supporting it)
* `notes`: Internal notes for challenge admins.
* `container`: Options for the container.
  * `build_command`: Specify a command to be run instead of docker build.  Must
    accept the tag name as an argument.
  * `prebuild_command`: A command to be run before building.  The working
    directory will be set to the root of the challenge.
* `deployment`:
  * `setuid`: If true, will ensure the deployment allows setuid binaries to
    work.
  * `ptrace`: If true, will ensure the deployment allows ptrace to work.
  * `port`: Network port the container listens on.
