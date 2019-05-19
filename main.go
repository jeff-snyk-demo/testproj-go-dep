package main

import (
	"fmt"
	accounting "github.com/leekchan/accounting"
	"github.com/aws/aws-sdk-go"
	"cloud.google.com/go/storage"
	"github.com/ghodss/yaml"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gruntwork-io/gruntwork-cli/collections"
	"github.com/gruntwork-io/gruntwork-cli/entrypoint"
	"github.com/gruntwork-io/gruntwork-cli/errors"
	"github.com/gruntwork-io/gruntwork-cli/files"
	"github.com/gruntwork-io/gruntwork-cli/logging"
	"github.com/jstemmer/go-junit-report/formatter"
	"github.com/jstemmer/go-junit-report/parser"
	"github.com/magiconair/properties/assert"
	"github.com/docker/docker"
	"github.com/gogits/gogs"
)

func main() {
	fmt.Println("hello again")

	// ac := accounting.Accounting{Symbol: "$", Precision: 2}
	// fmt.Println(ac.FormatMoney(123456789.213123))
}
