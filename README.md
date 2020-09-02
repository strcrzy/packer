你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Packer

* Website: http://www.packer.io
* IRC: `#packer-tool` on Freenode
* Mailing list: [Google Groups](http://groups.google.com/group/packer-tool)

Packer is a tool for building identical machine images for multiple platforms
from a single source configuration.

Packer is lightweight, runs on every major operating system, and is highly
performant, creating machine images for multiple platforms in parallel.
Packer comes out of the box with support for creating AMIs (EC2), VMware
images, and VirtualBox images. Support for more platforms can be added via
plugins.

The images that Packer creates can easily be turned into
[Vagrant](http://www.vagrantup.com) boxes.

## Quick Start

**Note:** There is a great
[introduction and getting started guide](http://www.packer.io/intro)
for those with a bit more patience. Otherwise, the quick start below
will get you up and running quickly, at the sacrifice of not explaining some
key points.

First, [download a pre-built Packer binary](http://www.packer.io/downloads.html)
for your operating system or [compile Packer yourself](#developing-packer).

After Packer is installed, create your first template, which tells Packer
what platforms to build images for and how you want to build them. In our
case, we'll create a simple AMI that has Redis pre-installed. Save this
file as `quick-start.json`. Be sure to replace any credentials with your
own.

```json
{
  "builders": [{
    "type": "amazon-ebs",
    "access_key": "YOUR KEY HERE",
    "secret_key": "YOUR SECRET KEY HERE",
    "region": "us-east-1",
    "source_ami": "ami-de0d9eb7",
    "instance_type": "t1.micro",
    "ssh_username": "ubuntu",
    "ami_name": "packer-example {{timestamp}}"
  }]
}
```

Next, tell Packer to build the image:

```
$ packer build quick-start.json
...
```

Packer will build an AMI according to the "quick-start" template. The AMI
will be available in your AWS account. To delete the AMI, you must manually
delete it using the [AWS console](https://console.aws.amazon.com/). Packer
builds your images, it does not manage their lifecycle. Where they go, how
they're run, etc. is up to you.

## Documentation

Full, comprehensive documentation is viewable on the Packer website:

http://www.packer.io/docs

## Developing Packer

If you wish to work on Packer itself, you'll first need [Go](http://golang.org)
installed (version 1.2+ is _required_). Make sure you have Go properly installed,
including setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH).

For some additional dependencies, Go needs [Mercurial](http://mercurial.selenic.com/)
and [Bazaar](http://bazaar.canonical.com/en/) to be installed.
Packer itself doesn't require these, but a dependency of a dependency does.

You'll also need [`gox`](https://github.com/mitchellh/gox)
to compile packer. You can install that with:

```
$ go get -u github.com/mitchellh/gox
```

Next, clone this repository into `$GOPATH/src/github.com/mitchellh/packer` and
then just type `make`. In a few moments, you'll have a working `packer` executable:

```
$ make
...
$ bin/packer
...
```

You can run tests by typing `make test`.

This will run tests for Packer core along with all the core builders and commands and such that come with Packer.

If you make any changes to the code, run `make format` in order to automatically
format the code according to Go standards.

When new dependencies are added to packer you can use `make updatedeps` to
get the latest and subsequently use `make` to compile and generate the `packer` binary.
