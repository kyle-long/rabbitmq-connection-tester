//usr/bin/env go run $0 $@; exit;
package main

import (
    "fmt"
    "log"
    "github.com/streadway/amqp"
    "github.com/docopt/docopt-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
    doc := `
Usage:
    rabbitmq-connection-tester <username> <password> <host> [options]

Options:
    -p port, --port port            Port to connect to. [default: 5672]
    -v vhost, --vhost vhost         vhost to connect to. [default: /]

    `

    a, _ := docopt.Parse(doc, nil, true, "rabbitmq-connection-tester", false)
    connection_string := fmt.Sprintf("amqp://%s:%s@%s:%s%s", a["<username>"], a["<password>"], a["<host>"], a["--port"], a["--vhost"])
    log.Print("Attempting to connect with: " + connection_string)
	conn, err := amqp.Dial(connection_string)
    failOnError(err, "Failed to connect.")
    fmt.Println("Successfully connected")
    conn.Close()
}
