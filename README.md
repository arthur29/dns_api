# Requisites

To run the project is necessary Docker installed on your system

# Build

Execute command: `make build-docker-image`

# Run

Execute command: `make run`

So DNS is on port 53 and API on port 9000

# Issues

* The only route that is working is `/list` to show DNS entries
* The tests is not completely
* The image base is alpine and the docs recommend to not use it, so how it is small and no problem was found it was used and in future will be changed
