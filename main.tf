provider "aws" {
  region = "us-west-2"
}


resource "aws_s3_bucket" "example" {
  bucket = "demo-terraform-eks-state-bucket-20"
  lifecycle {
    prevent_destroy = false
  }
}

resource "aws_dynamodb_table" "basic-dynamodb-table" {
  name           = "yomi-demo-eks-state-locks"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "LockID"


  attribute {
    name = "LockID"
    type = "S"
  }


}