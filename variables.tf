variable "creds" {
  type = string
}

variable "region" {
  type    = string
  default = "us-east-1"
}

variable "iam_name" {
  type    = string
  default = "tester"
}

variable "policy_name" {
  type = string
}

variable "bucket_name" {
  type = string
}

variable "key" {
  type    = string
  default = "key"
}