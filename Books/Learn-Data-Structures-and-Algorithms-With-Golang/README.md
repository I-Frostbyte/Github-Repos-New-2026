# 🚀 Learn Data Structures and Algorithms with Golang

## Author
Bhagvan Kommadi

## Description:
Explore Golang's data structures and algorithms to design, implement, and analyze code in the professional settingKey FeaturesLearn the basics of data structures and algorithms and implement them efficiently
Use data structures such as arrays, stacks, trees, lists and graphs in real-world scenarios
Compare the complexity of different algorithms and data structures for improved code performanceBook DescriptionGolang is one of the fastest growing programming languages in the software industry. Its speed, simplicity, and reliability make it the perfect choice for building robust applications. This brings the need to have a solid foundation in data structures and algorithms with Go so as to build scalable applications. Complete with hands-on tutorials, this book will guide you in using the best data structures and algorithms for problem solving.
The book begins with an introduction to Go data structures and algorithms. You'll learn how to store data using linked lists, arrays, stacks, and queues. Moving ahead, you'll discover how to implement sorting and searching algorithms, followed by binary search trees. This book will also help you improve the performance of your applications by stringing data types and implementing hash structures in algorithm design. Finally, you'll be able to apply traditional data structures to solve real-world problems.
By the end of the book, you'll have become adept at implementing classic data structures and algorithms in Go, propelling you to become a confident Go programmer.What you will learnImprove application performance using the most suitable data structure and algorithm
Explore the wide range of classic algorithms such as recursion and hashing algorithms
Work with algorithms such as garbage collection for efficient memory management 
Analyze the cost and benefit trade-off to identify algorithms and data structures for problem solving
Explore techniques for writing pseudocode algorithm and ace whiteboard coding in interviews
Discover the pitfalls in selecting data structures and algorithms by predicting their speed and efficiencyWho this book is forThis book is for developers who want to understand how to select the best data structures and algorithms that will help solve coding problems. Basic Go programming experience will be an added advantage.Downloading the example code for this book You can download the example code files for all Packt books you have purchased from your account at http://www.PacktPub.com. If you purchased this book elsewhere, you can visit http://www.PacktPub.com/support and register to have the files e-mailed directly to you.

## Folder Structure

This folder contains the code written while following the book, split into two
independent Go modules (each with its own `go.mod` / `Makefile`).

| Folder | Contents |
| --- | --- |
| [Searching](Searching) | Search algorithm implementations: [1-Linear](Searching/1-Linear), [2-Binary](Searching/2-Binary), plus a shared [generator](Searching/generator) used by the tests. |
| [Sorting](Sorting) | Six classic sorting algorithms: [1-Bubble](Sorting/1-Bubble), [2-Selection](Sorting/2-Selection), [3-Insertion](Sorting/3-Insertion), [4-Shell](Sorting/4-Shell), [5-Merge](Sorting/5-Merge), [6-Quick](Sorting/6-Quick). |

## Running the Code

Each sub-folder is a standalone Go module. Run the tests from within
`Searching/` or `Sorting/`:

```bash
make test        # go test -v ./...
```
