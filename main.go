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
			SchemeMerchantCategoryCode: "444",
			TotalAuthAmount:            990,
		},
	}
	txnRule := indigo.Rule{
		ID:         "transaction_rule",
		Schema:     txnRuleSchema,
		ResultType: indigo.Float{},
		Rules:      map[string]*indigo.Rule{},
	}

	txnRule.Rules["1_percent"] = &indigo.Rule{
		ID:         "1_percent",
		Schema:     txnRuleSchema,
		ResultType: indigo.Float{},
		Expr: `
		[
		   '111',
		   '222',
		   '333',
		].exists(mcc, mcc == txn.scheme_merchant_category_code) ?
		{
		   '111': 1.0,
		   '222': 2.2,
		   '333': 3.1,
		}[txn.scheme_merchant_category_code] : 0.0
		`,
	}

	txnRule.Rules["2_percent"] = &indigo.Rule{
		ID:         "2_percent",
		Schema:     txnRuleSchema,
		ResultType: indigo.Float{},
		Expr: `
		[
		   '111',
		   '222',
		   '333',
		   '444',
		].exists(mcc, mcc == txn.scheme_merchant_category_code) ?
		{
		   '111': 1.0,
		   '222': 2.2,
		   '333': 3.1,
           '444': 5.5,
		}[txn.scheme_merchant_category_code] : 0.0
		`,
	}

	engine := indigo.NewEngine(cel.NewEvaluator())
	err := engine.Compile(&txnRule)

	if err != nil {
		fmt.Printf("Error adding rule %v", err)

	}
	results, err := engine.Eval(context.Background(), txnRule.Rules["2_percent"], data)
	if err != nil {
		fmt.Printf("Error evaluating: %v", err)
		return
	}
	fmt.Println(results.Pass)
	sum := results.Value
	fmt.Println(sum)

}
