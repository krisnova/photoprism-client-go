# V1 SDK

This is the main SDK code. Here be dragons.

The V1 Client ships with the following HTTP methods

 - GET
 - POST
 - PUT
 - DELETE

The client can also be chained using the following two chain methods.

#### JSON Marshal

To send a GET request to `/api/v1/photos/:uuid` and marshal the results on to a Photo struct

```go 
    uuid := "123"
    photo := Photo{
        UUID: uuid,
    }
    err := v1.GET("/api/v1/photos/%s", uuid).JSON(&object)
    // 
    fmt.Println(err)
    fmt.Println(photo)
```

#### String

Sometimes it is helpful to just see what the Photoprism API returns.
The `String()` method implements the Go idiomatic `String()` and will
return the body of the response for debugging.

To send a GET request to `/api/v1/photos/:uuid` and see the raw JSON output

```go 
    uuid := "123"
    fmt.Println(v1.GET("/api/v1/photos/%s", uuid).String())
```
