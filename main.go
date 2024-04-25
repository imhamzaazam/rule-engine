package main

import (
	"context"
	"fmt"
	"github.com/ezachrisen/indigo"
	"github.com/ezachrisen/indigo/cel"
	"google.golang.org/genproto/googleapis/type/decimal"
	pb "rule-engine/protos"
	"strings"
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
			SchemeMerchantCategoryCode: "111",
			TotalAuthAmount:            100,
		},
	}

	//data2 := map[string]interface{}{
	//	"flights": map[string]int64{"1111": 1.0, "DL232": 1.0, "AA1622": 1.0},
	//}

	listOfMcc := []string{"111", "222", "333"}
	expr := formatString(listOfMcc)

	txnRule := indigo.Rule{
		ID:     "transaction_rule",
		Schema: txnRuleSchema,
		Expr:   expr,
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
	fmt.Println(results.ExpressionPass)
}

func formatString(listOfMcc []string) string {
	var mccExpressions []string
	for _, mcc := range listOfMcc {
		mccWithQuotes := fmt.Sprintf(`"%s"`, mcc)
		mccExpressions = append(mccExpressions, mccWithQuotes)
	}
	expr := fmt.Sprintf("[%s].exists(mcc, mcc == txn.scheme_merchant_category_code)",
		strings.Join(mccExpressions, ", "))
	return expr
}
