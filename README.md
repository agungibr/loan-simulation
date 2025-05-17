# 🏦 TraPinjaman Online

TraPinjaman Online is a simple loan and credit simulation application implemented in Go. The application simulates a loan system with fixed or variable interest schemes, allowing users to manage borrowers, loan amounts, tenors, and payment status.

## ✨ Main Features

1. **🔐 User Authentication**
   - User registration and login
   - Secure password handling
   - Login attempt limitation (3 attempts)

2. **💰 Loan Management**
   - Add, modify, and delete borrower data
   - Apply for loans with different amounts and tenors
   - Automatic calculation of interest rates and monthly installments

3. **🔍 Search Capabilities**
   - Sequential Search: Linear search through borrower data
   - Binary Search: Efficient search through sorted borrower data

4. **📊 Data Sorting**
   - Selection Sort: Sort loans by amount or tenor
   - Insertion Sort: Alternative sorting method for loans

5. **📝 Reporting**
   - Total loans granted
   - Payment status statistics
   - Loan value summaries

## 🧩 Code Structure

The application is structured into several modules:

- `main1.go`: Application entry point with enhanced authentication
- `auth.go`: Authentication and registration related functions
- `model.go`: Data models and structures with immutable patterns
- `loan.go`: Loan operations and calculations
- `search.go`: Search algorithm implementations
- `sort.go`: Sorting algorithm implementations
- `report.go`: Reporting and data analysis
- `db[1].go`: Data structure and database definitions
- `seed.go`: Sample data initialization


## 🔧 How to Run

```bash
go run $(ls *.go | grep -v main1.go)
```

The command above will run all Go files except main1.go to avoid the duplicate main function error.

## 📱 Usage

1. **Registration and Login**
   - Register as a new user or login with existing credentials
   - The system limits login attempts to 3 times

2. **Main Menu**
   - Navigate the main menu to access various features
   - Apply for loans with various terms and amounts
   - Manage existing loans
   - Search and sort loan data
   - View reports and statistics

3. **Loan Data Processing**
   - Using efficient sorting and searching algorithms
   - Implementation of higher-order functions for code flexibility
   - Immutable patterns for safer state management

## 🌟 Key Improvements

- **Registration System**: Ability to add new users
- **Login Security**: Login attempt limitation for system security
- **Modular Structure**: Separation of code into focused modules