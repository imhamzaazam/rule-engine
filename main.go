package main

import (
	"context"
	"fmt"
	"github.com/ezachrisen/indigo"
	"github.com/ezachrisen/indigo/cel"
	"google.golang.org/genproto/googleapis/type/decimal"
	pb "rule-engine/protos"
	"time"
)

type Turbo struct {
	CreatedAt       time.Time `json:"created_at"`
	Code            string    `json:"code"`
	TxCelExpression string    `json:"tx_cel_expression"`
	// TODO set Percentage to NOT NULL
	Percentage decimal.Decimal `json:"percentage"`
}

func main() {

	txnRuleSchema := indigo.Schema{
		Description: "Schema of the transaction",
		Elements: []indigo.DataElement{
			{Name: "txn", Type: indigo.Proto{Message: &pb.TransactionRuleSchema{}}},
		},
	}

	data := map[string]interface{}{
		"txn": &pb.TransactionRuleSchema{
			SchemeMerchantCategoryCode: "222",
			TotalAuthAmount:            990,
		},
	}

	txnRule := indigo.Rule{
		ID:         "transaction_rule",
		Schema:     txnRuleSchema,
		ResultType: indigo.Float{},
		Expr: "(txn.total_auth_amount >= 100 && " +
			"txn.scheme_merchant_category_code in " +
			"{'111': 1.0, " +
			"'222': 2.2, " +
			"'333': 3.0}) ? " +
			"({'111': 1.0, " +
			"'222': 2.2, " +
			"'333': 3.0}[txn.scheme_merchant_category_code]) " +
			": 0.0",
	}

	engine := indigo.NewEngine(cel.NewEvaluator())
	err := engine.Compile(&txnRule)

	if err != nil {
		fmt.Printf("Error adding rule %v", err)

	}
	results, err := engine.Eval(context.Background(), &txnRule, data)
	if err != nil {
		fmt.Printf("Error evaluating: %v", err)
		return
	}
	fmt.Println(results.Pass)
	sum := results.Value
	fmt.Println(sum)

}
