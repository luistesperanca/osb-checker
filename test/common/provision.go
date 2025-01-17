package common

import (
	"context"
	"testing"

	"github.com/antihax/optional"
	openapi "github.com/luistesperanca/osb-checker/autogenerated/go-client"
	. "github.com/luistesperanca/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProvision(
	t *testing.T,
	instanceID string,
	req *openapi.ServiceInstanceProvisionRequest,
	async bool,
) {
	Convey("PROVISIONING - request syntax", t, func() {

		Convey("should return 412 PreconditionFailed if missing X-Broker-API-Version header", func() {
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, "", instanceID, openapi.ServiceInstanceProvisionRequest{},
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 412)
		})

		if CONF.Authentication.AuthType != TypeNoauth {
			Convey("should return 401 Unauthorized if missing Authorization header", func() {
				_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
					context.Background(), CONF.APIVersion, instanceID,
					openapi.ServiceInstanceProvisionRequest{},
					&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

				So(err, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, 401)
			})
		}

		if async {
			Convey("should return 422 UnprocessableEntity if missing accepts_incomplete", func() {
				tempBody := openapi.ServiceInstanceProvisionRequest{}
				deepCopy(req, &tempBody)
				_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
					authCtx, CONF.APIVersion, instanceID, tempBody,
					&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(false)})

				So(err, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, 422)
			})
		}

		Convey("should reject if missing service_id", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			tempBody.ServiceId = ""
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})

		Convey("should reject if missing plan_id", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			tempBody.PlanId = ""
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})

		Convey("should reject if request payload is missing organization_guid", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			tempBody.OrganizationGuid = ""
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})

		Convey("should reject if request payload is missing space_guid", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			tempBody.SpaceGuid = ""
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})

		Convey("should reject if service_id is invalid", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			tempBody.ServiceId = "xxxx-xxxx-xxxx-xxxx"
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})

		Convey("should reject if plan_id is invalid", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			tempBody.PlanId = "xxxx-xxxx-xxxx-xxxx"
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})

		Convey("should reject if parameters are not following schema", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			tempBody.Parameters = map[string]interface{}{
				"can-not": "be-good",
			}
			if err := testCatalogSchema(&SchemaOpts{
				ServiceID:  tempBody.ServiceId,
				PlanID:     tempBody.PlanId,
				Parameters: tempBody.Parameters,
				SchemaType: TypeServiceInstance,
				Action:     ActionCreate,
			}); err == nil {
				return
			}
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldNotBeNil)
			So(resp.StatusCode, ShouldEqual, 400)
		})
	})

	Convey("PROVISIONING - new", t, func() {
		Convey("should accept a valid provision request", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldBeNil)
			if async {
				So(resp.StatusCode, ShouldEqual, 202)
			} else {
				So(resp.StatusCode, ShouldEqual, 201)
			}
		})
	})

	if async {
		Convey("PROVISIONING - poll", t, func(c C) {
			testPollInstanceLastOperation(instanceID)

			So(pollInstanceLastOperationStatus(instanceID, "provision"), ShouldBeNil)
		})
	}

	Convey("PROVISIONING - existed", t, func() {
		Convey("should return 200 OK when instance Id exists with identical properties", func() {
			tempBody := openapi.ServiceInstanceProvisionRequest{}
			deepCopy(req, &tempBody)
			_, resp, err := cli.ServiceInstancesApi.ServiceInstanceProvision(
				authCtx, CONF.APIVersion, instanceID, tempBody,
				&openapi.ServiceInstanceProvisionOpts{AcceptsIncomplete: optional.NewBool(async)})

			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 200)
		})
	})
}
