# AliyunGo: Go SDK for Aliyun Services
================
This is an unofficial Go SDK for Aliyun Services



Package Structure
-------------------

*  ecs: [API for Elastic Compute Service]()
*  util: Utility helpers

Documentation
-------------------
*  ecs: [https://godoc.org/github.com/denverdino/aliyungo/ecs](https://godoc.org/github.com/denverdino/aliyungo/ecs)


Test
-------------------

Modify "ecs/config_test.go" 

```sh
	TEST_ACCESS_KEY_ID     = "MY_ACCESS_KEY_ID"
	TEST_ACCESS_KEY_SECRET = "MY_ACCESS_KEY_ID"
	TEST_INSTANCE_ID       = "MY_INSTANCE_ID"
	TEST_I_AM_RICH         = false
```

*  TEST_ACCESS_KEY_ID: the Access Key Id
*  TEST_ACCESS_KEY_SECRET: the Access Key Secret.
*  TEST_INSTANCE_ID: the existing instance id for testing. It will be stopped and restarted during testing.
*  TEST_I_AM_RICH(Optional): If it is set to true, it will perform tests to create virtual machines and disks under your account. And you will pay the bill. :-)

Under "ecs" and run

```sh
go test
```

License and Authors
-------------------
Authors: denverdino@gmail.com