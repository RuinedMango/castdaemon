package main
import (
	"net"
	"os"
	"log"
	"syscall"
	"os/signal"
)

func main() {
    // Create a Unix domain socket and listen for incoming connections.
    socket, err := net.Listen("unix", "/tmp/castdaemon.sock")
    if err != nil {
        log.Fatal(err)
    }

    // Cleanup the sockfile.
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        os.Remove("/tmp/castdaemon.sock")
        os.Exit(1)
    }()

    for {
        // Accept an incoming connection.
        conn, err := socket.Accept()
        if err != nil {
            log.Fatal(err)
        }

        // Handle the connection in a separate goroutine.
        go func(conn net.Conn) {
            defer conn.Close()
            // Create a buffer for incoming data.
            buf := make([]byte, 4096)

            // Read data from the connection.
            n, err := conn.Read(buf)
            if err != nil {
                log.Fatal(err)
            }
			handle(buf[:n])
        }(conn)
    }
}
