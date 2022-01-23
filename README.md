# formicidae
Update .env files in Go

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/papasavva/formicidae/CI?style=flat-square)](https://github.com/papasavva/formicidae/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/papasavva/formicidae?style=flat-square)](https://goreportcard.com/report/github.com/papasavva/formicidae)

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
}
```
If you wish to update multiple variables, you can call the function multiple times.

## Licence
The code is licensed under the permissive MIT licence. [Read this](https://www.tldrlegal.com/l/mit) for a summary.