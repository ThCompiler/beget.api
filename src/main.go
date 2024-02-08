package main

import (
	"fmt"
	"log"

	"github.com/ThCompiler/go.beget.api/pkg/beget/api/dns"
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/dns/build"
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/result"
	"github.com/ThCompiler/go.beget.api/pkg/beget/core"
)

func main() {
	client := core.Client{
		Login:    "vetan2o5",
		Password: "932usRC26Tp",
	}

	req, err := core.PrepareRequest[result.BoolResult](
		client,
		dns.CallChangeRecords("tmp.thecompiler.pw",
			dns.SetBasicRecords(
				build.NewBasicRecordsCreator().
					AddARecords(
						build.NewARecordCreator().
							AddRecord(10, "8.8.8.8"),
					).Create(),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := req.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", *resp)
}
