
# about

This project is meant for displaying a small list of tasks in an e-ink display.

# structure

## backend

Sweet monolith for handling all the operations. It connects to a google calendar (that you will be able to configure).

## sbc-client

The client for the single board computer (hence sbc). Still haven't decided if this will be a raspberry pico or an arduino.

# tech stack

* `go` for the backend
* `tinygo` for te sbc
