<!-- DO NOT REMOVE - contributor_list:data:start:["gleich", "dependabot-preview[bot]"]:end -->

# statuser ![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/gleich/statuser?include_prereleases)

📣 A user-friendly status outputting library for go

![build](https://github.com/gleich/statuser/workflows/build/badge.svg)
![lint](https://github.com/gleich/statuser/workflows/lint/badge.svg)
![test](https://github.com/gleich/statuser/workflows/test/badge.svg)

## 🚀 Installing

Simply run the command below in the root of your go project:

```bash
go get -u github.com/gleich/statuser/v2
```

## 📚 Documentation [![GoDoc](https://godoc.org/github.com/gleich/statuser/v2?status.svg)](https://pkg.go.dev/github.com/gleich/statuser/v2)

Here is some basic usage:

### `Error()`

Output an error to the user

#### Parameters

1. message
   - Type: `string`
   - Description: A human readable message to help the user
2. err
   - Type: `err`
   - Description: The golang error
3. exitCode
   - Type: `int`
   - Description: Exit code used to exit the program

#### Example

Outputs the following in red text:

```
░░░░░░░░░░░░░
░🚨 ERROR 🚨░
░░░░░░░░░░░░░

message

GOLANG ERROR (SHOW DEVELOPER):
err
exit status 1
```

### `ErrorMsg()`

Output an error to the user with just a message

#### Parameters

1. message
   - Type: `string`
   - Description: A human readable message to help the user
2. exitCode
   - Type: `int`
   - Description: Exit code used to exit the program

#### Example

Outputs the following in red text:

```
░░░░░░░░░░░░░
░🚨 ERROR 🚨░
░░░░░░░░░░░░░

message
exit status 1
```

### `Warning()`

Output a warning the user

#### Parameters

1. message
   - Type: `string`
   - Description: The warning message to display to the user

#### Example

Outputs the following text in yellow:

```
⚠️ WARNING ⚠️
message
```

### `Success()`

Output a warning the user

#### Parameters

1. message
   - Type: `string`
   - Description: The success message to display to the user

#### Example

Outputs the following text in green:

```
✅ message
```

## Contributors

1. Matthew Gleich ([@gleich](http://www.github.com/gleich))

<!-- DO NOT REMOVE - contributor_list:start -->
## 👥 Contributors


- **[@gleich](https://github.com/gleich)**

- **[@dependabot-preview[bot]](https://github.com/apps/dependabot-preview)**

<!-- DO NOT REMOVE - contributor_list:end -->
