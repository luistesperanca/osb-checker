package common

import (
	"context"

	openapi "github.com/openservicebrokerapi/osb-checker/autogenerated/go-client"
	. "github.com/openservicebrokerapi/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func testPollInstanceLastOperation(
	instanceID string,
) {
	Convey("should return 412 PreconditionFailed if missing X-Broker-API-Version header", func() {
		_, resp, err := cli.ServiceInstancesApi.ServiceInstanceLastOperationGet(
			authCtx, "", instanceID, &openapi.ServiceInstanceLastOperationGetOpts{})

		So(err, ShouldNotBeNil)
		So(resp.StatusCode, ShouldEqual, 412)
	})

	Convey("should return 401 Unauthorized if missing Authorization header", func() {
		_, resp, err := cli.ServiceInstancesApi.ServiceInstanceLastOperationGet(
			context.Background(), CONF.APIVersion, instanceID,
			&openapi.ServiceInstanceLastOperationGetOpts{})

		So(err, ShouldNotBeNil)
		So(resp.StatusCode, ShouldEqual, 401)
	})

	Convey("should accept a valid poll instance last operation request", func() {
		_, resp, err := cli.ServiceInstancesApi.ServiceInstanceLastOperationGet(
			authCtx, CONF.APIVersion, instanceID,
			&openapi.ServiceInstanceLastOperationGetOpts{})

		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
	})
}

func testPollBindingLastOperation(
	instanceID, bindingID string,
) {
	Convey("should return 412 PreconditionFailed if missing X-Broker-API-Version header", func() {
		_, resp, err := cli.ServiceBindingsApi.ServiceBindingLastOperationGet(
			authCtx, "", instanceID, bindingID,
			&openapi.ServiceBindingLastOperationGetOpts{})

		So(err, ShouldNotBeNil)
		So(resp.StatusCode, ShouldEqual, 412)
	})

	Convey("should return 401 Unauthorized if missing Authorization header", func() {
		_, resp, err := cli.ServiceBindingsApi.ServiceBindingLastOperationGet(
			context.Background(), CONF.APIVersion, instanceID, bindingID,
			&openapi.ServiceBindingLastOperationGetOpts{})

		So(err, ShouldNotBeNil)
		So(resp.StatusCode, ShouldEqual, 401)
	})

	Convey("should accept a valid poll binding last operation request", func() {
		_, resp, err := cli.ServiceBindingsApi.ServiceBindingLastOperationGet(
			authCtx, CONF.APIVersion, instanceID, bindingID,
			&openapi.ServiceBindingLastOperationGetOpts{})

		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
	})
}