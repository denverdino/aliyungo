# Aliyun-Go: Go SDK for Aliyun Services
================
This is an unofficial Go SDK for Aliyun Services

Package Structure
-------------------

*  ecs: API for Elastic Compute Service
*  util: Utility helpers


Test
-------------------
Modify "ecs/config_test.go" with your Access Key Id and Access Key Secret.

Set TEST_INSTANCE_ID with the existing instance id.
Set TEST_I_AM_RICH with true to perform some actions to create virtual machines and disks. And you need pay your bill for it. :-)

```sh
	TEST_ACCESS_KEY_ID     = "MY_ACCESS_KEY_ID"
	TEST_ACCESS_KEY_SECRET = "MY_ACCESS_KEY_ID"
	TEST_INSTANCE_ID       = "MY_INSTANCE_ID"
	TEST_I_AM_RICH         = false
```

Under aliyun-go/ecs and run

```sh
go test
```


License and Authors
-------------------
Authors: denverdino@gmail.com