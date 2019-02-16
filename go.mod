module ggz-server

require (
	github.com/AndreasBriese/bbloom v0.0.0-20180913140656-343706a395b7 // indirect
	github.com/dgraph-io/badger v1.5.5-0.20190214192501-3196cc1d7a5f
	github.com/dgryski/go-farm v0.0.0-20190104051053-3adb47b1fb0f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/json-iterator/go v1.1.5
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/rs/cors v1.6.0
	github.com/sosop/gitlabClient v0.0.0-20190214050246-a02196e4fbf1
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/urfave/negroni v1.0.0
	golang.org/x/net v0.0.0-20190125091013-d26f9f9a57f3 // indirect
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4 // indirect
	golang.org/x/sys v0.0.0-20190130150945-aca44879d564 // indirect
)

replace (
	golang.org/x/net v0.0.0-20190125091013-d26f9f9a57f3 => github.com/golang/net v0.0.0-20190125091013-d26f9f9a57f3
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4 => github.com/golang/sync v0.0.0-20181221193216-37e7f081c4d4
	golang.org/x/sys v0.0.0-20190130150945-aca44879d564 => github.com/golang/sys v0.0.0-20190130150945-aca44879d564
)
