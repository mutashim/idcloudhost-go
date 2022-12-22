package idcloudhost

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	apikey string
	http   *http.Client
}

// Make a request to API
func (c *Client) request(method, path string, data *url.Values) ([]byte, error) {

	var body io.Reader = nil

	if data != nil {
		log.Println(data.Encode())
		body = strings.NewReader(data.Encode())
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%s/%s", BASEAPI, path), body)
	if err != nil {
		return nil, fmt.Errorf("client: could not create request: %s", err)
	}

	request.Header.Add("apikey", c.apikey)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := c.http.Do(request)
	if err != nil {
		return nil, fmt.Errorf("client: error making http request: %s", err)
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read response body: %s", err)
	}
	log.Println(string(responseBody[:]))

	if response.StatusCode != 200 {
		// var errorBody struct {
		// errors []string `json:"errors"`
		// }
		return nil, fmt.Errorf("client: error response: %s:%s", response.Status, response.Body)
	}

	return responseBody, nil
}

// Authenticated user can request data-model for themselves
// https://api.idcloudhost.com/#get-user-info
func (c *Client) GetUserInfo() (*User, error) {
	res, err := c.request("GET", "v1/user-resource/user", nil)
	if err != nil {
		return nil, err
	}

	userinfo := User{}
	err = json.Unmarshal(res, &userinfo)
	if err != nil {
		return nil, err
	}

	return &userinfo, nil
}

// Authenticated user can modify their own profile data.
// https://api.idcloudhost.com/#modify-profile-info
func (c *Client) ModifyProfileInfo(p *ProfileInput) (*Profile, error) {
	data, err := parseForm(p)
	if err != nil {
		return nil, err
	}

	res, err := c.request("PATCH", "v1/user-resource/user/profile", data)
	if err != nil {
		return nil, err
	}

	profile := Profile{}
	err = json.Unmarshal(res, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// Delete API token.
// https://api.idcloudhost.com/#delete-token
func (c *Client) DeleteToken(tokenID int64) error {
	data := url.Values{}
	data.Add("token_id", fmt.Sprint(tokenID))

	_, err := c.request("DELETE", "v1/user-resource/token", &data)

	return err
}

// Create new token and register it at API Gateway.
// https://api.idcloudhost.com/#create-token
func (c *Client) CreateToken(in *TokenInput) (*Token, error) {
	data, err := parseForm(in)
	if err != nil {
		return nil, err
	}

	result, err := c.request("POST", "v1/user-resource/token", data)
	if err != nil {
		return nil, err
	}

	token := Token{}
	err = json.Unmarshal(result, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// List user API tokens.
// https://api.idcloudhost.com/#list-tokens
func (c *Client) ListTokens() (*[]Token, error) {
	result, err := c.request("GET", "v1/user-resource/token/list", nil)
	if err != nil {
		return nil, err
	}

	tokens := []Token{}
	err = json.Unmarshal(result, &tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

// Update API token options.
// https://api.idcloudhost.com/#update-token
func (c *Client) UpdateToken(ID int64, t *TokenInput) (*Token, error) {

	data, err := parseForm(t)
	if err != nil {
		return nil, err
	}

	data.Add("token_id", fmt.Sprint(ID))

	res, err := c.request("PATCH", "v1/user-resource/token", data)
	if err != nil {
		return nil, err
	}

	token := Token{}
	if err := json.Unmarshal(res, &token); err != nil {
		return nil, err
	}

	return &token, err
}

// Locations are different data centres or resource pools where
// virtual machines can be hosted. By default API calls manage
// resources in the "is_default": true location.
// To access a specific location, its slug must be used in the
// API URL right after version number : v1/{slug}/. In the
// following example Cycletown location is used if no location
// is specified.
func (c *Client) ListLocations() (*[]Location, error) {
	res, err := c.request("GET", "v1/config/locations", nil)
	if err != nil {
		return nil, err
	}

	locs := []Location{}
	err = json.Unmarshal(res, &locs)
	if err != nil {
		return nil, err
	}

	return &locs, nil
}

// Get virtual machine list
func (c *Client) ListVM() (*[]VM, error) {
	res, err := c.request("GET", "v1/user-resource/vm/list", nil)
	if err != nil {
		return nil, err
	}

	vms := []VM{}
	if err := json.Unmarshal(res, &vms); err != nil {
		return nil, err
	}

	return &vms, nil
}

// Delete virtual machine
// https://api.idcloudhost.com/#delete-vm
func (c *Client) DeleteVM(uuid string) error {
	data := url.Values{}
	data.Add("uuid", uuid)

	_, err := c.request("DELETE", "v1/user-resource/vm", &data)
	if err != nil {
		return err
	}

	return nil
}

// https://api.idcloudhost.com/#get-vm-info
func (c *Client) GetVMInfo(uuid string) (*VM, error) {
	res, err := c.request("GET", fmt.Sprintf("%s?uuid=%s", "v1/user-resource/vm", uuid), nil)
	if err != nil {
		return nil, err
	}

	vm := VM{}
	if err := json.Unmarshal(res, &vm); err != nil {
		return nil, err
	}

	return &vm, nil
}

func (c *Client) StopVM(uuid string) (*VM, error) {
	data := url.Values{}
	data.Add("uuid", uuid)
	res, err := c.request("POST", "v1/user-resource/vm/stop", &data)
	if err != nil {
		return nil, err
	}

	vm := VM{}
	if err := json.Unmarshal(res, &vm); err != nil {
		return nil, err
	}

	return &vm, nil
}

// Set a new password for an existing user on the virtual machine.
// The VM must be running, otherwise the password cannot be changed
// and an error will be returned.
// https://api.idcloudhost.com/#change-password
func (c *Client) ChangePassword(uuid, username, password string) error {
	data := url.Values{}
	data.Add("uuid", uuid)
	data.Add("username", username)
	data.Add("password", password)

	res, err := c.request("PATCH", "v1/user-resource/vm/user", &data)
	if err != nil {
		return err
	}

	var response struct {
		Success bool `json:"success"`
	}

	if err := json.Unmarshal(res, &response); err != nil {
		return err
	}

	if !response.Success {
		return errors.New("error change password")
	}

	return nil
}

// Currently only name, vcpu and ram can be changed. Do note that vcpu
// and ram can only be changed when the machine is in stopped state.
// https://api.idcloudhost.com/#modify-vm
func (c *Client) ModifyVM(uuid, name string, ram, vcpu int64) (*VM, error) {
	data := url.Values{}
	data.Add("uuid", uuid)
	data.Add("name", name)
	data.Add("ram", fmt.Sprint(ram))
	data.Add("vcpu", fmt.Sprint(vcpu))

	res, err := c.request("PATCH", "v1/user-resource/vm", &data)
	if err != nil {
		return nil, err
	}

	vm := VM{}
	if err := json.Unmarshal(res, &vm); err != nil {
		return nil, err
	}

	return &vm, nil
}

// Create a new virtual machine. billing_account_id is optional if using
// an API token that is restricted to one billing account.
//
// Specify source_uuid (VM UUID) and source_replica (snapshot or backup
// UUID) to create the new virtual machine as a copy of an existing backup
// or snapshot.
//
// Specify disk_uuid to use an existing unattached disk as the boot disk
// or the new VM. In this case disks parameter has no effect and should
// be left empty.
//
// If a VM without public IPv4 address is needed, set reserve_public_ip
// to False. Specify network_uuid to create the VM in the given network.
// If the value is emtpy the VM is created in the default network.
//
// VMs are initialised with cloud-init, it is possible to add custom
// cloud-init configuration to cloud_init parameter as JSON, currently
// runcmd and write_files are supported. See cloud-init documentation for
// more information.
//
// https://api.idcloudhost.com/#create-vm
func (c *Client) CreateVM(vm *VMInput) (*VM, error) {

	data, err := parseForm(vm)
	if err != nil {
		return nil, err
	}

	res, err := c.request("POST", "v1/user-resource/vm", data)
	if err != nil {
		return nil, err
	}

	response := VM{}
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Describes VM creation parameters and their expected and allowed values.
// https://api.idcloudhost.com/#vm-parameters
func (c *Client) GetVMParameter() (*VMParameter, error) {

	res, err := c.request("GET", "v1/api/parameters/vm", nil)
	if err != nil {
		return nil, err
	}

	response := []ParameterItem{}
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	final := VMParameter{}

	for _, v := range response {

		switch v.Parameter {
		case "os_version":
			os := []VMOS{}
			err := json.Unmarshal(v.Limits, &os)
			if err != nil {
				fmt.Println(err)
			}
			final.OS = os
			fmt.Println(os)
		case "disks":
			disk := VMRange{
				Min: v.Min,
				Max: v.Max,
			}
			final.Disks = disk
		}
	}
	return &final, nil
}

// S3 API Info returns S3 API URL.
// https://api.idcloudhost.com/#s3-api-info
func (c *Client) GetS3APIInfo() (*S3APIInfo, error) {
	res, err := c.request("GET", "v1/storage/api/s3", nil)
	if err != nil {
		return nil, err
	}

	s3info := S3APIInfo{}
	err = json.Unmarshal(res, &s3info)
	if err != nil {
		return nil, err
	}

	return &s3info, nil
}

// Create an S3 object storage bucket. Bucket names must be globally unique across all users.
// https://api.idcloudhost.com/#create-bucket
func (c *Client) CreateBucket(name string, billingAccountID int64) (*Bucket, error) {

	data := url.Values{}
	data.Add("name", name)
	if billingAccountID != 0 {
		data.Add("billing_account_id", fmt.Sprint(billingAccountID))
	}

	res, err := c.request("POST", "v1/storage/bucket", &data)
	if err != nil {
		return nil, err
	}

	buck := Bucket{}
	err = json.Unmarshal(res, &buck)
	if err != nil {
		return nil, err
	}

	return &buck, nil
}

// Modify a bucket's billing account
// Info: https://api.idcloudhost.com/#modify-bucket
func (c *Client) ModifyBucket(name string, billingAccountID int64) (*Bucket, error) {

	data := url.Values{}
	data.Add("name", name)
	data.Add("billing_account_id", fmt.Sprint(billingAccountID))

	res, err := c.request("PATCH", "v1/storage/bucket", &data)
	if err != nil {
		return nil, err
	}

	buck := Bucket{}
	err = json.Unmarshal(res, &buck)
	if err != nil {
		return nil, err
	}

	return &buck, nil
}

// Delete an S3 object storage bucket. Only empty buckets can be deleted via this API.
// https://api.idcloudhost.com/#delete-bucket
func (c *Client) DeleteBucket(name string) error {
	data := url.Values{}
	data.Add("name", name)

	_, err := c.request("DELETE", "v1/storage/bucket", &data)
	if err != nil {
		return err
	}

	return nil

}

// Get bucket information.
// https://api.idcloudhost.com/#get-bucket
func (c *Client) GetBucket(BucketName string) (*Bucket, error) {
	res, err := c.request("GET", fmt.Sprintf("v1/storage/bucket?=%s", BucketName), nil)
	if err != nil {
		return nil, err
	}

	buck := Bucket{}
	err = json.Unmarshal(res, &buck)
	if err != nil {
		return nil, err
	}

	return &buck, nil
}

// List user's buckets. Optionally filter the list by billing account.
// https://api.idcloudhost.com/#list-buckets
func (c *Client) ListBucket() (*[]Bucket, error) {
	res, err := c.request("GET", "v1/storage/bucket/list", nil)
	if err != nil {
		return nil, err
	}

	bucks := []Bucket{}
	err = json.Unmarshal(res, &bucks)
	if err != nil {
		return nil, err
	}

	return &bucks, nil
}

// Get S3 user info, including their access and secret keys.
// User and keys will be generated, if they do not exist already.
// https://api.idcloudhost.com/#get-s3-user
func (c *Client) GetS3Users() (*S3User, error) {
	res, err := c.request("GET", "v1/storage/user", nil)
	if err != nil {
		return nil, err
	}

	user := S3User{}
	err = json.Unmarshal(res, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete a S3 key.
// https://api.idcloudhost.com/#delete-key
func (c *Client) DeleteKey(accessKey string) error {
	data := url.Values{}
	data.Add("access_key", accessKey)

	_, err := c.request("DELETE", "v1/storage/user/keys", &data)
	if err != nil {
		return err
	}

	return nil
}

// Returns all user keys.
// https://api.idcloudhost.com/#get-keys
func (c *Client) GetKeys() (*[]S3Credential, error) {
	res, err := c.request("GET", "v1/storage/user/keys", nil)
	if err != nil {
		return nil, err
	}

	keys := []S3Credential{}
	err = json.Unmarshal(res, &keys)
	if err != nil {
		return nil, err
	}

	return &keys, nil
}

// Generate a new S3 key pair. Returns the list of all keys.
// https://api.idcloudhost.com/#generate-key
func (c *Client) GenerateKey() (*[]S3Credential, error) {
	res, err := c.request("POST", "v1/storage/user/keys", nil)
	if err != nil {
		return nil, err
	}

	keys := []S3Credential{}
	err = json.Unmarshal(res, &keys)
	if err != nil {
		return nil, err
	}

	return &keys, nil
}

func (c *Client) ListBillingAccounts() (*[]BillingAccount, error) {
	res, err := c.request("GET", "v1/payment/billing_account/list", nil)
	if err != nil {
		return nil, err
	}

	accs := []BillingAccount{}
	err = json.Unmarshal(res, &accs)
	if err != nil {
		return nil, err
	}

	return &accs, nil
}
