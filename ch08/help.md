# redis-benchmark --help
Usage: redis-benchmark [-h <host>] [-p <port>] [-c <clients>] [-n <requests>] [-k <boolean>]

-h <hostname>      Server hostname (default 127.0.0.1)
-p <port>          Server port (default 6379)
-s <socket>        Server socket (overrides host and port)
-a <password>      Password for Redis Auth
--user <username>  Used to send ACL style 'AUTH username pass'. Needs -a.
-c <clients>       Number of parallel connections (default 50)
-n <requests>      Total number of requests (default 100000)
-d <size>          Data size of SET/GET value in bytes (default 3)
--dbnum <db>       SELECT the specified db number (default 0)
--threads <num>    Enable multi-thread mode.
--cluster          Enable cluster mode.
--enable-tracking  Send CLIENT TRACKING on before starting benchmark.
-k <boolean>       1=keep alive 0=reconnect (default 1)
-r <keyspacelen>   Use random keys for SET/GET/INCR, random values for SADD,
random members and scores for ZADD.
Using this option the benchmark will expand the string __rand_int__
inside an argument with a 12 digits number in the specified range
from 0 to keyspacelen-1. The substitution changes every time a command
is executed. Default tests use this to hit random keys in the
specified range.
-P <numreq>        Pipeline <numreq> requests. Default 1 (no pipeline).
-q                 Quiet. Just show query/sec values
--precision        Number of decimal places to display in latency output (default 0)
--csv              Output in CSV format
-l                 Loop. Run the tests forever
-t <tests>         Only run the comma separated list of tests. The test
names are the same as the ones produced as output.
-I                 Idle mode. Just open N idle connections and wait.
--tls              Establish a secure TLS connection.
--sni <host>       Server name indication for TLS.
--cacert <file>    CA Certificate file to verify with.
--cacertdir <dir>  Directory where trusted CA certificates are stored.
If neither cacert nor cacertdir are specified, the default
system-wide trusted root certs configuration will apply.
--insecure         Allow insecure TLS connection by skipping cert validation.
--cert <file>      Client certificate to authenticate with.
--key <file>       Private key file to authenticate with.
--tls-ciphers <list> Sets the list of prefered ciphers (TLSv1.2 and below)
in order of preference from highest to lowest separated by colon (":").
See the ciphers(1ssl) manpage for more information about the syntax of this string.
--tls-ciphersuites <list> Sets the list of prefered ciphersuites (TLSv1.3)
in order of preference from highest to lowest separated by colon (":").
See the ciphers(1ssl) manpage for more information about the syntax of this string,
and specifically for TLSv1.3 ciphersuites.
--help             Output this help and exit.
--version          Output version and exit.

Examples:

Run the benchmark with the default configuration against 127.0.0.1:6379:
$ redis-benchmark

Use 20 parallel clients, for a total of 100k requests, against 192.168.1.1:
$ redis-benchmark -h 192.168.1.1 -p 6379 -n 100000 -c 20

Fill 127.0.0.1:6379 with about 1 million keys only using the SET test:
$ redis-benchmark -t set -n 1000000 -r 100000000

Benchmark 127.0.0.1:6379 for a few commands producing CSV output:
$ redis-benchmark -t ping,set,get -n 100000 --csv

Benchmark a specific command line:
$ redis-benchmark -r 10000 -n 10000 eval 'return redis.call("ping")' 0

Fill a list with 10000 random elements:
$ redis-benchmark -r 10000 -n 10000 lpush mylist __rand_int__

On user specified command lines __rand_int__ is replaced with a random integer
with a range of values selected by the -r option.


