package util

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func CustomHttpError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Internal, err.Error())
	}

	httpStatus := runtime.HTTPStatusFromCode(s.Code())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	resp := map[string]interface{}{
		"status": httpStatus,
		"error":  s.Message(),
	}
	json.NewEncoder(w).Encode(resp)
}
