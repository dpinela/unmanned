# unmanned

A wrapper over man(1) that displays manpages in your browser instead of the console. It uses
[mandoc][] to render pages and runs a (locally-accessible only by default) web server that lets
you navigate between them, without having to pre-convert them all to HTML first.

It does not provide any indexing or search features; just a way to read manpages locally in a
nicer format than the console allows.

## How to use

Ensure you have mandoc installed, then run `unmanned <page>` or `unmanned <section> <page>`
just like you would do with regular man. When you're done browsing manuals, press Ctrl-C
to shut down the server.

If you just want to start the server, run just `unmanned`.

In any of these cases, you can add `-p <address>:<port>` to specify the address and port that
unmanned should listen on. By default, it binds to the loopback interface on a system-chosen
port.

### URL structure

- `/<page>`: displays the same page as `man <page>`
- `/<section>/<page>`: displays the same page as `man <section> <page>`

[mandoc]: https://mandoc.bsd.lv