# formicidae
Update .env files in Go with code.

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/papasavva/formicidae/CI?style=flat-square)](https://github.com/papasavva/formicidae/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/papasavva/formicidae?style=flat-square)](https://goreportcard.com/report/github.com/papasavva/formicidae)

## What is fomicidae?
formicidate is a small tool for Go application. You can update the value of environment variables in a .env file with code.
This can be handy in situations that you want to update the values programmatically like in unit tests.  

## Installation
```shell
go get github.com/papasavva/formicidae
```

## Usage
Add your application configuration to your .env file of your project:

```shell
S3_BUCKET=myBucket
SECRET_KEY=mySecretKeyGoesHere
```

Then in your app you can do something like:
```go
package main

import (
    "github.com/papasavva/formicidae"
)

func main() {
	updatedContent, err := formicidae.UpdateVariable(".env", "SECRET_KEY", "myNewSecretKey")
	...
}
```

## Notes
- If you wish to update multiple variables, you can call the function multiple times.
- The function will return the content of the whole file.

## Contributing
Contributions are welcomed.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Added some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request


## Licence
The code is licensed under the permissive MIT licence. [Read this](https://www.tldrlegal.com/l/mit) for a summary.