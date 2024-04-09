# A go package to make http requests easier

## Usage

1. Install package
   ```shell
    go get github.com/teepee78/requests
    ```
2. Import the package
   ```go
   package main
   
   import "github.com/teepee78/requests"
    ```
3. Create a new `Request` object
   ```go
    baseUrl := "http://localhost:5000"
    headers := make(map[string]string)
    headers["Content-Type"] = "application/json"
   
    request := requests.CreateRequest(baseUrl, headers)
    ```
4. Request away
   ```go
      response := request.Get("/users", nil)
   ```