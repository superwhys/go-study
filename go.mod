module go-study

go 1.16

require (
	gitee.com/superwhys/go-study/v_protobuf/person v0.0.0 // indirect
	gitee.com/superwhys/go-study/w_grpc/pb v0.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20211008194852-3b03d305991f // indirect
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211008145708-270636b82663 // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace (
	gitee.com/superwhys/go-study/w_grpc/pb => ./w_grpc/pb
	gitee.com/superwhys/go-study/v_protobuf/person => ./v_protobuf/person
)
