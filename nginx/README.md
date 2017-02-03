# polly-nginx

This directory contains parts of our NGINX server configuration. These
are not complete and serve merely as an example. Handle with care.

The [NGINX documentation](http://nginx.org/en/docs/) is actually fairly
good. Before you apply any of the statements you see in this directory
please read up on the respective command there! 

Things also change over time, both in terms of NGINX directives, but also
in terms of best practices for SSL etc.; do not expect this repository
to reflect the latest developments, it is merely a snapshot.

## Features

* Log Sanitation: do not log the last IPv4 octet/IPv6 hextet of visitors
* LetsEncrypt ACME
* GZIP Compression Optimization

## TODO

* Content-Security-Policy
