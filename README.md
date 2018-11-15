# Fridge Temperatures

This project is a demo of an API taking temperatures from fridge sensors and outputing some statistics around those figures. It is entirely written in Golang and deployable on AWS Lambda and API Gateway

An example of input is:
```json
[{"id": "a","timestamp": 1509493641,"temperature": 3.53},
{"id": "b","timestamp": 1509493642,"temperature": 4.13},
{"id": "c","timestamp": 1509493643,"temperature": 3.96},
{"id": "a","timestamp": 1509493644,"temperature": 3.63},
{"id": "c","timestamp": 1509493645,"temperature": 3.96},
{"id": "a","timestamp": 1509493645,"temperature": 4.63},
{"id": "a","timestamp": 1509493646,"temperature": 3.53},
{"id": "b","timestamp": 1509493647,"temperature": 4.15},
{"id": "c","timestamp": 1509493655,"temperature": 3.95},
{"id": "a","timestamp": 1509493677,"temperature": 3.66},
{"id": "b","timestamp": 1510113646,"temperature": 4.15},
{"id": "c","timestamp": 1510127886,"temperature": 3.36},
{"id": "c","timestamp": 1510127892,"temperature": 3.36},
{"id": "a","timestamp": 1510128112,"temperature": 3.67},
{"id": "b","timestamp": 1510128115,"temperature": 3.88}]
```

The output would then be:
```json
[{"id":"c","average":3.72,"median":3.95,"mode":[3.36,3.96]},
{"id":"a","average":3.78,"median":3.65,"mode":[3.53]},
{"id":"b","average":4.08,"median":4.14,"mode":[4.15]}]
```

## Development && Testing

### Requirements

