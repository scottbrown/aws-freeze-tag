# aws-freeze-tag

Discovers tags on an EC2 instance and saves them to disk in JSON format

This tool is primarily used by user data scripts to discovers the tags
given to an EC2 machine.  The tool attempts to download all of the tags
in an efficient manner, using only one network call.

This tool uses the EC2 metadata service, which is only accessible from
EC2 machines.  Therefore, you must run this tool on an EC2 machine.  It
will not work outside of EC2.

## Usage

```
$ freezetag --directory /etc/aws
Wrote tags to /etc/aws/tags.json

$ jq . /etc/aws/tags.json
[
  {
    "Key": "aws:cloudformation:stack-name",
    "Value": "example-stack"
  },
  {
    "Key": "Name",
    "Value": "example-server"
  }
]
```

