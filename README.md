rlistener
=========

You can now use a `net.Listener` on a remote server!

Installation
------------

```
go get github.com/thinxer/rlistener
```

Usage
-----

1. Create an access point with `rlistener.NewServer`.
2. `rlistener.Dial` to the access point and get a `net.Listener`!

Examples
--------

See [examples](examples). You can first start both `recho` and `rlistener`, and then try `nc` to the `rlistener` bind address.
