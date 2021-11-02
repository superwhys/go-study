module go-study

go 1.16

require (
	gitee.com/superwhys/go-study/v_protobuf/person v0.0.0
	gitee.com/superwhys/go-study/w_grpc/pb v0.0.0
	github.com/gogo/protobuf v1.3.2
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/net v0.0.0-20211008194852-3b03d305991f
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211008145708-270636b82663 // indirect
	google.golang.org/grpc v1.41.0
)

replace (
	gitee.com/superwhys/go-study/v_protobuf/person => ./v_protobuf/person
	gitee.com/superwhys/go-study/w_grpc/pb => ./w_grpc/pb
)
