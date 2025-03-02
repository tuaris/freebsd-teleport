# Teleport for FreeBSD

Currently just the client and cluster management tool

## Status

There's a slight bug while fetching go modules:

```
go: github.com/gravitational/teleport/api@v0.0.0 (replaced by ./api): reading api/go.mod: open /usr/ports/distfiles/go/shells_teleport/gravitational-teleport-v17.3.0_GH0/api/go.mod: no such file or directory
```

The work around for now, is to extract the `api` folder from the tar archive and 
put it where it's expected.  If you know how to fix this, please share.

The client and cluster management tool build, but there's some clean up needed
with patching to do things "the right way".

I haven't tested anything to see if it actually conencts and works.  All I know
is it builds and runs locally.

