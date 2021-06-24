package common

import (
	"context"
	"testing"

	openapi "github.com/luistesperanca/osb-checker/autogenerated/go-client"
	. "github.com/luistesperanca/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetInstance(
	t *testing.T,
	instanceID string,
) {
	Convey("RETRIEVE INSTANCE", t, func() {

		Convey("should return 412 PreconditionFailed if missing X-Broker-API-Version header", func() {
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceGet(
				authCtx, "", instanceID, &openapi.ServiceInstanceGetOpts{})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 412)
		})

		if CONF.Authentication.AuthType != TypeNoauth {
			Convey("should return 401 Unauthorized if missing Authorization header", func() {
				_, resp, err := cli.ServiceInstancesApi.ServiceInstanceGet(
					context.Background(), CONF.APIVersion, instanceID,
					&openapi.ServiceInstanceGetOpts{})

				So(err, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, 401)
			})
		}

		Convey("should accept a valid get instance request", func() {
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceGet(
				authCtx, CONF.APIVersion, instanceID,
				&openapi.ServiceInstanceGetOpts{})

			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 200)
		})
	})
}

func TestGetBinding(
	t *testing.T,
	instanceID, bindingID string,
) {
	Convey("RETRIEVE BINDING", t, func() {

		Convey("should return 412 PreconditionFailed if missing X-Broker-API-Version header", func() {
			_, resp, err := cli.ServiceBindingsApi.ServiceBindingGet(
				authCtx, "", instanceID, bindingID,
				&openapi.ServiceBindingGetOpts{})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 412)
		})

		if CONF.Authentication.AuthType != TypeNoauth {
			Convey("should return 401 Unauthorized if missing Authorization header", func() {
				_, resp, err := cli.ServiceBindingsApi.ServiceBindingGet(
					context.Background(), CONF.APIVersion, instanceID, bindingID,
					&openapi.ServiceBindingGetOpts{})

				So(err, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, 401)
			})
		}

		Convey("should accept a valid get binding request", func() {
			_, resp, err := cli.ServiceBindingsApi.ServiceBindingGet(
				authCtx, CONF.APIVersion, instanceID, bindingID,
				&openapi.ServiceBindingGetOpts{})

			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 200)
		})
	})
}
