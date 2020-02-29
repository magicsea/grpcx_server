package share

import (
	"strconv"
	"google.golang.org/grpc/metadata"
	"context"
	"errors"
)
const (
	UserID_Key string = "uid"
)
func GetUIDFromContext(ctx context.Context) (int64,error) {
	oc,_:= metadata.FromIncomingContext(ctx)
	list, ok := oc[UserID_Key]
	if !ok {
		return 0,errors.New("not exist key")
	}
	v,err := strconv.Atoi(list[0])
	return int64(v),err
}


