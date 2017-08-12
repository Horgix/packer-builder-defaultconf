# Packer Builder - DefaultConf

This repository intends to be used as a demo for a question / feature request
to Packer.

The main question is: **How to properly assign defaults to some configurations
field of a Packer Builder?**. This can, of course, probably be applied to other
components such as provisionners and post-processors.

## How to run this demo

1. Setup requirements:

```bash
glide install
```

2. Go to either `defaultconf_1` or `defaultconf_2`

```bash
cd defaultconf_1/
# OR
cd defaultconf_2/
```

3. Build:

```bash
make
```

4. Run a really simple call to this builder:

```bash
make test
```

## Considered solutions

Let's take this (made-of for the example) usecase:

- I have 2 Parameters for my Builder
    1. `Create`, a boolean that represents if we should create something or not
    2. `Number`, an integer that represents the number of things to create
- I want `Create` to default to `true`
- I want `Number` to default to 42


- The
  [`Decode()`](https://github.com/hashicorp/packer/blob/master/helper/config/decode.go#L14)
  function of Packer doesn't look like it is able to take default values to
  assign if some parameters are not found in the configuration, neither does
  the underlying mapstructure lib
- I can't just go the if `IntParam == 0 { IntParam = 42 }` way after the call
  to config.Decode(), because then it wouldn't allow people to explicitely give
  IntParam: 0 as parameter
- I could have a separate function like InitConfig() that is called right
  before config.Decode(), but that seems a bit weird to me (I will do it if
  that's the only way, but expected to be able to do something cleaner)

### Solution 1

> Let's just have something that initializes the Config with default values
> before the `config.Decode()` happens

The reason this solution is innapropriate is demonstrated in the
`defaultconf_1` directory.

Result:

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

### Solution 2

> Let's check the values after the `config.Decode()` and change them to
> defaults if the user hasn't changed them

The reason this solution is annapropriate is demonstrated in the
`defaultconf_1` directory.

The main problem is that **there is no way to know if the user defined a value
or not**, we can just check if the value is the type's default one; what if the
user (him/her)self defined it to the type's default one?

## Details




## The issue submitted to the Packer repo


