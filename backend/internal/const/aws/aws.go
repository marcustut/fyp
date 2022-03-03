package aws

// BucketName is the S3 Bucket's ID.
const BucketName string = "ai-text-summarizer"

// AmiUbuntux86_64 is the AMI Image ID for
// Ubuntu 20.04 on x86_64.
const AmiUbuntux86_64 string = "ami-055d15d9cfddf7bd3"

// AmiAmazonLinux2x86_64 is the AMI Image ID for
// Amazon Linux 2 on x86_64.
const AmiAmazonLinux2x86_64 string = "ami-02a45d709a415958a"

// EC2KeyPairName is the default key pair
// used to connect to EC2 instances.
const EC2KeyPairName string = "FYP"

// EC2InstanceProfileArn is the ARN for
// IAM Role of "fyp-ec2-role".
const EC2InstanceProfileArn string = "arn:aws:iam::607729451680:instance-profile/fyp-ec2-role"
