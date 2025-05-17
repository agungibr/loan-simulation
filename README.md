# TraPinjaman Online

TraPinjaman Online is a simple loan and credit simulation system implemented in Go. The application simulates a loan system with fixed or variable interest schemes, allowing users to manage borrowers, loan amounts, tenors, and payment status.

## Features

1. **User Authentication**
   - User registration and login
   - Secure password handling
   - Login attempt limitation (3 attempts)

2. **Loan Management**
   - Add, modify, and delete borrower data
   - Apply for loans with different amounts and tenors
   - Calculate interest rates and monthly installments

3. **Search Capabilities**
   - Sequential Search: Linear search through borrower data
   - Binary Search: Efficient search through sorted borrower data

4. **Data Sorting**
   - Selection Sort: Sort loans by amount or tenor
   - Insertion Sort: Alternative sorting method for loans

5. **Reporting**
   - Total loans granted
   - Payment status statistics
   - Loan value summaries

## Code Structure

The application is structured into several modules:

- `main.go`: Entry point of the application
- `auth.go`: Authentication-related functions
- `model.go`: Data models and structures
- `loan.go`: Loan operations and calculations
- `search.go`: Search algorithms implementation
- `sort.go`: Sorting algorithms implementation
- `report.go`: Reporting and data analysis
- `util.go`: Utility functions and helpers

## How to Run

```bash
go run *.go
```

## Usage

1. Register as a new user or login with existing credentials
2. Navigate the main menu to access different features
3. Apply for loans with various terms and amounts
4. Manage existing loans
5. Search and sort through loan data
6. View reports and statistics