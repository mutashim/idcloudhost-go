package idcloudhost_test

import (
	"fmt"
	"testing"

	"github.com/mutashim/idcloudhost-go"
	"github.com/stretchr/testify/assert"
)

var Config idcloudhost.Config = idcloudhost.Config{
	ApiKey: "SopOdexwpt62GRvHJp9R4NzCKBYf0cW2",
}

func TestGetUserInfo(t *testing.T) {
	idc, err := idcloudhost.New(&Config)

	fmt.Println(err)

	fmt.Println(idc.GetUserInfo())
}

func TestModifyProfileInfo(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)

	res, err := idc.ModifyProfileInfo(&idcloudhost.ProfileInput{
		PhoneNumber: "+6281122334455",
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)
}

func TestListToken(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)

	res, err := idc.ListTokens()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)
}

func TestCreateToken(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)

	res, err := idc.CreateToken(&idcloudhost.TokenInput{
		Description:      "Second",
		Restricted:       false,
		BillingAccountID: 0,
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)
}

func TestUpdateToken(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)
	res, err := idc.UpdateToken(42812, &idcloudhost.TokenInput{
		Description: "Diubah",
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)
}

func TestDeleteToken(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)
	err := idc.DeleteToken(42814)
	assert.Nil(t, err)
}

func TestLocation(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)
	res, err := idc.ListLocations()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)

}

func TestListVM(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)
	res, err := idc.ListVM()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)
}

func TestCreateVM(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)
	res, err := idc.CreateVM(&idcloudhost.VMInput{
		BillingAccountID: 1200152307,
		Backup:           false,
		Description:      "Test Create VM",
		Name:             "Hollys",
		OSName:           "ubuntu",
		OSVersion:        "22.04-lts",
		ReservePublicIP:  true,
		Password:         "HeALA991",
		Username:         "ubuntu",
		RAM:              1024,
		VCPU:             1,
		Disks:            "20",
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)
}

func TestStopVM(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)
	res, err := idc.StopVM("26e2a83e-162e-4d95-9b94-135efd3c7b97")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)

}

func TestGetBillingAccountList(t *testing.T) {
	idc, _ := idcloudhost.New(&Config)

	acc, err := idc.ListBillingAccounts()
	fmt.Println(err)
	fmt.Println(acc)
}

func TestGetBuckets(t *testing.T) {
	idc, err := idcloudhost.New(&Config)

	fmt.Println(err)
	fmt.Println(idc.GetS3APIInfo())
}

func TestGetParam(t *testing.T) {
	c, _ := idcloudhost.New(&Config)

	fmt.Println(c.GetVMParameter())
}

func TestGetBillingDetails(t *testing.T) {
	// idc, err := idcloudhost.New(&Config)

	// idc.
}
