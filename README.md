# advent-of-code
Advent of Code 

For the 2022 you can test the go programs if you have podman installed with:

```
podman run --rm -v "$PWD":/usr/src/myapp:Z -w /usr/src/myapp golang:1.19.4 go run <name_of_program.go>
```
If you have docker, probably you can run the same command but replacing `podman` with `docker` in the command.
