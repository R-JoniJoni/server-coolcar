package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const privateKey string = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA3egom/VmJG6k6z7uR8kvjGr9rVJ0foFvKZ/wNabyRPWR3+oj\namJo7VT9lgs7fZbLuE1p+V150ZWf+rwswqvXkoFwDs7kcGUN6+ja6/n3eNEWoSzU\ngtMosspwulrF7D0Xh30I/BkwWOi2eGk8r0MAojcsTNdPwTXVtbQG0k8cfIJJf9t7\nXa9FkvEubbTMBwMfNnNlp3qUmTih4Z7CJQZWbe+MSQYhW6taFGW0Gd3Ut9YHQV7l\ni374zPpGaRFWLOPspa4zUNXUm66l0meg3KEo5hAr98GwZLLpXzeSqYhFcnpE7Dqk\nU08nRnySABV5TaFRA5y6Ww2KqKthqcr5Ah8FqwIDAQABAoIBAAhHWgR/gkEHs7Vn\nEqKw+comT7KAqgHyOEB4TBDkWpOCCeAtrwaQB1QbYJ6RarXDw3Prx3FbMGlGSMKk\n9JKKsK8xjwmuZE1hJ8TOWiSFndrvCgYXUxJSaGlLLit1qX6nxNH0MzqcgfY/MFeL\nrYzoVgS3RDrVqY8OwBtYTQzZkLmAXQa+O+d5gAGMdSazzDTI9CXCH92X803uBaCJ\nKsQTOjuvn5s28DzSiUChkB5MuHxucDW6CfaLVucvCKAujvLN5rmFM5fBIArW4Rip\nRHeDvBWtZADpKSbB6MeyFF6/n61w/Qy5A4d3Qa6HSM9JVfCPdFbxUmlsBQcAOajl\n40Z7vCkCgYEA/QHk+0+gH8Hq9qRiQd4T2EWR+WT9YgmiQ3+vjIBQhqU27i3PRL+j\nZ/2OzpXIbhlcdKbLS7dRWcfXaB5bRToO7yqpf6JNtf2SoMxWHO0HIXHGPGOcFBRV\nsnRE6z71qAK1vkRavqFMfohB7JTo8AS4pq8zhj3DM9GrABe7UNNysO0CgYEA4IgX\nhMwNvHrOBgkXmei8FY0R1dSR1bOSOhjEqxhoghpnprV9fT/qaVQ3CmLOA/jyVVXC\npS4xB5BmANJIbZK63RVBGpwiKOHcXdxA7sGdoYjdmeZSasw4032m8kwayiNefk9J\n73y9Ht8BO7aLlVYqOF7qr0XD9SY8Tps3vYDCdfcCgYA9klF0a2tPbzTMMzMKTo5L\nypp8s/wJ+Mg5XwCM08lFCz3z9sgDNhQBQa6YTdFKIffjF5hP49vzWnPsjb6ueTOb\nNqmrOwdoa75cTX4DahebJwIUPjWEmXJFjJAdI+RKr+Yk2KOw0rY38NcCSTbq+msp\nfWevmqY/nR3dVukqAVte/QKBgQDTZW53OkNIjHrC1fahkqzawYnkQBHaGQp4pm8s\nA3wJ2mByZfezX+UMrBxyK9p5hP5r96WeLWI+E+blqRZGC9rhYix8qDnFMflvaXq9\nEA+gUaMKTf6UzJhIDsqK37ptTGWgGHitAU1x7lZT6Sd8P7baggsFYHMBsbEf9SXC\nxqLPuwKBgDmxyYuK8ue+/but7q6C3yL0f1AxBNO01OnQtMA1U8BZyAMLCy0Hkgfa\nwNwQeagPtoBck8WzYOKm+0mezRi+Mb2fV+Qv7pPUPKO0BGMoT9rDxASeR4FMneve\nW0Uz4v2yTpBuSwPHhmjQFL0g86Lw1ZIJ22U7wgCOY09VaoQ1qyhC\n-----END RSA PRIVATE KEY-----"
const want string = "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyMzkxMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY2Fubm9JbkRNYWpvciIsInN1YiI6IjYyMTg0NzBiNDdiMzc2ZWUwYmRmMDVhNCJ9.DKE2ZZSifSeC1S_c-03GfQ35BXOyj-h33IvZpEjJIlmRwJCypeNN90TZM4A_6glvWHrgJFUCaVkuLChzizcc28MURCyXvX3B--g-5x7Zf0ZuhY3jBwHThMrYDQKjf6CdEd68cK6KEIn6R7_u1T2AUr3XnAizLNYjAiW7fZ-Vnc81xBF5xFdKf3U9NinMp2amHLJByhzQ2cnx04OC3VyjpMtzB2CLHUFSS69SSCRHvCSR7aJ4ggjdgZbACtGEtNTTZ0Su23DZeaIhe-GpUddHMbO6yFBFnVQVQiGofbwP_BBBlRpeliOwIFkrYdimSoOcI0QvTMwLyhJ6mgEEnjrLVA"

func TestJWTTokenGen_GenerateToken(t *testing.T){
	// change the format of privateKey
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}

	// make a NewJWTTokenGen and set its time function
	jwtObj := NewJWTTokenGen("cannoInDMajor", key)
	jwtObj.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}

	// use the GenerateToken function. Get the token we want
	token, err := jwtObj.GenerateToken("6218470b47b376ee0bdf05a4", 1e11)
	if err != nil {
		t.Errorf("cannot use method GenerateToken: %v", err)
	}

	if want != token {
		t.Errorf("wrong answer. \nWant: %q\nGot:  %q\n", want, token)
	}
}