# log2ltsv
apache and nginx access log to ltsv format

## Installation

Download from https://github.com/tkuchiki/log2ltsv/releases

## Usage

```console
usage: log2ltsv [<flags>]

apache and nginx access log to ltsv format

Flags:
  --help            Show context-sensitive help (also try --help-long and --help-man).
  --format="nginx"  access log format (apache or nginx)
  --file=FILE       access log
  --version         Show application version.
```

## Example

### Apache

```diff
- LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-Agent}i\"" combined
+LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-Agent}i\" %T" combined
```

```console
$ cat logs/apache.access.log
127.0.0.1 - - [03/Sep/2017:12:23:39 +0000] "GET / HTTP/1.1" 403 3839 "-" "curl/7.47.1" 1
$ cat logs/apache.access.log | ./log2ltsv --format apache
user:- 	time:2017-09-03T12:23:39+00:00	method:GET	uri:/	status:403	size:3839	reqtime:1	apptime:1
```

### Nginx

```
log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                  '$status $body_bytes_sent "$http_referer" '
-                  '"$http_user_agent" "$http_x_forwarded_for"';
+                  '"$http_user_agent" "$http_x_forwarded_for" '
+                  '$request_time $upstream_response_time';
```

```console
$ cat logs/nginx.access.log
127.0.0.1 - - [03/Sep/2017:12:21:17 +0000] "GET / HTTP/1.1" 200 3770 "-" "curl/7.47.1" "-" 0.0 0.1
$ cat logs/nginx.access.log | ./log2ltsv
user:-	time:2017-09-03T12:21:17+00:00	method:GET	uri:/	status:200	size:3770	reqtime:0.0	apptime:0.1
```
