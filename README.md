# Packer Builder - DefaultConf

This repository intends to be used as a demo for a question / feature request
to Packer.

The main question is: **How to properly assign defaults to some configurations
fields of a Packer Builder?**. This can, of course, probably be applied to
other components such as provisionners and post-processors.

## How to run this demo

1. Setup requirements:

```bash
glide install
```

2. Go to either `defaultconf_preset` or `defaultconf_2` depending of which solution
   you want to see/test:

```bash
cd defaultconf_preset/
# OR
cd defaultconf_postcmp/
```

3. Build:

```bash
# This will just compile the plugin binary
make
```

4. Run a really simple call to this builder:

```bash
# This will call a "packer build" on a really simple configuration
make test
```

## Considered solutions

Let's take this usecase, made-of for this demo:

- I have 2 Parameters for my Builder
    1. `Create`, a boolean that represents if we should create something or not
    2. `Number`, an integer that represents the number of things to create
- I want `Create` to default to `true`
- I want `Number` to default to `42`

### Solution 1 - Pass defaults to `config.Decode()`

The
[`Decode()`](https://github.com/hashicorp/packer/blob/master/helper/config/decode.go#L14)
function of Packer doesn't look like it is able to take default values to
assign if some parameters are not found in the configuration, neither does the
underlying [mapstructure](https://github.com/mitchellh/mapstructure) library


- I can't just go the if `IntParam == 0 { IntParam = 42 }` way after the call
  to config.Decode(), because then it wouldn't allow people to explicitely give
  IntParam: 0 as parameter
- I could have a separate function like InitConfig() that is called right
  before config.Decode(), but that seems a bit weird to me (I will do it if
  that's the only way, but expected to be able to do something cleaner)

### Solution 2

Idea:

> Let's just have something that initializes the Config with default values
> before the `config.Decode()` happens

Problem:

The `config.Decode()` initializes a totally new config and assigns it,
overriding pre-set values. This is demonstrated in the `defaultconf_preset`
directory:

```raw
NewBuilder(): Creating new Builder...
NewBuilder(): Initializing new Config...
NewBuilder(): Assigning value Create = true
NewBuilder(): Assigning value Number = 42
NewBuilder(): Created new Builder. Returning.
[...]
Prepare(): Starting...
Prepare(): Current config:  &{{  false false  map[]} {<nil> map[] map[] false   } true 42}
Prepare(): Create config Parameter:  true
Prepare(): Number config Parameter:  42
Prepare(): Calling config.Decode()...
Prepare(): Called config.Decode()
Prepare(): Current config:  &{{defaultconf defaultconf false false  map[]} {<nil> map[] map[] false   } false 0}
Prepare(): Create config Parameter:  false
Prepare(): Number config Parameter:  0
Prepare(): Prepared. Returning.
```

### Solution 3

Idea:

> Let's check the values after the `config.Decode()` and change them to
> defaults if the user hasn't changed them

Problem:

**There is no way to know if the user defined a value or not**, we can just
check if the value is the type's default one; what if the user (him/her)self
defined it to the type's default? This is demonstrated in the
`defaultconf_postcmp` directory.

### Solution 4

Suggested by *mcbadass* on IRC (`#packer-tool` on Freenode).

Just accept strings as parameters, and parse/sanitize it.

- `Create` would then no longer be a boolean, but a string, and we would then
  be able to differenciate empty string

wait. In any way, such a thing is counter-intuitive and it would be required to
read the documentation.
