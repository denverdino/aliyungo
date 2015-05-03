# AliyunGo: Go SDK for Aliyun Services
================
This is an unofficial Go SDK for Aliyun Services



Package Structure
-------------------

*  ecs: [Elastic Compute Service](http://docs.aliyun.com/?spm=5176.100054.3.1.1uupTM#/ecs)
*  oss: [Open Storage Service](http://docs.aliyun.com/?spm=5176.100054.3.5.1uupTM#/oss)
*  util: Utility helpers

Documentation
-------------------
*  ecs: [https://godoc.org/github.com/denverdino/aliyungo/ecs](https://godoc.org/github.com/denverdino/aliyungo/ecs)
*  oss: [https://godoc.org/github.com/denverdino/aliyungo/oss](https://godoc.org/github.com/denverdino/aliyungo/oss)

Test ECS
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

Test OSS
-------------------

Modify "oss/config_test.go" 

```sh
	TEST_ACCESS_KEY_ID     = "MY_ACCESS_KEY_ID"
	TEST_ACCESS_KEY_SECRET = "MY_ACCESS_KEY_ID"
	TEST_I_AM_RICH         = false
	TEST_REGION            = BEIJING
	TEST_BUCKET            = "denverdino"
```

*  TEST_ACCESS_KEY_ID: the Access Key Id
*  TEST_ACCESS_KEY_SECRET: the Access Key Secret.
*  TEST_I_AM_RICH(Optional): If it is set to true, it will perform tests to create virtual machines and disks under your account. And you will pay the bill. :-)
*  TEST_REGION: the region of OSS for testing
*  TEST_BUCKET: the bucket name for testing


Under "oss" and run

```sh
go test
```

License and Authors
-------------------
Authors: denverdino@gmail.com