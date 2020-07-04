module github.com/xiaokangwang/V2CloudSign

go 1.14

require (
	github.com/aws/aws-lambda-go v1.17.0
	github.com/xiaokangwang/V2BuildAssist v0.0.0-20200702083956-84eb6c6b7ae3
)

replace (
	github.com/xiaokangwang/V2BuildAssist => ../V2BuildAssist
)