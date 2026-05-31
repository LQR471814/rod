module github.com/go-rod/rod

go 1.21

require (
	github.com/bytedance/sonic v1.15.1
	github.com/ysmood/fetchup v0.2.3
	github.com/ysmood/goob v0.4.0
	github.com/ysmood/got v0.40.0
	github.com/ysmood/gotrace v0.6.0
	github.com/ysmood/gson v0.7.3
	github.com/ysmood/leakless v0.9.0
)

require (
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.5.1 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ysmood/gop v0.2.0 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/sys v0.22.0 // indirect
)

replace encoding/json => ./rod/lib/sonicjson
