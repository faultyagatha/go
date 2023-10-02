# Effective Go

## Packages

- lower case, single-word names; there should be no need for underscores or mixedCaps
- the package name is the base name of its source directory; the package in src/encoding/base64 is imported as "encoding/base64" but has name base64, not encoding_base64 and not encodingBase64.

## Getters
- no support for getters and setters but it's ok to have them
- idiomatic practice:
  - no `Get` into the getter's name
  - e.g., if you have a field called `owner` (lower case, unexported), the getter method should be called `Owner` (upper case, exported), not GetOwner. 
  - the use of upper-case names for export provides the hook to discriminate the field from the method
  - a setter function, if needed, will likely be called SetOwner

  ```go
  owner := obj.Owner()
  if owner != user {
    obj.SetOwner(user)
  }
  ```