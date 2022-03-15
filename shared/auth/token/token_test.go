package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const publicKey string = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3egom/VmJG6k6z7uR8kv\njGr9rVJ0foFvKZ/wNabyRPWR3+ojamJo7VT9lgs7fZbLuE1p+V150ZWf+rwswqvX\nkoFwDs7kcGUN6+ja6/n3eNEWoSzUgtMosspwulrF7D0Xh30I/BkwWOi2eGk8r0MA\nojcsTNdPwTXVtbQG0k8cfIJJf9t7Xa9FkvEubbTMBwMfNnNlp3qUmTih4Z7CJQZW\nbe+MSQYhW6taFGW0Gd3Ut9YHQV7li374zPpGaRFWLOPspa4zUNXUm66l0meg3KEo\n5hAr98GwZLLpXzeSqYhFcnpE7DqkU08nRnySABV5TaFRA5y6Ww2KqKthqcr5Ah8F\nqwIDAQAB\n-----END PUBLIC KEY-----"

func TestJWTTokenVerifier_Verify(t *testing.T) {
	// 转换 public key 的格式
	pk, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key string: %v", err)
	}

	// 设置jwt包里的时间，即这个测试文档所认为的现在的时间
	jwt.TimeFunc = func() time.Time {
		return time.Unix(1516239032, 0)
	}

	verifierObj := &JWTTokenVerifier{PublicKey: pk}

	// 设计test table
	cases := []struct{
		name 	string
		tkn 	string
		wantErr	bool
		want 	string
	}{
		{
			name: "good_token",
			tkn: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyMzkxMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY2Fubm9JbkRNYWpvciIsInN1YiI6IjYyMTg0NzBiNDdiMzc2ZWUwYmRmMDVhNCJ9.DKE2ZZSifSeC1S_c-03GfQ35BXOyj-h33IvZpEjJIlmRwJCypeNN90TZM4A_6glvWHrgJFUCaVkuLChzizcc28MURCyXvX3B--g-5x7Zf0ZuhY3jBwHThMrYDQKjf6CdEd68cK6KEIn6R7_u1T2AUr3XnAizLNYjAiW7fZ-Vnc81xBF5xFdKf3U9NinMp2amHLJByhzQ2cnx04OC3VyjpMtzB2CLHUFSS69SSCRHvCSR7aJ4ggjdgZbACtGEtNTTZ0Su23DZeaIhe-GpUddHMbO6yFBFnVQVQiGofbwP_BBBlRpeliOwIFkrYdimSoOcI0QvTMwLyhJ6mgEEnjrLVA",
			wantErr: false,
			want: "6218470b47b376ee0bdf05a4",
		},
		{
			name: "expired",
			tkn: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyMzkwMjMsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY2Fubm9JbkRNYWpvciIsInN1YiI6IjYyMTg0NzBiNDdiMzc2ZWUwYmRmMDVhNCJ9.DDERnWdOWXA24tKubItnQrnNmDqpuwpjCTqoKj4SbubvG1dVlz6ciJN5qJm0rZkJWq8isKqVfgr8Drpf7HASv03SW4NNrJikA5VvqZezAl5RDCxpad07QNQ6acOtbQGT8hSlYOu29UQk-WgjWeTf6Db4wWB1SobiPpiI41utMGiW4OGyPOUaYpAnBEZgPKaBtMypDpM2EJOvMqTf-Lyzk__4j9iskZeCY1QgLKPP5tYymDibdkxYj5tU_XU9FZmOhnv8Xj2Dx_UOhL80iUVr4ls-dxK249bjAeIKTlcGN1oGl_bk6tHMKX-IL6MUhIuRDRoLRwEErQ_c11cGmaBkvw",
			wantErr: true,
			want: "",
		},
		{
			name: "not_token",
			tkn: "团长你就是歌姬吧",
			wantErr: true,
			want: "",
		},
		{
			name: "fake_signiture",
			tkn: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyMzkxMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY2Fubm9JbkRNYWpvciIsInN1YiI6IjYyMTg0NzBiNDdiMzc2ZWUwYmRmMDVhNSJ9.DKE2ZZSifSeC1S_c-03GfQ35BXOyj-h33IvZpEjJIlmRwJCypeNN90TZM4A_6glvWHrgJFUCaVkuLChzizcc28MURCyXvX3B--g-5x7Zf0ZuhY3jBwHThMrYDQKjf6CdEd68cK6KEIn6R7_u1T2AUr3XnAizLNYjAiW7fZ-Vnc81xBF5xFdKf3U9NinMp2amHLJByhzQ2cnx04OC3VyjpMtzB2CLHUFSS69SSCRHvCSR7aJ4ggjdgZbACtGEtNTTZ0Su23DZeaIhe-GpUddHMbO6yFBFnVQVQiGofbwP_BBBlRpeliOwIFkrYdimSoOcI0QvTMwLyhJ6mgEEnjrLVA",
			wantErr: true,
			want: "",
		},
	}

	// 进行各个情况的测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {	// 进入一个子函数（可见Run的定义）
			got, err := verifierObj.Verify(c.tkn)
			if err != nil && !c.wantErr {
				t.Errorf("Wanted correct but error occured: %v", err)
			}
			if err == nil && c.wantErr {
				t.Errorf("Wanted error but it turned out correct: %v", err)
			}
			if c.want != got {
				t.Errorf("wrong result.\n Wanted %q\nGot    %q", c.want, got)
			}
		})
	}
}

