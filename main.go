package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fortanix/sdkms-client-go/sdkms"
	"github.com/fxamacker/cbor/v2"
	"google.golang.org/grpc"
)

const (
	defaultConfigPath = "/etc/fortanix/k8s-sdkms-plugin.json"
	netProtocol       = "unix"
	version           = "v2beta1"
	healthz           = "ok"
	runtimeName       = "k8s-sdkms-plugin"
	runtimeVersion    = "0.3.0"
)

func main() {
	configFile := flag.String("config", defaultConfigPath, "config file location")
	flag.Parse()

	log.Println("Reading config...")
	config, err := readConfigFromFile(*configFile)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	if err := config.validate(); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	log.Println("Starting gRPC service...")
	server, err := startServer(*config)
	if err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

	log.Printf("version: %v, runtime: %v (%v)", version, runtimeName, runtimeVersion)
	log.Println("Service started successfully.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	log.Printf("Signal: '%v', shutting down gRPC service...\n", sig)
	server.server.GracefulStop()
}

type pluginConfig struct {
	SdkmsEndpoint *string `json:"sdkms_endpoint,omitempty"`
	ApiKey        *string `json:"api_key,omitempty"`
	KeyName       *string `json:"key_name,omitempty"`
	KeyID         *string `json:"key_id,omitempty"`
	SocketFile    *string `json:"socket_file,omitempty"`
}

func readConfigFromFile(configFilePath string) (*pluginConfig, error) {
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	var config pluginConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %v: %v", configFilePath, err)
	}
	return &config, nil
}

func (p pluginConfig) validate() error {
	if p.SdkmsEndpoint == nil {
		return errors.New("required field `sdkms_endpoint` is missing")
	}
	if p.ApiKey == nil {
		return errors.New("required field `api_key` is missing")
	}
	if p.SocketFile == nil {
		return errors.New("required field `socket_file` is missing")
	}
	if p.KeyName == nil && p.KeyID == nil {
		return errors.New("neither `key_name` nor `key_id` was specified")
	}
	if p.KeyName != nil && p.KeyID != nil {
		return errors.New("cannot specify `key_name` and `key_id` at the same time")
	}
	// verify configuration by authenticating and getting the encryption key
	ctx := context.Background()
	client := p.makeClient()
	_, err := client.AuthenticateWithAPIKey(ctx, *p.ApiKey)
	if err != nil {
		return fmt.Errorf("invalid `api_key`: %v", err)
	}
	defer client.TerminateSession(ctx)

	descriptor := p.makeSobjectDescriptor()
	encoding := sdkms.SobjectEncodingJson
	key, err := client.GetSobject(ctx, &sdkms.GetSobjectParams{View: &encoding}, *descriptor)
	if err != nil {
		return fmt.Errorf("invalid key: %v", err)
	}
	if key.ObjType != sdkms.ObjectTypeAes {
		return fmt.Errorf("invalid key type, expected AES, found: %v", key.ObjType)
	}
	return nil
}

func (p pluginConfig) makeClient() sdkms.Client {
	client := sdkms.Client{
		HTTPClient: http.DefaultClient,
		Endpoint:   *p.SdkmsEndpoint,
		Auth:       sdkms.APIKey(*p.ApiKey),
	}
	return client
}

func (p pluginConfig) makeSobjectDescriptor() *sdkms.SobjectDescriptor {
	if p.KeyName != nil {
		return sdkms.SobjectByName(*p.KeyName)
	}
	return sdkms.SobjectByID(*p.KeyID)
}

type kmsServer struct {
	server *grpc.Server
	config pluginConfig
	hash   string
}

