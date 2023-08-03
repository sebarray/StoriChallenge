package helpers

import (
  "storie/pkg/domain"
  "fmt"
)

func SendHtml(mail domain.Mail) string {
html:=`	<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Resumen de Transacciones</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      max-width: 600px;
      margin: 0 auto;
      padding: 20px;
    }

    .logo {
      display: block;
      max-width: 200px;
      margin: 0 auto;
    }

    .summary {
      margin-top: 20px;
      padding: 10px;
      border: 1px solid #ccc;
      background-color: #f9f9f9;
    }

    .summary-title {
      font-weight: bold;
      text-decoration: underline;
    }

    .summary-item {
      margin-bottom: 10px;
    }
  </style>
</head>
<body> <p>Dear `+mail.Name+`,</p>
<p>Please find attached the transaction history of your credit card. If you have any inquiries or need further assistance, please don't hesitate to contact us.</p>

<img src="https://i.ibb.co/vwrZ8DP/complete-logo-0f6b7ce5-1.png" alt="complete-logo-0f6b7ce5-1" border="0">

  <div class="summary">
    <h2 class="summary-title">transaction history</h2>
    <div class="summary-item">
      <span>Total balance:</span>
      <span>` + fmt.Sprintf("%.2f", mail.Balance) + `</span>
    </div>`
  
 
 HTMLEnd  :=  `<div class="summary-item">
      <span>average debit amount:</span>
      <span>` + fmt.Sprintf("%.2f", mail.AverageDebit) +`</span>
    </div>    
    <div class="summary-item">
      <span>average amount of credit:</span>
      <span>` + fmt.Sprintf("%.2f", mail.AverageCredit) +`</span>
    </div> 
    <a href="`+mail.Link +`" download style="display: inline-block; background-color: #f2f2f2; color: #333; border: 1px solid #ccc; padding: 10px 20px; text-decoration: none; border-radius: 4px; transition: background-color 0.3s ease;">
    download file
  </a>
  </div>
</body>
</html>`


return  html + CountTransactionsPerMonth(mail.Transaction) + HTMLEnd 
}