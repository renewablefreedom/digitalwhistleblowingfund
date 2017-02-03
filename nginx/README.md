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

## :warning: WARNING :warning:

* Sets expiry of many files to maximum, because this works well for EmberJS; this is a terrible idea in many other cases
* HSTS Preload is enabled in this configuration. You need to educate yourself and think hard whether you want that for your own site(s). [This is dangerous.](https://scotthelme.co.uk/death-by-copy-paste/)

## TODO

* Content-Security-Policy