// Hash of endPoint, KeyID and KeyName
func (p pluginConfig) hash() string {
	h := sha256.New()

	if p.SdkmsEndpoint != nil {
		h.Write([]byte(*p.SdkmsEndpoint))
	}
	if p.KeyID != nil {
		h.Write([]byte(*p.KeyID))
	}
	if p.KeyName != nil {
		h.Write([]byte(*p.KeyName))
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func startServer(config pluginConfig) (*kmsServer, error) {
	if err := os.Remove(*config.SocketFile); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to remove stale socket: %v", err)
	}

	listener, err := net.Listen(netProtocol, *config.SocketFile)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	s := &kmsServer{
		config: config,
		server: server,
		hash:   config.hash(),
	}
	RegisterKeyManagementServiceServer(server, s)
	go server.Serve(listener)
	return s, nil
}

type wrappedData struct {
	Version int
	KID     string
	Cipher  []byte
	IV      []byte
	Tag     []byte
}

func (s *kmsServer) Status(ctx context.Context, request *StatusRequest) (*StatusResponse, error) {
	logRequest("Status", fmt.Sprintf("healtcheck status is %v", healthz), nil)
	return &StatusResponse{Version: version, Healthz: healthz, KeyId: s.hash}, nil
}

func (s *kmsServer) Encrypt(ctx context.Context, request *EncryptRequest) (*EncryptResponse, error) {
	resp, msg, err := s.encrypt(ctx, request)
	logRequest("Encrypt", msg, err)
	return resp, err
}

func (s *kmsServer) Decrypt(ctx context.Context, request *DecryptRequest) (*DecryptResponse, error) {
	resp, msg, err := s.decrypt(ctx, request)
	logRequest("Decrypt", msg, err)
	return resp, err
}

func (s *kmsServer) encrypt(ctx context.Context, request *EncryptRequest) (*EncryptResponse, string, error) {
	client := s.config.makeClient()
	tagLen := uint(128)
	resp, err := client.Encrypt(ctx, sdkms.EncryptRequest{
		Key:    s.config.makeSobjectDescriptor(),
		Alg:    sdkms.AlgorithmAes,
		Plain:  request.Plaintext,
		Mode:   sdkms.CryptModeSymmetric(sdkms.CipherModeGcm),
		TagLen: &tagLen,
	})
	if err != nil {
		return nil, "", err
	}
	data, err := cbor.Marshal(wrappedData{
		Version: 1, // signifies AES GCM without AAD
		KID:     *resp.Kid,
		Cipher:  resp.Cipher,
		IV:      *resp.Iv,
		Tag:     *resp.Tag,
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to serialize encrypt response: %v", err)
	}
	return &EncryptResponse{Ciphertext: data, KeyId: s.hash}, fmt.Sprintf("plain: %v bytes, cipher: %v bytes", len(request.Plaintext), len(data)), nil
}

func (s *kmsServer) decrypt(ctx context.Context, request *DecryptRequest) (*DecryptResponse, string, error) {
	var data wrappedData
	if err := cbor.Unmarshal(request.Ciphertext, &data); err != nil {
		return nil, "", fmt.Errorf("failed to deserialize wrapped cipher data: %v", err)
	}
	if data.Version != 1 {
		return nil, "", fmt.Errorf("unknown version for wrapped cipher data: %v", data.Version)
	}
	if request.KeyId != s.hash {
		return nil, "", fmt.Errorf("KeyId does not match. Expected: %v, found: %v", request.KeyId, s.hash)
	}
	client := s.config.makeClient()
	alg := sdkms.AlgorithmAes
	resp, err := client.Decrypt(ctx, sdkms.DecryptRequest{
		Key:    sdkms.SobjectByID(data.KID),
		Alg:    &alg,
		Cipher: data.Cipher,
		Mode:   sdkms.CryptModeSymmetric(sdkms.CipherModeGcm),
		Iv:     &data.IV,
		Tag:    &data.Tag,
	})
	if err != nil {
		return nil, "", err
	}
	return &DecryptResponse{Plaintext: resp.Plain}, fmt.Sprintf("cipher: %v bytes, plain: %v bytes", len(request.Ciphertext), len(resp.Plain)), nil
}

func logRequest(kind string, msg string, err error) {
	outcome := "was successful"
	if err != nil {
		outcome = fmt.Sprintf("ran into error: %v", err)
	}
	log.Printf("%s request %s. %s", kind, outcome, msg)
}
