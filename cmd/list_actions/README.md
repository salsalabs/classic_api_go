# Fetch actions and convert to PDF

This app creates a bash shell that causes an application to read the
contents of actions and then convert them to PDF.  The application is
a NodeJS app that leverates the `puppeteer` library.  That's the only
thing that I've found that does a nice, neat PDF with out a lot of
empty pages at the top.

# Installation

The app expects that `hn.js` is in the current directory and that the current
directory is initialized properly to be both a Go and a NodeJS directory. Use
these steps to make that happen.

### Go
1. Latest version of [Go](https://golang.org}.
1. Correct directory structure.
```
HOME
  |
  + go
    |
    + bin
    + pkg
    + src
      |
      + github.com
        |
        + salsalabs
```
1. Change the directory to `salsalabs`.
1. Install this repository.
```bash
git clone https://github.com/salsalabs/godig
```
1. Change the directory to `godig`.
1. Install all Go dependencies.
```bash
go get ./...
```
1. Build this application
```bash
go build -o list_events cmd/list_events
```

### NodeJS

Execute these statements in the `godig` directory.  This is important.

1. Install [Node Version Manager (NVM)](https://github.com/nvm-sh/nvm)
1. Install the latest version of nvm in the current directory.
```bash
nvm install default
```
1. Initialize to be a NodeJS package.
```bash
npm init #then tap the enter key a bunch of times
```
1. Install the `puppeteer` library for NoeJS](https://github.com/GoogleChrome/puppeteer)
```bash
npm install --save puppeteer
```

### Salsa
1. Campaign manager credentials for the organization.
1. Create a YAML file with the campaign manager credentials.  We'll refer to this as `whatever.yaml` in this document.
1. [click here](https://help.salsalabs.com/hc/en-us/articles/115000341773-Salsa-Classic-API#api_host) to find the hostname.
```yaml
hostname: Salsa node name
email:  campaign manager's email address
password: campaign manager's password
```

### Execution

1. Run the Go app to create a bash script to retrieve actions.
```bash
./list_events --login whatever.yaml
```
1. Run the script generated by the Go command.
```bash
bash fetch_actions.bash
```
1. There may be errors. You'll have to fix those yourself.
1. When the process is done, there will be a directory named `pdfs/actions`. It contains all of the action PDFs.

# Questions?
Use the "Issues" link at the top of this repository to ask questions and report
errors.  **Do not contact Salsalabs support**.  They are surly and bite-y during
their nesting season.  Use the "Issues" link.  We'll be back in touch eventually.