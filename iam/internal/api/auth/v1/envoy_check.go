package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	"github.com/HeyReyHR/twitch-clone/iam/internal/utils/jwt_tokens"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	authv3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	typev3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

const (
	AccessTokenCookieName = "X-Access-Token"

	HeaderUserId       = "X-User-Id"
	HeaderUserUsername = "X-User-Username"
	HeaderUserRole     = "X-User-Role"

	HeaderContentType = "content-type"
	HeaderAuthStatus  = "X-Auth-Status"

	HeaderCookie        = "cookie"
	HeaderAuthorization = "authorization"

	ContentTypeJSON = "application/json"

	AuthStatusDenied = "denied"
)

func (a *api) Check(ctx context.Context, req *authv3.CheckRequest) (*authv3.CheckResponse, error) {
	accessToken, err := extractAccessToken(req)
	if accessToken == "" {
		return deniedResponse("Missing access token", 403), nil
	}

	claims, err := jwt_tokens.ValidateAccessToken(accessToken)
	if err != nil {
		return deniedResponse("Invalid access token", 403), nil
	}

	return allowedResponse(claims), nil
}

func extractAccessToken(req *authv3.CheckRequest) (string, error) {
	if req.Attributes == nil || req.Attributes.Request == nil {
		return "", model.ErrNoHTTPRequest
	}
	headers := req.Attributes.Request.Http.Headers

	if authHeader, ok := headers[HeaderAuthorization]; ok && authHeader != "" {
		if strings.HasPrefix(authHeader, "Bearer ") {
			return strings.TrimPrefix(authHeader, "Bearer "), nil
		}
	}

	if cookieHeader, ok := headers[HeaderCookie]; ok && cookieHeader != "" {
		accessToken := extractAccessTokenFromCookies(cookieHeader)
		if accessToken != "" {
			return accessToken, nil
		}
	}

	return "", model.ErrMalformedToken
}

func extractAccessTokenFromCookies(cookieHeader string) string {
	req := &http.Request{Header: make(http.Header)}
	req.Header.Add(HeaderCookie, cookieHeader)

	if cookie, err := req.Cookie(AccessTokenCookieName); err == nil {
		var accessToken string
		accessToken, err = url.QueryUnescape(cookie.Value)
		if err != nil {
			return cookie.Value
		}
		return accessToken
	}
	return ""
}

func allowedResponse(claims *model.Claims) *authv3.CheckResponse {
	return &authv3.CheckResponse{
		Status: &status.Status{
			Code: int32(codes.OK),
		},
		HttpResponse: &authv3.CheckResponse_OkResponse{
			OkResponse: &authv3.OkHttpResponse{
				Headers: []*corev3.HeaderValueOption{
					{
						Header: &corev3.HeaderValue{
							Key:   HeaderUserId,
							Value: claims.UserId,
						},
					},
					{
						Header: &corev3.HeaderValue{
							Key:   HeaderUserUsername,
							Value: claims.Username,
						},
					},
					{
						Header: &corev3.HeaderValue{
							Key:   HeaderUserRole,
							Value: string(claims.Role),
						},
					},
				},
			},
		},
	}
}

// deniedResponse creates a denied authorization response
func deniedResponse(message string, statusCode int32) *authv3.CheckResponse {
	return &authv3.CheckResponse{
		Status: &status.Status{
			Code: int32(codes.PermissionDenied),
		},
		HttpResponse: &authv3.CheckResponse_DeniedResponse{
			DeniedResponse: &authv3.DeniedHttpResponse{
				Status: &typev3.HttpStatus{
					Code: typev3.StatusCode(statusCode),
				},
				Body: fmt.Sprintf(`{"error": "%s", "timestamp": "%s"}`,
					message, time.Now().Format(time.RFC3339)),
				Headers: []*corev3.HeaderValueOption{
					{
						Header: &corev3.HeaderValue{
							Key:   HeaderAuthStatus,
							Value: AuthStatusDenied,
						},
					},
					{
						Header: &corev3.HeaderValue{
							Key:   HeaderContentType,
							Value: ContentTypeJSON,
						},
					},
				},
			},
		},
	}
}
