rabbitmq-connection-tester
==========================
A tool for testing a rabbitmq connection.

You would only need this tool if you are using an older version of `rabbitmqctl` which does not already contain the `authenticate_user` command or perhaps if you'd like to test the connection from another machine that doesn't have rabbitmq installed.

Note: This is just a utility I've thrown together because I personally needed it quickly. It isn't written well and is not full featured.

Install
=======

    curl -L https://github.com/kyle-long/rabbitmq-connection-tester/releases/download/v1.0/rabbitmq-connection-tester > rabbitmq-connection-tester
    chmod +x rabbitmq-connection-tester
