# Asana Priority Projects List

Assuming you use Asana on a large enough team, you may find yourself tracking
so many projects that it becomes difficult to find signal in the noise.

At UP Global, we use Asana quite heavily and have found ways to use
Asana projects to communicate at various levels of granularity for a given
audience.

The Asana projects list isn't aware of this hierarchy so we must find
a better way.

This project makes use of the Asana API and a standard format to

- communicate top-tier projects in your workspace in a clear way
- communicate the projects status

We rely on the following project naming format to determine status:

`[STATUS] YOUR PROJECT NAME`

The `STATUS` value should be either "R", "Y", "G" for "red, yellow, and green",
respectively.

Once you do that and spin up this server, we'll take care of the rest.

To run, execute `go run main.go`. Otherwise, you can use the supplied
`Procfile` to run this on Heroku.

# Technical Requirements and Options

- `asanakey` - Supply this value in `config/production.json`, which is a file
must create. Take a look at `config/defaults.json` to see which settings the
project expects
- `port` - We default to `5000` for testing, but you can set this value with
`PORT` environment variable
- `appmode` - Tells the server whether to run in `development`, `test`, or
`production` mode
