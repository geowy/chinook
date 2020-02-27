# Training day

For training day we will be looking at writing web services in Go.

## AM

In the AM we will each add our own section to this example app.

Connect to the database (`sqlite3 chinook.db`) and pick one of the tables 
(`.tables`). You can inspect the tables with `.schema TABLE`.

The albums section is aready finished so we can use that as a pattern for the
for our own sections.

Main modules used are here:

* [database/sql](https://golang.org/pkg/database/sql/)
* [log](https://golang.org/pkg/log/)
* [net/http](https://golang.org/pkg/net/http/)

## PM

In the afternoon we can each spend some time looking into things that interest
us.

Suggestions:

* Sessions and authentication
* Serialising & deserialising JSON
* What web frameworks are available for Go?
