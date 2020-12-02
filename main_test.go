package go_bdd_terratest

import (
	"errors"
	"fmt"
	"github.com/go-bdd/assert"
	"github.com/gruntwork-io/terratest/modules/aws"
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

var bucketName string
var region string


func setS3BucketContext(t gobdd.StepTest, ctx gobdd.Context, name string) {

	if name != bucketName {
		t.Error(errors.New("the names don't match"))
		return
	}
	fmt.Println("Setting the Context")
	ctx.Set("bucket", name)
}

func verifyS3BucketExists(t gobdd.StepTest, ctx gobdd.Context, name string) {

	fmt.Println("Performing verification and validation of the S3 bucket")
	received, err := ctx.GetString("bucket")

	if err != nil {
		t.Error(err)
		return
	}

	err = assert.Equals(received, name)

	if err != nil {
		t.Error(errors.New("the expected and the actual names of the bucket do not match"))
	}


}

func TestScenarios(t *testing.T) {
	suite := gobdd.NewSuite(t)
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "/home/hugh/Documents/Programming/Golang/github.com/hcampbel/go-bdd-terratest/terraform",
	})

	region = "us-west-2"

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketName = terraform.Output(t, terraformOptions, "bucket_name")


	suite.AddStep(`I create new bucket (\w+\W+\w+\W+\w+\W+\w+)`, setS3BucketContext)
	aws.AssertS3BucketExists(t, region, bucketName)
	suite.AddStep(`the (\w+\W+\w+\W+\w+\W+\w+) creation succeeded`, verifyS3BucketExists)
	suite.Run()

}
