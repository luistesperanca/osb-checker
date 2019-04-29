package client

import (
	"encoding/json"
	"strconv"

	v2 "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
	. "github.com/openservicebrokerapi/osb-checker/config"
)

var Default = NewClient(NewReceiver())

func NewClient(r Receiver) *Client {
	return &Client{
		Receiver: r,
	}
}

type Client struct {
	Receiver
}

func (c *Client) GetCatalog() (int, *v2.Catalog, error) {
	var catalog v2.Catalog
	params := &BrokerRequestParams{
		URL:    GenerateCatalogURL(),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &catalog)

	return res.Code, &catalog, nil
}

func (c *Client) Provision(
	instanceID string,
	req *v2.ServiceInstanceProvisionRequest,
	async bool,
) (int, *v2.ServiceInstanceProvision, *v2.ServiceInstanceAsyncOperation, error) {

	var provision v2.ServiceInstanceProvision
	var asyncOperation v2.ServiceInstanceAsyncOperation
	params := &BrokerRequestParams{
		URL: GenerateInstanceURL(instanceID) +
			"?accepts_incomplete=" + strconv.FormatBool(async),
		Method: "PUT",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username:  CONF.Authentication.Username,
		Password:  CONF.Authentication.Password,
		InputData: req,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, nil, err
	}
	switch res.Code {
	case 200, 201:
		json.Unmarshal(res.Message, &provision)
		break
	case 202:
		json.Unmarshal(res.Message, &asyncOperation)
		break
	default:
		break
	}

	return res.Code, &provision, &asyncOperation, nil
}

func (c *Client) PollInstanceLastOperation(
	instanceID string,
) (int, *v2.LastOperationResource, error) {

	var lastOperation v2.LastOperationResource
	params := &BrokerRequestParams{
		URL:    GeneratePollInstanceLastOperationURL(instanceID),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &lastOperation)

	return res.Code, &lastOperation, nil
}

func (c *Client) GetInstance(
	instanceID string,
) (int, *v2.ServiceInstanceResource, error) {

	var instance v2.ServiceInstanceResource
	params := &BrokerRequestParams{
		URL:    GenerateInstanceURL(instanceID),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &instance)

	return res.Code, &instance, nil
}

func (c *Client) UpdateInstance(
	instanceID string,
	req *v2.ServiceInstanceUpdateRequest,
	async bool,
) (int, *v2.ServiceInstanceAsyncOperation, error) {

	var asyncOperation v2.ServiceInstanceAsyncOperation
	params := &BrokerRequestParams{
		URL: GenerateInstanceURL(instanceID) +
			"?accepts_incomplete=" + strconv.FormatBool(async),
		Method: "PATCH",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username:  CONF.Authentication.Username,
		Password:  CONF.Authentication.Password,
		InputData: req,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	switch res.Code {
	case 202:
		json.Unmarshal(res.Message, &asyncOperation)
		break
	default:
		break
	}

	return res.Code, &asyncOperation, nil
}

func (c *Client) Deprovision(
	instanceID string,
	serviceID, planID string,
	async bool,
) (int, *v2.AsyncOperation, error) {

	var asyncOperation v2.AsyncOperation
	params := &BrokerRequestParams{
		URL: GenerateInstanceURL(instanceID) +
			"?accepts_incomplete=" + strconv.FormatBool(async) +
			"&service_id=" + serviceID + "&plan_id=" + planID,
		Method: "DELETE",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	switch res.Code {
	case 202:
		json.Unmarshal(res.Message, &asyncOperation)
		break
	default:
		break
	}

	return res.Code, &asyncOperation, nil
}

func (c *Client) Bind(
	instanceID, bindingID string,
	req *v2.ServiceBindingRequest,
	async bool,
) (int, *v2.ServiceBinding, *v2.AsyncOperation, error) {

	var binding v2.ServiceBinding
	var asyncOperation v2.AsyncOperation
	params := &BrokerRequestParams{
		URL: GenerateBindingURL(instanceID, bindingID) +
			"?accepts_incomplete=" + strconv.FormatBool(async),
		Method: "PUT",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username:  CONF.Authentication.Username,
		Password:  CONF.Authentication.Password,
		InputData: req,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, nil, err
	}
	switch res.Code {
	case 200, 201:
		json.Unmarshal(res.Message, &binding)
		break
	case 202:
		json.Unmarshal(res.Message, &asyncOperation)
		break
	default:
		break
	}

	return res.Code, &binding, &asyncOperation, nil
}

func (c *Client) PollBindingLastOperation(
	instanceID, bindingID string,
) (int, *v2.LastOperationResource, error) {

	var lastOperation v2.LastOperationResource
	params := &BrokerRequestParams{
		URL:    GeneratePollBindingLastOperationURL(instanceID, bindingID),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &lastOperation)

	return res.Code, &lastOperation, nil
}

func (c *Client) GetBinding(
	instanceID, bindingID string,
) (int, *v2.ServiceBindingResource, error) {

	var binding v2.ServiceBindingResource
	params := &BrokerRequestParams{
		URL:    GenerateBindingURL(instanceID, bindingID),
		Method: "GET",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	json.Unmarshal(res.Message, &binding)

	return res.Code, &binding, nil
}

func (c *Client) Unbind(
	instanceID, bindingID,
	serviceID, planID string,
	async bool,
) (int, *v2.AsyncOperation, error) {

	var asyncOperation v2.AsyncOperation
	params := &BrokerRequestParams{
		URL: GenerateBindingURL(instanceID, bindingID) +
			"?accepts_incomplete=" + strconv.FormatBool(async) +
			"&service_id=" + serviceID + "&plan_id=" + planID,
		Method: "DELETE",
		HeaderOption: map[string]string{
			"X-Broker-API-Version": CONF.APIVersion,
		},
		Username: CONF.Authentication.Username,
		Password: CONF.Authentication.Password,
	}

	res, err := c.Recv(params)
	if err != nil {
		return -1, nil, err
	}
	switch res.Code {
	case 202:
		json.Unmarshal(res.Message, &asyncOperation)
		break
	default:
		break
	}

	return res.Code, &asyncOperation, nil
}
