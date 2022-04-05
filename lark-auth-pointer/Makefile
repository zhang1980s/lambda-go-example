Build:
	GOARCH=amd64 GOOS=linux go build -o lark-auth-pointer main.go
	zip lark-auth-pointer.zip lark-auth-pointer
	aws s3 cp lark-auth-pointer.zip s3://zzhe-cn-bucket-1/ --profile cn --region cn-north-1
