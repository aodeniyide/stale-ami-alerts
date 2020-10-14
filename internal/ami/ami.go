package ami

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"time"
)

// QueryAmi get image descriptions.
func QueryAmi (ownerid string, region string ) map[string]string {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	svc := ec2.New(sess)
	input := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(ownerid),
		},
	}

	result, err := svc.DescribeImages(input)
	if err != nil {
		fmt.Printf("Unable to describe images, %v\n", err)
	}

	img := make(map[string]string)
	for _, j := range result.Images {
		img[*j.ImageId] = *j.CreationDate
	}
	//fmt.Println("map:", img)
	return img
}

// convert aws timestamp
func UpdateTime(t string) time.Time {
	f, err :=  time.Parse(time.RFC3339, t)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

// process the aws timestamp of ami to get days
func ProcessStaleAmi(t time.Time) int {
	createTime := t
	currentTime := time.Now()
	delta := currentTime.Sub(createTime).Hours() / 24
	return int(delta)

}

// Create the alert
func AlertStaleAmi(a string, t int, dayLimit int)  {
	if t >= dayLimit {
		fmt.Println("ami_id:", a, "has an uptime of:", t, "days. Please deregister")
	}
}
