# DUCK Application

**_Attention: These are Instructions for Developers._** If you just want to try DUCK, you might want to download it [here](https://github.com/Microsoft/DUCK/releases).
A user manual can be found [here](docs/usermanual.md), the DUCK architecture is described [here](docs/architecture.md).

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.


This is a Gulp-powered build system with the following features:

- Sass compilation and prefixing
- JavaScript transpilation based on Babel and concatenation
- Go compilation
- Dynamic browser reloading using BrowserSync
- For production builds:
  - CSS compression
  - JavaScript compression
  - Image compression


## Installation

The project build requires:

- [Git](https://git-scm.com/)
- [Go](https://golang.org/)  (1.7 or later)
- [NodeJS](https://nodejs.org/en/) (0.12 or greater, LTS version >4.0.0 recommended)
- [CouchDB](http://couchdb.apache.org/) (1.6 or greater)


### Setup

First clone the project repository. Make sure it is under the GOPATH directory, for example if your GOPATH is:

```bash
/gocode
```
DUCK should be checked out to:

```bash
/gocode/src/github.com/Microsoft/DUCK
```
To clone the codebase, use: 
```bash
git clone https://github.com/Microsoft/DUCK DUCK
```

Then, from the cloned directory, install the required dependencies:

```bash
cd DUCK
npm install
npm install -g bower
npm install -g gulp
bower install
```


Make sure Couch DB is running.

Finally, run `npm start` to execute the build. The application will be accessible at:

```
http://localhost:8000
```
Dynamic reloading will be enabled. Both frontend (Javascript, CSS, HTML) and backend (go) assets are watched for changes, which will automatically trigger an
application update.  

To create compressed, production-ready assets, run `npm run build`.

### Building a Distrubution

Execute the distribution build using:

```
npm run distro
```

A binary archive will be generated in the /image directory

### Configuration

This project reads its configuration from the file `backend/configuration.json`, environment variables and also command-line flags. 
The following precendence order is used. Each item takes precedence over the item below it:


- flag
- env
- config
- default

#### Default
The default configuration has these values:

```yaml
  database: 
      location: "http://127.0.0.1"
      port: 5984
      name: "duck"
	  username: ""
	  password: ""
  jwtkey: "c2VjcmV0"
  webdir: "/src/github.com/Microsoft/DUCK/frontend/dist"
  rulebasedir: "/src/github.com/Microsoft/DUCK/RuleBases"
```
##### jwtkey
The field jwtkey is a base64 encoded string. If this field is empty, a random key will be generated.

##### regarding path variables

If rulebasedir or webdir have an absolute path it is used as an absolute path.
If it is a relative path it will be assumed to be relative to the GOPATH environment variable if present. 
If GOPATH is not found, the path is assumed to be relative to the go executable.

#### env
The environment variable names are prefixed with DUCK_ and all uppercase. Fields in the database object are referenced using the `.` operator, e.g. 
`DUCK_DATABASE.NAME`.

#### flags
The flags are handled in the go standard way described in https://golang.org/pkg/flag/. 
Main points are:

>Command line flag syntax:
>```
-flag
-flag=x
-flag x  // non-boolean flags only
```
>One or two minus signs may be used; they are equivalent. The last form is not permitted for boolean flags because the meaning of the command
>
>`cmd -x *`
>
>will change if there is a file called 0, false, etc. You must use the -flag=false form to turn off a boolean flag.
>
>Integer flags accept 1234, 0664, 0x1234 and may be negative. Boolean flags may be:
>
>`1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False`

It is *not* possible to configure the database connection via flags.
