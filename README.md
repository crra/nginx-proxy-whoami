# About

Refreshed version of [whoami](https://github.com/jwilder/whoami), to test an [nginx docker proxy](https://github.com/nginx-proxy/nginx-proxy) setup. Works on the raspberry pi 4 (ARM).

# Usage

    $ docker-compose build
    $ docker-compose up -d

    $ curl -H "Host: whoami.local" localhost
    > I'm f619deadbeef