Working on this project requires certain dependancies:
- Golang [https://golang.org/dl/]
- AWS CLI (for testing) [https://aws.amazon.com/cli/]
- AWS SAM (for testing) [https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html]
- Docker (for testing) [https://www.docker.com/get-started]

Not much really.

### Unit testing & Code Coverage

To run the unit test, all you have to do is running the following:
```bash
make test
```
Which will look for all `*_test.go` files and run them

You can also get a code coverage by running:
```bash
make cover
```
Which includes the test run everytime

### Local simulation

Finally, you can even invoke the function locally by running:
```bash
make local-invoke
```
This will:
1. build the binaries for a Lambda environment
2. start a Lambda Docker container locally
3. and finally run the binary on the container, using the content of the file `event.json` as a payload to the handler.

There should be a nice output similar to what can be seen on CloudWatch:
```
START RequestId: c381418d-f241-15b0-bc18-3ca13bbaf716 Version: $LATEST
END RequestId: c381418d-f241-15b0-bc18-3ca13bbaf716
REPORT RequestId: c381418d-f241-15b0-bc18-3ca13bbaf716	Duration: 1.72 ms	Billed Duration: 100 ms	Memory Size: 128 MB	Max Memory Used: 5 MB
{"statusCode":200,"headers":null,"body":"[{\"id\":\"a\",\"average\":3.78,\"median\":3.65,\"mode\":[3.53]},{\"id\":\"b\",\"average\":4.08,\"median\":4.14,\"mode\":[4.15]},{\"id\":\"c\",\"average\":3.72,\"median\":3.95,\"mode\":[3.36,3.96]}]"}
```

## Deployment

### CircleCI

```
       +----------+
       |  GitHub  |
       +----+-----+
            |
            |  Commit
            |
     +------v------+
   +-+  CircleCI   +------+
   | +-------------+      |
   |                      |
   |  Upload bins      Create/Update resources
   |                      |
   |                      |
+--v---+        +---------v-------+
|  S3  +--------> CloudFormation  |
+------+        +--+------------+-+
                   |            |
           API Endpoint         | Binary
                   |            |
           +-------v-----+    +-v------+
           | API Gateway |    | Lambda |
           +-------------+    +--------+
```

This project is configured to run on CircleCI [https://circleci.com/]. All you need to do if you have such an environment setup is making sure that
1. you have your AWS credentials setup [https://circleci.com/docs/2.0/deployment-integrations/#aws] and the env
2. you have the environment variable `AWS_BUCKET` added to the project, which would be the name of the bucket to upload the project artifact to (Go binaries)

### Locally

The project can also be deployed from locally by simply running the command
```bash
AWS_BUCKET=my_bucket make publish
```

You will need to have your access to AWS configure on your machine for that to succeed [https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html]

### Time to run it!

Once deployed, you should be able to grab the url of the API from API Gateway, stage `Latest`, and then run a following request or through Postman:

```bash
curl -X POST \
  https://n4algq1vc1.execute-api.ap-southeast-2.amazonaws.com/Latest/readings \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '[{"id": "a","timestamp": 1509493641,"temperature": 3.53},
{"id": "b","timestamp": 1509493642,"temperature": 4.13},
{"id": "c","timestamp": 1509493643,"temperature": 3.96},
{"id": "a","timestamp": 1509493644,"temperature": 3.63},
{"id": "c","timestamp": 1509493645,"temperature": 3.96},
{"id": "a","timestamp": 1509493645,"temperature": 4.63},
{"id": "a","timestamp": 1509493646,"temperature": 3.53},
{"id": "b","timestamp": 1509493647,"temperature": 4.15},
{"id": "c","timestamp": 1509493655,"temperature": 3.95},
{"id": "a","timestamp": 1509493677,"temperature": 3.66},
{"id": "b","timestamp": 1510113646,"temperature": 4.15},
{"id": "c","timestamp": 1510127886,"temperature": 3.36},
{"id": "c","timestamp": 1510127892,"temperature": 3.36},
{"id": "a","timestamp": 1510128112,"temperature": 3.67},
{"id": "b","timestamp": 1510128115,"temperature": 3.88}]'
```

## Improvements

As you know, a project is never done! It can always be improved. At least the good news is that there's CI/CD setup for this project, so it's a bit easier. The things that could be improved to make it even nicer are:

### Safe Deployment

Using AWS SAM has this huge advantage of `safe deploying` the lambdas, just because unit tests are never enough. The principle is simple to understand:
1. Deploy the Lamda without moving the `latest` alias to this newest version yet
2. Run a pre traffic function that would test the Lambda function, directly in AWS
3. if the pre traffic hook gives the green light, then move the `latest` alias to the final version, which means traffic will start coming through the new code
4. if it's a red light, then don't do anything, the new code will be technically deployed, but API Gateway will still route the traffic to the previous version.

CloudFormation makes a safe deployment approach really easy to do thanks to SAM's template which involves Code Deploy

### Logging

The current code has absolutely no logging. Yes of course, if any error happens, it'll be printed nicely on CloudWatch. But it won't be any useful to see `invalid character` without any other information. We would need more debug or trace infomation for that.

A very nice and easy library that we could use is [https://github.com/sirupsen/logrus].

Not only it becomes easy to log while adding some context to the messages, rather than just `fmt.Println`'ing everything, but log messages can also be JSON formatted, which then becomes handy when using a proper log manager such as Sumo.

On top of the enhanced logging, and especially if this microservice starts having more than one function processing the same event at different stage, we could add a `correlation ID` to improve the tracability of each processing.

### Security

`https://n4algq1vc1.execute-api.ap-southeast-2.amazonaws.com/Latest/readings` is actually currently live as of 16 Nov 2018. Nothing prevents you from hammering the API. And that is because there is absolutely no security. A simple API key would do the job. We can even make it required to have the payload signed, with the signature sent as a header. Plenty of solutions here.

### CircleCI Docker image

I'm currently using a CircleCI prebuild Docker image for Golang. However this image does not have AWS CLI installed, which is why is in the CircleCI configs, making the build a little bit more time consuming. That could be easily improved by making a new image with this dependancy installed.
