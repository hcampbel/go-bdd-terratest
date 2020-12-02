terraform {
  # This module is now only being tested with Terraform 0.13.x. However, to make upgrading easier, we are setting
  # 0.12.26 as the minimum version, as that version added support for required_providers with source URLs, making it
  # forwards compatible with 0.13.x code.
  required_version = ">= 0.12.26"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }

}

provider "aws" {
  region                  = var.region
  shared_credentials_file = var.creds
}


//resource "aws_iam_user_policy" "tester_policy" {
//  name        = var.policy_name
//  user        = var.iam_name
//
//  policy = file("policy.json")
//}
//
//resource "aws_iam_user" "tester" {
//  name = var.iam_name
//  path = "/system/"
//}
//
//resource "aws_iam_access_key" "tester" {
//  user = aws_iam_user.tester.name
//}

resource "aws_s3_bucket" "test-bucket" {
  bucket = var.bucket_name
  acl = "private"

  tags = {
    Name = "Test-Bucket"
  }
}

