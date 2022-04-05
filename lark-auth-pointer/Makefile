Build:
	GOARCH=amd64 GOOS=linux go build -o lark-auth main.go
	zip lark-auth.zip lark-auth
	aws s3 cp lark-auth.zip s3://zzhe-cn-bucket-1/ --profile cn --region cn-north-1