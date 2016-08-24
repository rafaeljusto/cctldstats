cctldstats
==========

Project developed to allow exporting public ccTLD statistics in a common format. The idea is to create a system that will query this service in all ccTLDs so we can store the number of domains history for further studies.

Protocol
--------

For now there's only one service defined to retrieve the number of registered domains. To retrieve it just send a HTTP `GET` to `/domains/registered`. The response will be something like:

```json
{
  "number": 400000
}
```

**Possible HTTP codes:** 200 (OK), 401 (Forbidden), 500 (Internal Server Error)

How to use it
-------------

First you will need to download and install the Go compiler. For that go to the http://golang.org/dl website and choose the version for your OS and architecture.

Now you will need to setup the Go project structure and set the GOPATH environment variable. Choose a place to store the Go projects, as example $HOME/go/src, and set the GOPATH for it (GOPATH=$HOME/go).

Build the project with following line:

    go get -u github.com/rafaeljusto/cctldstats

Now you can get it running, just set some environment variables that are requirements:

  Environment Variable         | Explanation
  ---------------------------- | ------------------------------------------
  CCTLDSTATS_DATABASE_KIND     | Database type: "mysql" or "postgres"
  CCTLDSTATS_DATABASE_NAME     | Database name
  CCTLDSTATS_DATABASE_USERNAME | Database username
  CCTLDSTATS_DATABASE_PASSWORD | Database password
  CCTLDSTATS_DATABASE_HOST     | Database host with port (127.0.0.1:3306)
  CCTLDSTATS_DOMAIN_TABLE_NAME | Name of the table that contain the domains
  CCTLDSTATS_ACL               | Allowed IP address separated by comma

You're good to go! The server runs on port 8888. Logs will be printed on the standard output.

Testing
-------

    curl -i http://localhost:8888/domains/registered
