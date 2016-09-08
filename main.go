package main

import (
	"./parser"
	"fmt"
	"strings"
)

func main() {
	source := `
	{Withdrawal, sort_by_score, ...} = import('records/activity')

	; Internal record used to handle the combine algorithm
	record State{
	  activities: [],
	  withdrawals: [],
	  fees: []
	}

	; Return a vector of activities w/ withdrawals & ATM fees merged.
	; @param [Vector] activities
	;   uncombined feed input
	; @return [Vector]
	;   activities with withdrawals and ATM fees combined
	function combine_fees(activities) {
	  result = reduce(
	    #(state = {activities, withdrawals, fees}, activity) {
	      match (activity.tags) {
		when [..., 'non_bank_atm_withdrawal', ...]:
                  State{...state, :withdrawals: push(withdrawals, activity)}
	        when [..., 'non_bank_atm_operator_fee', ...]:
                  State{...state, :fees: push(fees, activity)}
               else:
                 State{...state, :activities: push(activities, activity)}
	      }
	    },
	    activities,
	    State{}
	  )

	  process_result(result)
	}

	; Combine withdrawals and fees while possible
	function process_result({activities, withdrawals, fees}) {
	  match ([withdrawals, fees]) {
	    when [[], []]:
	      sort_by_score(activities)
	    else:
	      state = match ([withdrawals, fees]) {
		when [[w, ...ws], [f, ...fs]]:
		  State{
		    :activities: push(activities, Withdrawal{...w, :fee: f}),
		    :withdrawals: ws,
		    :fees: fs
		  }
		else:
		  State{:activities: push(activities, head(withdrawals) or head(fees))}
	      }
	      process_result(state)
	  }
	}
	`
	err, ast := parser.Parse(strings.NewReader(source), "/some/example.bjx")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", ast)
}
