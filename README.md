## 📝 Go SSL Certificate Generator 📝

A utility package for generating self-signed SSL certificates within your Go application. No third-party tools required.
Self-signed certs will be generated in the `"etc/ssl/"` directory of your application. As of now, this path is not changeable.
I will likely change this in the future to be configurable but defaulting to `"etc/ssl"` to ensure backwards compatibility.

```
[ your-awesome-go-app ]
 |_ cmd
 |_ pkg
 |_ internal
 |_ etc
 |   |_ssl
 |  	|_ server.key   <- Generated key
 |	|_ server.pem   <- Generated pem
 |
 |_ README.md
```
____
### Install Instructions:
> ``go install https://github.com/KeithAlt/go-cert-generator``
____
### Example Implementation:
```go
// An example of package usage with Gin ...

func main() {
	cert, err := gencert.Generate()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.RunTLS(":8080", cert.PemPath, cert.KeyPath)
}
```
___
### Accessible Fields:
```go
type Cert struct {
	PemPath   string `binding:"required"`
	KeyPath   string `binding:"required"`
	PemBytes  []byte `binding:"required"`
	KeyBytes  []byte `binding:"required"`
	CertBytes []byte `binding:"required"`
}
```

```go
  cert, err := gencert.Generate()
  if err != nil {
    log.Fatal(err)
  }
  
  fmt.Println(cert.PemPath) // -> String path to the generated .pem file
  fmt.Println(cert.KeyPath) // -> String path to the generated .key file
  fmt.Println(cert.PemBytes) // -> []Byte field containing the pem file contents
  fmt.Println(cert.KeyBytes) // -> []Byte field containing the key file contents
  fmt.Println(cert.CertBytes) // -> []Byte field containing the cert contents
```
___

If you are checking out this package for usage in an internally communicating API, consider [GRPC](https://github.com/grpc/grpc)
