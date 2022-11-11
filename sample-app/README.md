# Sample App

This is bad code.

We shell exec the start/stop/create/destroy docker commands (poorly)
and this is how the unit testing suite attempts to start/stop/create/destroy
the local persistent store.

(H3) Removed custom line by line shell execution and simply run shell scripts which 
will not be debugged line by line.  The problem with the line by line interpreter is that it 
was not a full bash shell, so short of extending simpler to remove. So the scripts below run by hand
are run in the same way be the automate tests (go test ./...)

### Running The Sample App

```bash 
./pcreate  # Will create the sample app running locally
./pdestroy # Will destroy the sample app, but the data will persist regardles of running this command
./pstop    # Will stop the photoprism app from running/serving
./plogs    # Will tail the photoprism logs
./pstart   # Will start an already created, and then stopped Photoprism application
```