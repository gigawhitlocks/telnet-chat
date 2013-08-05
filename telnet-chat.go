package main

import (
  "net"
  "fmt"
  "os"
  "bufio"
  "container/list"
)

var clients = list.New()

func main() {
  ln, err := net.Listen("tcp", ":6667")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Something went wrong before loop: %s", err)
  }
  for {
    conn, _ := ln.Accept()
    go handleConnection(conn)

  }
}

func handleConnection ( conn net.Conn ) {
  connp := clients.PushFront(conn)
  defer clients.Remove(connp)

  fmt.Println("connection handled")
  for {
    status, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Printf("%s",status)
    for c := clients.Front(); c != nil; c = c.Next() {
      if c != connp {
        writer := bufio.NewWriter((c.Value).(net.Conn))
        writer.WriteString(status)
        writer.Flush()
      }
    }
  }
}
