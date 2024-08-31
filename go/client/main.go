package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2/google"
	// "google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	// "google.golang.org/api/option"
	"golang.org/x/oauth2"

	pb "github.com/williamjiuchengcai/medlmpp-incubation-platform-jcwc-dev/helloworld"
	grpcMetadata "google.golang.org/grpc/metadata"
)

// NewConn creates a new gRPC connection.
// host should be of the form domain:port, e.g., example.com:443
func NewConn(host string, insecure bool) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}

	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		// Note: On the Windows platform, use of x509.SystemCertPool() requires
		// go version 1.18 or higher.
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	return grpc.Dial(host, opts...)
}

// pingRequestWithAuth mints a new Identity Token for each request.
// This token has a 1 hour expiry and should be reused.
// audience must be the auto-assigned URL of a Cloud Run service or HTTP Cloud Function without port number.
func pingRequestWithAuth(conn *grpc.ClientConn, p *pb.HelloRequest, audience string) (*pb.HelloReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Construct the GoogleCredentials object which obtains the default configuration from your
	// working environment.
	// credentials, err := google.FindDefaultCredentials(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to generate default credentials: %w", err)
	// }

	// If that fails, we use our Application Default Credentials to fetch an id_token on the fly
	gts, err := google.DefaultTokenSource(ctx)
	if err != nil {
		return nil, err
	}
	tokenSource := oauth2.ReuseTokenSource(nil, &idTokenSource{TokenSource: gts})

	// Create an identity token.
	// With a global TokenSource tokens would be reused and auto-refreshed at need.
	// A given TokenSource is specific to the audience.
	// tokenSource, err := idtoken.NewTokenSource(ctx, audience, option.WithCredentials(credentials))
	// if err != nil {
	// 	return nil, fmt.Errorf("idtoken.NewTokenSource: %w", err)
	// }
	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("TokenSource.Token: %w", err)
	}

	// Add token to gRPC Request.
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken)

	// Send the request.
	client := pb.NewGreeterClient(conn)
	return client.SayHello(ctx, p)
}

// idTokenSource is an oauth2.TokenSource that wraps another
// It takes the id_token from TokenSource and passes that on as a bearer token
type idTokenSource struct {
	TokenSource oauth2.TokenSource
}

func (s *idTokenSource) Token() (*oauth2.Token, error) {
	token, err := s.TokenSource.Token()
	if err != nil {
		return nil, err
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("token did not contain an id_token")
	}

	return &oauth2.Token{
		AccessToken: idToken,
		TokenType:   "Bearer",
		Expiry:      token.Expiry,
	}, nil
}

func main() {
	conn, err := NewConn("helloworld-836582501653.us-central1.run.app:443", false)
	if err != nil {
		log.Fatalf("Could not create new connection: %w", err)
	}
	reply, err := pingRequestWithAuth(conn, &pb.HelloRequest{
		Name: "William Cai",
	}, "https://helloworld-836582501653.us-central1.run.app")
	if err != nil {
		log.Fatalf("HelloReply returned error: %w", err)
	}
	log.Printf("Successful reply: %v", reply)
}
