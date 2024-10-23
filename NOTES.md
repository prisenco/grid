
This is a very htmx-friendly problem, so I may have handicapped myself a bit
by not using it (and crudely recreating the functionality of HTMX in vanilla
js).

**Refactoring**

1. Move functionality out of single `main.go` file and into separate files
for cleaner organization.
2. Move the markup to `tpl/` directory, embed using `go:embed` directive
and parse using `fasttemplate` library. Makes for easier updating of markup
as regular html and not mixed in with go code.
3. Move js and css to files that are statically served. These could be 
embedded the way the templates are so that the binary would include them
for deploy.
4. Only send back the grid itself, so we aren't double-loading the css and
js.
5. Removed some unnecessary code, like the `<form>` tag

**Additional thoughts**

1. This should be multi-user safe, since the browser isn't holding the state and
any updates of the grid result in the full grid being returned.
2. There are still possibilities of race conditions though (user A and B
update the grid at the same time, so one user gets back the other users
value since it's newer by nanoseconds). Gets really complicated if you want
to prevent this.
3. Sending the whole grid back works for a small grid since the markup
isn't much.
4. If you were to do go for a huge grid (say, 10000x10000) it would have to
isolate just the viewable part of the grid and update that, otherwise request
payload would be too big.
5. If this were in production, it would need rate-limiting, timeouts,
structured logging and all that good stuff.
