module github.com/allocamelus/allocamelus

go 1.16

// TODO: Wait for https://github.com/kubernetes/klog/pull/242
exclude github.com/go-logr/logr v1.0.0

require (
	github.com/ProtonMail/go-crypto v0.0.0-20210707164159-52430bf6b52c // indirect
	github.com/ProtonMail/gopenpgp/v2 v2.2.2
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/go-redis/redis/v8 v8.11.3
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gofiber/fiber/v2 v2.17.0
	github.com/gofiber/helmet/v2 v2.2.0
	github.com/h2non/bimg v1.1.5
	github.com/json-iterator/go v1.1.11
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/microcosm-cc/bluemonday v1.0.15
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mr-tron/base58 v1.2.0
	github.com/nbutton23/zxcvbn-go v0.0.0-20210217022336-fa2cb2858354
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/tinylib/msgp v1.1.6
	github.com/tkrajina/typescriptify-golang-structs v0.1.5
	github.com/valyala/fasthttp v1.28.0 // indirect
	github.com/valyala/fastjson v1.6.3
	github.com/valyala/quicktemplate v1.6.3
	github.com/xhit/go-simple-mail/v2 v2.10.0
	golang.org/x/crypto v0.0.0-20210813211128-0a44fdfbc16e
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/gographics/imagick.v3 v3.4.0
	k8s.io/klog/v2 v2.10.0
)
