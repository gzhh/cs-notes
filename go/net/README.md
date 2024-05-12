# package net
- https://pkg.go.dev/net

## 一、包级别对外方法

### net.Listen()

```go
func Listen(network, address string) (Listener, error) {
   var lc ListenConfig
   return lc.Listen(context.Background(), network, address)
}
```

### net.Dial()

```go
func (d *Dialer) Dial(network, address string) (Conn, error) {
   return d.DialContext(context.Background(), network, address)
}
```

### net.Pipe()

## 二、包内结构

### ListenConfig struct

```go
func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error) {

}
```

### Listener interface

```go
// A Listener is a generic network listener for stream-oriented protocols.
//
// Multiple goroutines may invoke methods on a Listener simultaneously.
type Listener interface {
   // Accept waits for and returns the next connection to the listener.
   Accept() (Conn, error)

   // Close closes the listener.
   // Any blocked Accept operations will be unblocked and return errors.
   Close() error

   // Addr returns the listener's network address.
   Addr() Addr
}
```

### Conn interface

```go
// Conn is a generic stream-oriented network connection.
//
// Multiple goroutines may invoke methods on a Conn simultaneously.
type Conn interface {
   // Read reads data from the connection.
   // Read can be made to time out and return an Error with Timeout() == true
   // after a fixed time limit; see SetDeadline and SetReadDeadline.
   Read(b []byte) (n int, err error)

   // Write writes data to the connection.
   // Write can be made to time out and return an Error with Timeout() == true
   // after a fixed time limit; see SetDeadline and SetWriteDeadline.
   Write(b []byte) (n int, err error)

   // Close closes the connection.
   // Any blocked Read or Write operations will be unblocked and return errors.
   Close() error

   ...
}
```

```go
// Implementation of the Conn interface.

// Read implements the Conn Read method.
func (c *conn) Read(b []byte) (int, error) {
   if !c.ok() {
      return 0, syscall.EINVAL
}
   n, err := c.fd.Read(b)
   if err != nil && err != io.EOF {
      err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
   }
   return n, err
}

// Write implements the Conn Write method.
func (c *conn) Write(b []byte) (int, error) {
   if !c.ok() {
      return 0, syscall.EINVAL
}
   n, err := c.fd.Write(b)
   if err != nil {
      err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
   }
   return n, err
}

// Close closes the connection.
func (c *conn) Close() error {
   if !c.ok() {
      return syscall.EINVAL
}
   err := c.fd.Close()
   if err != nil {
      err = &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
   }
   return err
}
```

### Ref

- https://zhuanlan.zhihu.com/p/31644462

- 基于net/http框架个部分总结 https://learnku.com/docs/eudore/32-frame-of-golang-http/9669

- Golang net 包学习和实战 https://segmentfault.com/a/1190000022577103


### 优秀开源项目
- Centrifugo is an open-source scalable real-time messaging server. https://github.com/centrifugal/centrifugo
