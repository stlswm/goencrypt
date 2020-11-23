package sha256

import "testing"

var privateKey = `-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALG0DtrgftRRQZWq
YHQN3X6ug4PwTV89KRq9LTmggPFX1FDS9QXHlOManPbbz1SIVHhnvpNXdrjOv8Gv
YjJq4Zhv85DklYNVCDdX2UcZWwJI0FePY7fgkeyAR9154S+8RXoTf4gILPUrYZIq
hlekWxppNJIYus8FRmZYhHaVV43vAgMBAAECgYAl+46RShrR1uxlyR0EsVH+md6S
fPKMwM3gYT8haiUlcNh4zh6BNb+lKzMRsLoIH3uuoL4jUDaYGdjy1UHv9NhzGTHl
jqCJIeQZ5YzGKzdO1VzP1uRn6kRpsHDSPSbh45KfIB0Cx+CsLVyIsQOd3hSdo1HS
p7ZNWmwefLV8U8kWwQJBAOj63sG9oiCtNYet/adk5vx6HekA5zWTx+wq/zZ8PKn9
qGhMdIMPUUif2/zobM+kkDJkFG6yz9btZbNa0hEsQ/MCQQDDQv1GQlG8X8M1XYMx
+I37RjGv9S9uYs3KvxqJHL4/X+SgAzcanhPpnX00Q3BIUEbQ694hKyOSeCCAih38
r9kVAkEAxawdFjrEdX8gvPbWEZIeHbeq6wgWJDI33Vx534u3mO+cVKoR7SUw7TKd
X08BC4hvxCF/6JCUoaIYuP1k4KczGwJAT+bfwZpt3AnL1dCYWSB+6S4GMxy/9gnU
cHzXHPh4GXLiGWB3chrxxw09KN/X2J+Q+vkqAqSmf89MbY0XTEmhKQJAIpizyQQ4
1AOZ0pT65tBtXpqW3pHs0IJIeSJNHc6vD/zj2CGvsyvTg38otlkW0xJkVnvg0jOy
sx9fOSv5eiB5Qw==
-----END PRIVATE KEY-----`
var publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCxtA7a4H7UUUGVqmB0Dd1+roOD
8E1fPSkavS05oIDxV9RQ0vUFx5TjGpz2289UiFR4Z76TV3a4zr/Br2IyauGYb/OQ
5JWDVQg3V9lHGVsCSNBXj2O34JHsgEfdeeEvvEV6E3+ICCz1K2GSKoZXpFsaaTSS
GLrPBUZmWIR2lVeN7wIDAQAB
-----END PUBLIC KEY-----`

func TestSign(t *testing.T) {
	err, signed := Sign([]byte(privateKey), "wang min")
	if err != nil {
		t.Error(err)
	}
	if signed != "BDjfqVF3HczZxz818YlIqBuV18YA4arb5wGrO+Ke+e+3hhRgR4gvH0dvBVeFl9aZjz/oJecyFTImdX63olSJT/91D3BqnIq5dyBBjx8410+3Da6olFfnB7LrOMqSdgj+kZtHpyoehM1CvOVj0Ul7EqtJREmbxxXqH2WQrtJmSb4=" {
		t.Error("mismatch")
	}
}

func TestVerify(t *testing.T) {
	err := Verify([]byte(publicKey), "BDjfqVF3HczZxz818YlIqBuV18YA4arb5wGrO+Ke+e+3hhRgR4gvH0dvBVeFl9aZjz/oJecyFTImdX63olSJT/91D3BqnIq5dyBBjx8410+3Da6olFfnB7LrOMqSdgj+kZtHpyoehM1CvOVj0Ul7EqtJREmbxxXqH2WQrtJmSb4=", "wang min")
	if err != nil {
		t.Error(err)
	}
}
