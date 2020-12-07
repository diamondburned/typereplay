# typereplay

A small program that parses an input format to replay keyboard commands.

## Dependencies

### Linux

- Robotgo dependencies
	- Whatever Nix defines the build inputs for `st` is
	- `with xlibs`
		- `libXi`
		- `libXtst`
		- `libxkbcommon`

## Documentation

### Tape Syntax

```ini
# A comment is prefixed with a hash/pound/whatever.

# The default duration is 100ms. setduration will apply to the whole program
# regardless of where it's put.
setduration 100ms

# Note that trailing spaces on either ends will be trimmed.

type This line will be inputted followed by an enter press.

puts This line will not have an enter follow it.
enter

tap up
enter
tap up

type We're inserting another line here.

# Wait for an extra 100ms before running the next command.
wait 100ms

type End of program.
```

### CLI Usage

```sh
./typereplay \
	-w 5s    # wait for 5 seconds before starting
	-i input # the input file
# all flags are optional.
```
