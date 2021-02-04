module model-service

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.1.0
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210125195502-f46fe6c6624a
	google.golang.org/grpc v1.35.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.2.8
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
	k8s.io/apimachinery v0.20.0
	k8s.io/client-go v0.20.0
)
