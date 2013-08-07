package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"net"
	"os"
	//"unicode/utf8"
)

var bgcolors = map[string]string{
	"red":     "\033[41m",
	"yellow":  "\033[43m",
	"green":   "\033[42m",
	"blue":    "\033[46m",
	"indigo":  "\033[44m",
	"violet":  "\033[45m",
	"black":   "\033[40m",
	"default": "\033[49m",
}

var fgcolors = map[string]string{
	"black":   "\033[30m",
	"white":   "\033[37m",
	"red":     "\033[31m",
	"yellow":  "\033[33m",
	"green":   "\033[32m",
	"blue":    "\033[36m",
	"indigo":  "\033[34m",
	"violet":  "\033[35m",
	"default": "\033[39m",
}

var clients = list.New()

func main() {
	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong before loop: %s", err)
	}
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	connp := clients.PushFront(conn)
	defer func() {
		clients.Remove(connp)
	}()

  for {
    fmt.Println("connection handled")
    for c := clients.Front(); c != nil; c = c.Next() {
      if c != connp {
        paintRainbow((c.Value).(net.Conn))

      }
    }
  }
}

func paintRainbow(client net.Conn) {
	var buf bytes.Buffer
	for _, color := range bgcolors {
		buf.Write([]byte(color))
		for i := 0; i < 80; i++ {
			buf.Write([]byte(" "))
		}
		buf.Write([]byte(bgcolors["default"]))
		buf.Write([]byte("\n"))
	}
	writer := bufio.NewWriter(client)
	writer.WriteString(buf.String())
	writer.Flush()

}

/*
func handleConnection ( conn net.Conn ) {
  connp := clients.PushFront(conn)
  defer func () {
    clients.Remove(connp)
  }()

  fmt.Println("connection handled")
  for {
    input, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Printf("%s",input)
    for c := clients.Front(); c != nil; c = c.Next() {
        var b bytes.Buffer
        b.Write([]byte("\033[42m"))
        for i := 0; i < 80/*utf8.RuneCountInString(input) ; i++ {
          // do stuff
          /*
          b.Write([]byte(" "))
        }
        b.Write([]byte("\033[49m"))
        writer := bufio.NewWriter((c.Value).(net.Conn))
        writer.WriteString(b.String())
        writer.Flush()/*
      if c != connp {
        writer := bufio.NewWriter((c.Value).(net.Conn))
        writer.WriteString(input)
        writer.Flush()
      }
    }
  }
}


*/
