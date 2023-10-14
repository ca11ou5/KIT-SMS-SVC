proto:
	protoc ./internal/pb/sms.proto --go_out=plugins=grpc:.

buildimage:
	docker build -t sms-svc .