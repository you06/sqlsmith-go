module github.com/you06/sqlsmith-go

go 1.13

require (
	github.com/juju/errors v0.0.0-20190930114154-d42613fe1ab9
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/pingcap/parser v0.0.0-20191031081038-bfb0c3adf567
	github.com/pingcap/tidb v1.1.0-beta.0.20191106105829-1b72ce5987b3
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/stretchr/testify v1.3.0
)

replace github.com/pingcap/tidb => github.com/you06/tidb v1.1.0-beta.0.20191107073350-1f019a46fa2c
