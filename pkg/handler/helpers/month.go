package helpers

import (
	
	"fmt"
	"storie/pkg/domain"
    "strings"
	"time"
)

func CountTransactionsPerMonth(transactions []domain.Transaction) string {
	counts := make(map[string]int)

	for _, transaction := range transactions {
		date, err := ParseDate(transaction.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}

	

		monthNumber := date.Month()

		month :=  monthNumber.String()

		
		counts[month]++
	}

	
	months := make([]domain.Month, 0)


	for month, count := range counts {
		months = append(months, domain.Month{Count: count, Name: month})
	}

	return GenerateSummaryString( months)
}

func ParseDate(dateStr string) (date time.Time, err error) {
	date, err = time.Parse("2006-01-02", dateStr)
	return
}

func 	GenerateSummaryString(months []domain.Month) string {
	var summaryStrings []string

	for _, month := range months {
		summaryItem := fmt.Sprintf(`<div class="summary-item">
      <span>Number of transaction in %s:</span>
      <span>%d</span>
    </div>`, month.Name, month.Count)

		summaryStrings = append(summaryStrings, summaryItem)
	}

	return strings.Join(summaryStrings, "\n")
}
