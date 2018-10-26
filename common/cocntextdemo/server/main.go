package main

import (
	"net/http"
	"math/rand"
	"time"
	"fmt"
	"context"
)

const requestIDKey = "rid"

func lazyHandler(w http.ResponseWriter, r *http.Request) {
	ranNum := rand.Intn(2)
	if ranNum == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "slow response, %d\n", ranNum)
		fmt.Printf("slow response, %d\n", ranNum)
		return
	}
	fmt.Fprintf(w, "quick response, %d\n", ranNum)
	fmt.Printf("quick response, %d\n", ranNum)
	return
}

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	reqID := req.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = "0"
	}
	return context.WithValue(ctx, requestIDKey, reqID)
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := newContextWithRequestID(req.Context(), req)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func h(w http.ResponseWriter, req *http.Request)  {
	reqID := requestIDFromContext(req.Context())
	fmt.Fprintln(w, "Request ID: ", reqID)
	return
}

func main() {
	http.HandleFunc("/lazy", lazyHandler)
	http.Handle("/middle", middleWare(http.HandlerFunc(h)))
	http.ListenAndServe(":9201", nil)
}
