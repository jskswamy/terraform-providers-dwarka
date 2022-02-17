# Terraform Provider Dwarka

This repo is a companion repo to the learn terraform

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-dwarka
```

## Local release build

```shell
$ go install github.com/goreleaser/goreleaser@latest
```

```shell
$ make release
```

You will find the releases in the `/dist` directory. You will need to rename the provider binary to `terraform-provider-dwarka` and move the binary into [the appropriate subdirectory within the user plugins directory](https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers#install-hashicups-provider).
## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
