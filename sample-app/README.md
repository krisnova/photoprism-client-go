# Sample App

This is bad code.

We shell exec the start/stop/create/destroy docker commands (poorly)
and this is how the unit testing suite attempts to start/stop/create/destroy
the local persistent store.

### Running The Sample App

```bash 
./pcreate  # Will create the sample app running locally
./pdestroy # Will destroy the sample app, but the data will persist regardles of running this command
./pstop    # Will stop the photoprism app from running/serving
./plogs    # Will tail the photoprism logs
./pstart   # Will start an already created, and then stopped Photoprism application
```